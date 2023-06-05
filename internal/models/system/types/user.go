package types

import "github.com/toma-photo/internal/models/system"

type VerifyType string

var (
	VerifyPhone   VerifyType = "phone"
	VerifyAccount VerifyType = "account"
)

type UserLoginRequest struct {
	LoginType string `json:"loginType" example:"登录类型 必填,phone/account" binding:"required"` // 手机验证码或账号密码
	Username  string `json:"username" example:"用户名"`                                       // 账户密码登录, 用户名, 用户名支持手机号密码登录
	Password  string `json:"password" example:"密码"`                                        // 账户密码登录, 密码
	Phone     string `json:"phone" example:"手机号"`                                          // 手机号验证码登录, 手机号码
	PhoneCode string `json:"phoneCode" example:"手机验证码"`                                    // 手机验证码登录, 验证码
}

type UserLoginResponse struct {
	User  system.User `json:"user"`
	Token string      `json:"token"`
}

type UserRegisterRequest struct {
	Username  string `json:"username" example:"用户名" binding:"required"`
	Password  string `json:"passWord" example:"密码" binding:"required"`
	NickName  string `json:"nickName" example:"昵称"`
	HeaderImg string `json:"headerImg" example:"头像链接"`
	Phone     string `json:"phone" example:"电话号码, 必填" binding:"required"`
	PhoneCode string `json:"phoneCode" example:"手机号验证码" binding:"required"` // 手机验证码登录, 验证码
}

type UserUpdateInfoRequest struct {
	// ID        uint   `json:"-"`
	NickName  string `json:"nickName" example:"这是昵称"`
	HeaderImg string `json:"headerImg" example:"头像链接"`
	Sex       int8   `json:"sex" example:"1"`
	City      string `json:"city" example:"四川"`
	Province  string `json:"province" example:"成都"`
	Country   string `json:"country" example:"中国"`
}

type UserInfoResponse struct {
	User system.User `json:"user"`
}

var (
	UserAccountType  = "account"
	UserPasswordType = "password"
	UserPhoneType    = "phone"
)

type ChangeUserPasswordRequest struct {
	ChangeType  string `json:"changeType" example:"修改类型, 原密码修改(password), 手机验证码修改(phone)"`
	OldPassword string `json:"oldPassword" example:"旧密码"`
	PhoneCode   string `json:"phoneCode" example:"手机号验证码"`

	NewPassword      string `json:"newPassword" example:"新密码"`
	NewPasswordAgain string `json:"newPasswordAgain" example:"新密码确认"`
}

type ChangeUserPhoneRequest struct {
	OldPhoneCode string `json:"oldPhoneCode" example:"旧手机的验证码"`
	NewPhone     string `json:"newPhone" example:"新手机号码"`
	NewPhoneCode string `json:"newPhoneCode" example:"新手机验证码"`
}

type RecoverUserPasswordRequest struct {
	Phone            string `json:"phone" example:"手机号"`
	PhoneCode        string `json:"phoneCode" example:"手机号验证码"`
	NewPassword      string `json:"newPassword" example:"新密码"`
	NewPasswordAgain string `json:"newPasswordAgain" example:"新密码确认"`
}
