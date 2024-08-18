package model

type WxUser struct {
	ID         int    `gorm:"primary_key" json:"id"`
	Uuid       string `gorm:"type:char(36);unique" json:"uuid"`           // 用户唯一标识
	Openid     string `gorm:"type:varchar(100);index" json:"openid"`      // openid
	Unionid    string `gorm:"type:varchar(100);index" json:"unionid"`     // unionid
	Password   string `gorm:"type:varchar(100)" json:"password"`          // 密码
	Phone      string `gorm:"type:varchar(20)" json:"phone"`              // 手机号
	Email      string `gorm:"type:varchar(100)" json:"email"`             // 邮箱
	Name       string `gorm:"type:varchar(50)" json:"name"`               // 姓名
	Avatar     string `gorm:"type:varchar(200)" json:"avatarUrl"`         // 头像
	NickName   string `gorm:"type:varchar(50)" json:"nickName"`           // 昵称
	City       string `gorm:"type:varchar(50)" json:"city"`               // 城市
	Status     int    `gorm:"type:int" json:"status"`                     // 状态  1:启用 2:禁用: 3:未激活
	Gender     int    `gorm:"type:int" json:"gender"`                     // 性别 0:未知 1:男 2:女
	CreatedAt  string `gorm:"autoCreateTime" json:"created_at"`           // 创建时间
	UpdatedAt  string `gorm:"autoUpdateTime" json:"updated_at"`           // 更新时间
	IsDeleted  int    `gorm:"type:int" json:"is_deleted"`                 // 是否删除 1:删除 0:未删除
	InviteCode string `gorm:"type:varchar(100);idnex" json:"invite_code"` // 邀请码
	// 证件类型
	CertificateType string `gorm:"type:varchar(100)" json:"certificate_type"`
	// 证件号码
	CertificateNumber string `gorm:"type:varchar(100)" json:"certificate_number"`
	// 证件图片
	CertificateImage string `gorm:"type:varchar(200)" json:"certificate_image"`

	// 是否实名认证
	IsRealName int `gorm:"type:int" json:"is_real_name"`
}

type WxUserRes struct {
	WxUser
	Role             string `json:"role"`
	TeamMemberStatus string `json:"team_member_status"`
}

type Watermark struct {
	AppID     string `json:"appid"`
	TimeStamp int64  `json:"timestamp"`
}

type WXUserInfo struct {
	OpenID    string    `json:"openId,omitempty"`
	NickName  string    `json:"nickName"`
	AvatarUrl string    `json:"avatarUrl"`
	Gender    int       `json:"gender"`
	Country   string    `json:"country"`
	Province  string    `json:"province"`
	City      string    `json:"city"`
	UnionID   string    `json:"unionId,omitempty"`
	Language  string    `json:"language"`
	Watermark Watermark `json:"watermark,omitempty"`
}

type ReqWXUserInfo struct {
	UserInfo      WXUserInfo `json:"userInfo"`
	RawData       string     `json:"rawData"`
	Signature     string     `json:"signature"`
	EncryptedData string     `json:"encryptedData"`
	IV            string     `json:"iv"`
}

type AuthLoginBody struct {
	Code     string        `json:"code"`
	UserInfo ReqWXUserInfo `json:"userInfo"`
}
