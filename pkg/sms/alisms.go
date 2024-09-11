package sms

import (
	"encoding/json"
	"sgin/pkg/app"
	"strings"

	"errors"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v4/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

// CreateClient initializes the Aliyun SMS client using the provided AccessKeyId and AccessKeySecret.
func CreateClient(accessKeyId, accessKeySecret string) (*dysmsapi20170525.Client, error) {
	config := &openapi.Config{
		AccessKeyId:     tea.String(accessKeyId),
		AccessKeySecret: tea.String(accessKeySecret),
		Endpoint:        tea.String("dysmsapi.aliyuncs.com"),
	}
	return dysmsapi20170525.NewClient(config)
}

type SmsCodeTemplate struct {
	Code string `json:"code"`
}

// GenerateSmsCodeTemplate generates a JSON string for the SMS code template.
func GenerateSmsCodeTemplate(code string) string {
	template := &SmsCodeTemplate{
		Code: code,
	}
	templateJSON, _ := json.Marshal(template)
	return string(templateJSON)
}

// SendSMS sends an SMS message using the Aliyun SMS service.
func SendSMS(ctx *app.Context, phoneNumber string, templateParam string) error {

	if !ctx.Config.AliyunSMSConfig.Enable {
		ctx.Logger.Info("Aliyun SMS is disabled")
		return nil
	}

	client, err := CreateClient(ctx.Config.AliyunSMSConfig.AccessKeyId, ctx.Config.AliyunSMSConfig.AccessKeySecret)
	if err != nil {
		ctx.Logger.Error("Failed to create Aliyun SMS client", err)
		return err
	}

	signName := ctx.Config.AliyunSMSConfig.SignName
	templateCode := ctx.Config.AliyunSMSConfig.TemplateCode

	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		SignName:      tea.String(signName),
		TemplateCode:  tea.String(templateCode),
		PhoneNumbers:  tea.String(phoneNumber),
		TemplateParam: tea.String(templateParam),
	}

	runtime := &util.RuntimeOptions{}
	resp, err := client.SendSmsWithOptions(sendSmsRequest, runtime)
	if err != nil {
		var sdkErr *tea.SDKError
		if e, ok := err.(*tea.SDKError); ok {
			sdkErr = e
		} else {
			sdkErr = &tea.SDKError{Message: tea.String(err.Error())}
		}
		ctx.Logger.Error("Failed to send SMS", sdkErr.Message)
		if data := sdkErr.Data; data != nil {
			var decodedData interface{}
			d := json.NewDecoder(strings.NewReader(tea.StringValue(data)))
			if err := d.Decode(&decodedData); err == nil {
				if m, ok := decodedData.(map[string]interface{}); ok {
					if recommend, ok := m["Recommend"]; ok {
						// fmt.Println("Recommendation:", recommend)
						ctx.Logger.Error("Recommendation:", recommend)
					}
				}
			}
		}
		return err
	}

	// if *resp.Body.Code != "OK" {
	// 	ctx.Logger.Error("Failed to send SMS", *util.ToJSONString(resp))
	// 	return errors.New("Failed to send SMS")
	// }

	ctx.Logger.Info("SMS send:", *util.ToJSONString(resp))

	if *resp.Body.Code != "OK" {
		ctx.Logger.Error("Failed to send SMS", *util.ToJSONString(resp))
		return errors.New("发生短信失败,请联系管理员,Error:" + *resp.Body.Message)
	}
	// console.Log(util.ToJSONString(resp))
	return nil
}
