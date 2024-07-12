package main

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

// NullableTime is a custom type to handle empty time strings
type NullableTime struct {
	time.Time
}

// UnmarshalJSON handles empty string as null value for time
func (nt *NullableTime) UnmarshalJSON(data []byte) error {
	// Unmarshal into a string to check for empty value
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	// If the string is empty, set to zero time
	if str == "" {
		*nt = NullableTime{time.Time{}}
		return nil
	}

	// Otherwise, parse the time
	parsedTime, err := time.Parse(time.RFC3339, str)
	if err != nil {
		return err
	}

	*nt = NullableTime{parsedTime}
	return nil
}

// MarshalJSON handles the JSON marshaling of NullableTime
func (nt NullableTime) MarshalJSON() ([]byte, error) {
	if nt.Time.IsZero() {
		return json.Marshal("")
	}
	return json.Marshal(nt.Time.Format(time.RFC3339))
}

// Value implements the driver Valuer interface for database serialization
func (nt NullableTime) Value() (driver.Value, error) {
	if nt.Time.IsZero() {
		return nil, nil
	}
	return nt.Time, nil
}

// Scan implements the sql Scanner interface for database deserialization
func (nt *NullableTime) Scan(value interface{}) error {
	if value == nil {
		*nt = NullableTime{time.Time{}}
		return nil
	}
	val, ok := value.(time.Time)
	if !ok {
		return fmt.Errorf("can not convert %v to timestamp", value)
	}
	*nt = NullableTime{val}
	return nil
}

// Customer struct with time pointers
type Customer struct {
	ID           uint         `json:"id" gorm:"primaryKey;comment:'主键ID'"`             // 主键ID
	Name         string       `json:"name" gorm:"comment:'企业名称'"`                      // 企业名称
	Address      string       `json:"address" gorm:"comment:'地址'"`                     // 地址
	ContactInfo  string       `json:"contact_info" gorm:"comment:'联系方式'"`              // 联系方式
	BankAccount  string       `json:"bank_account" gorm:"comment:'银行账号'"`              // 银行账号
	CreditStatus string       `json:"credit_status" gorm:"comment:'信用状态'"`             // 信用状态
	Discount     float64      `json:"discount" gorm:"comment:'折扣'"`                    // 折扣
	Status       int          `json:"status" gorm:"comment:'状态'"`                      // 状态
	CreatedAt    NullableTime `json:"created_at" gorm:"autoCreateTime;comment:'创建时间'"` // 创建时间
	UpdatedAt    NullableTime `json:"updated_at" gorm:"autoUpdateTime;comment:'更新时间'"` // 更新时间
}

func main() {
	jsonData := `{
		"id": 1,
		"name": "Test Company",
		"address": "123 Test St",
		"contact_info": "123-456-7890",
		"bank_account": "1234567890",
		"credit_status": "Good",
		"discount": 10.5,
		"status": 1,
		"created_at": "",
		"updated_at": ""
	}`

	var customer Customer
	if err := json.Unmarshal([]byte(jsonData), &customer); err != nil {
		fmt.Println("Error:", err)
	} else {
		customer.CreatedAt = NullableTime{time.Now()}
		fmt.Println("Customer:", customer)
	}

}
