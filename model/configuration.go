package model

const (
	// ConfigurationCategorySystem 系统配置
	ConfigurationCategorySystem = "system"
)

const (
	// 用户协议
	ConfigurationNameUserAgreement = "user_agreement"
	// 隐私政策
	ConfigurationNamePrivacyPolicy = "privacy_policy"
)

type Configuration struct {
	Id        int    `json:"id"`
	Category  string `json:"category"`
	Name      string `json:"name"`
	Value     string `json:"value"`
	CreatedAt string `json:"created_at" gorm:"autoCreateTime;comment:'创建时间'"` // 创建时间
	UpdatedAt string `json:"updated_at" gorm:"autoUpdateTime;comment:'更新时间'"` // 更新时间
}
