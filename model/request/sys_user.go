package request

// User register structure
type Register struct {
	Username string `json:"username"`
	Password string `json:"password"`

	AuthorityId string `json:"authority_id" gorm:"default:888"`
}

// User login structure
type Login struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Captcha   string `json:"captcha"`
	CaptchaId string `json:"captchaId"`
}

// Modify password structure
type ChangePasswordStruct struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	NewPassword string `json:"new_password"`
}

// Modify  user's auth structure
type SetUserAuth struct {
	Id          float64 `json:"id" form:"id"`
	AuthorityId string  `json:"authority_id"`
}

// Modify  user's auth structure
type SetAdminInfo struct {
	Id        float64 `json:"id" form:"id"`
	Email     string  `json:"email" gorm:"default:''"`
	Phone     string  `json:"phone" gorm:"default:''"`
	Name      string  `json:"nickName" gorm:"default:'QMPlusUser'"`
	HeaderImg string  `json:"headerImg" gorm:"default:'http://www.henrongyi.top/avatar/lufu.jpg'"`
}
