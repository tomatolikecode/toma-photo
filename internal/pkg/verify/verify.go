package verify

var (
	AccountLogin         = Rules{"UserName": {NotEmpty()}, "Password": {NotEmpty()}}
	PhoneLogin           = Rules{"Phone": {NotEmpty()}, "PhoneCode": {NotEmpty()}}
	UserRegister         = Rules{"Username": {NotEmpty()}, "Password": {NotEmpty()}, "Phone": {NotEmpty()}, "PhoneCode": {NotEmpty()}}
	UserChangePwdByPwd   = Rules{"OldPassword": {NotEmpty()}, "NewPassword": {NotEmpty()}, "NewPasswordAgain": {NotEmpty()}}
	UserChangePwdByPhone = Rules{"PhoneCode": {NotEmpty()}, "NewPassword": {NotEmpty()}, "NewPasswordAgain": {NotEmpty()}}
	UserChangePhone      = Rules{"OldPhoneCode": {NotEmpty()}, "NewPhone": {NotEmpty()}, "NewPhoneCode": {NotEmpty()}}
	RecoverUserPassword  = Rules{"Phone": {NotEmpty()}, "PhoneCode": {NotEmpty()}, "NewPassword": {NotEmpty()}, "NewPasswordAgain": {NotEmpty()}}
)
