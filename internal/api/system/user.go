package system

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/toma-photo/internal/global"
	"github.com/toma-photo/internal/models/common/response"
	"github.com/toma-photo/internal/models/system"
	"github.com/toma-photo/internal/models/system/types"

	"github.com/toma-photo/internal/pkg/verify"
	"github.com/toma-photo/internal/utils"
	"go.uber.org/zap"
)

type UserApi struct{}

// @Tags 用户模块
// @Summary 用户登录
// @Description 用户登录, 账号密码登录, 手机验证码注册登录
// @Produce   application/json
// @Param    data  body  types.UserLoginRequest true  "登录类型, 用户名, 密码, 手机号, 验证码"
// @Success 200 {object} response.Response{data=types.UserLoginResponse,msg=string,code=int} "返回包括用户信息,token,过期时间"
// @Router /user/login [post]
func (u *UserApi) UserLogin(c *gin.Context) {
	user := new(system.User)
	var err error

	var loginParam types.UserLoginRequest
	if err := c.ShouldBindJSON(&loginParam); err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}

	switch loginParam.LoginType {
	case string(types.VerifyAccount):
		if err := verify.ParamVerify(loginParam, verify.AccountLogin); err != nil {
			response.FailWithMsg(err.Error(), c)
			return
		}

		user, err = userServce.UserLogin(&system.User{Username: loginParam.Username, Password: loginParam.Password})
		if err != nil {
			response.FailWithMsg(err.Error(), c)
			return
		}
	case string(types.VerifyPhone):
		// TODO: 手机验证码登录
	default:
		response.FailWithMsg("错误登录类型", c)
		return
	}

	// 初始化用户信息
	user.InitUserInfo()

	// 生成token
	var token string
	if token = creatToken(*user); token == "" {
		response.FailWithMsg("获取Token失败", c)
		return
	}

	c.Header("x-token", token)

	response.OkWithDetail(types.UserLoginResponse{User: *user, Token: token}, "登录成功", c)
}

func creatToken(user system.User) string {
	j := &utils.JWT{SigningKey: []byte(global.CONFIG.JWT.SigningKey)} // 唯一签名
	claims := j.CreateClaims(types.BaseClaims{
		ID:       user.ID,
		UserUID:  user.UserUID,
		NickName: user.NickName,
		Username: user.Username,
		Phone:    user.Phone,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		global.ZAP_LOG.Error("生成Token错误", zap.Error(err))
		return ""
	}
	return token
}

// @Tags 用户模块
// @Summary 用户注册
// @Description
// @Produce   application/json
// @Param    data  body  types.UserRegisterRequest  true  "注册信息"
// @Success 200 {object} response.Response{msg=string,code=int}
// @Router /user/register [post]
func (u *UserApi) UserRegister(c *gin.Context) {
	var registerParam types.UserRegisterRequest
	if err := c.ShouldBindJSON(&registerParam); err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}

	// 验证基本注册信息的有效性
	if err := verify.ParamVerify(registerParam, verify.UserRegister); err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}
	// TODO: 校验验证码的有效性

	registerUser := system.User{
		Username:  registerParam.Username,
		Password:  registerParam.Password,
		NickName:  registerParam.NickName,
		HeaderImg: registerParam.HeaderImg,
		Phone:     registerParam.Phone,
	}
	_, err := userServce.UserRegister(registerUser)
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}
	response.OkWithMsg("注册成功", c)
}

// @Tags 用户模块
// @Summary 修改密码
// @Description 修改用户密码, 通过手机号验证码修改, 通过原密码新密码验证修改
// @Accept    application/json
// @Produce   application/json
// @Param    data  body  types.ChangeUserPasswordRequest  true  "修改密码"
// @Success 200 {object} response.Response{code=int,msg=string}
// @Router /user/changepwd [put]
func (u *UserApi) ChangeUserPassword(c *gin.Context) {
	userUID := utils.GetUserUID(c)

	var param types.ChangeUserPasswordRequest
	if err := c.ShouldBindJSON(&param); err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}

	switch strings.ToLower(param.ChangeType) {
	case types.UserPasswordType:
		if err := verify.ParamVerify(param, verify.UserChangePwdByPwd); err != nil {
			response.FailWithMsg(err.Error(), c)
			return
		}
		if param.NewPassword != param.NewPasswordAgain {
			response.FailWithMsg("两次密码不一致", c)
			return
		}

		if err := userServce.ChangeUserPasswordByPwd(userUID, param.OldPassword, param.NewPassword); err != nil {
			response.FailWithMsg(err.Error(), c)
			return
		}
	case types.UserPhoneType:
		if err := verify.ParamVerify(param, verify.UserChangePwdByPhone); err != nil {
			response.FailWithMsg(err.Error(), c)
			return
		}
		// TODO: 验证码验证

		// 跟新密码
		if err := userServce.ChangeUserPassword(userUID, param.NewPassword); err != nil {
			response.FailWithMsg(err.Error(), c)
			return
		}
	default:
		response.FailWithMsg("错误登录类型", c)
		return
	}

	response.OkWithMsg("修改成功", c)
}

// @Tags 用户模块
// @Summary 更换手机号
// @Description 更换用户手机号, 原手机号验证修改
// @Accept    application/json
// @Produce   application/json
// @Param    data  body  types.ChangeUserPhoneRequest  true  "更换手机号"
// @Success 200 {object} response.Response{code=int,msg=string}
// @Router /user/changephone [put]
func (u *UserApi) ChangeUserPhone(c *gin.Context) {
	userUID := utils.GetUserUID(c)

	var param types.ChangeUserPhoneRequest
	if err := c.ShouldBindJSON(&param); err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}

	if err := verify.ParamVerify(param, verify.UserChangePhone); err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}

	// TODO: 旧手机验证码验证

	// TODO: 新手机验证码验证

	// 更新手机号
	if err := userServce.ChangeUserPhone(userUID, param.NewPhone); err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}
	response.OkWithMsg("修改成功", c)
}

// @Tags 用户模块
// @Summary 找回密码
// @Description 找回密码
// @Accept    application/json
// @Produce   application/json
// @Param    data  body  types.RecoverUserPasswordRequest  true  "找回密码"
// @Success 200 {object} response.Response{code=int,msg=string}
// @Router /user/recoverpwd [put]
func (u *UserApi) RecoverUserPassword(c *gin.Context) {
	var param types.RecoverUserPasswordRequest
	if err := c.ShouldBindJSON(&param); err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}

	if err := verify.ParamVerify(param, verify.RecoverUserPassword); err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}

	if param.NewPassword != param.NewPasswordAgain {
		response.FailWithMsg("两次密码不一致", c)
		return
	}

	// TODO: 手机号验证码验证

	// 更新密码
	if err := userServce.ChangeUserPasswordByPhone(param.Phone, param.NewPassword); err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}

	response.OkWithMsg("找回成功", c)
}

// @Tags 用户模块
// @Summary 更新用户信息
// @Description 更新用户信息
// @Accept    application/json
// @Produce   application/json
// @Param    data  body  types.UserUpdateInfoRequest  true  "用户更新信息"
// @Success 200 {object} response.Response{code=int,msg=string}
// @Router /user/info [put]
func (u *UserApi) UpdateUserInfo(c *gin.Context) {
	userID := utils.GetUserID(c)
	if userID == 0 {
		response.FailWithMsg("请求数据错误", c)
		return
	}

	var updateParam types.UserUpdateInfoRequest
	if err := c.ShouldBindJSON(&updateParam); err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}

	userUpdate := make(map[string]interface{})
	userUpdate["nick_name"] = updateParam.NickName
	userUpdate["header_img"] = updateParam.HeaderImg
	userUpdate["sex"] = updateParam.Sex
	userUpdate["city"] = updateParam.City
	userUpdate["province"] = updateParam.Province
	userUpdate["country"] = updateParam.Country

	err := userServce.UpdateUserInfo(userID, userUpdate)
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}

	response.OkWithMsg("更新成功", c)
}

// @Tags 用户模块
// @Summary 获取用户信息
// @Description 获取用户信息
// @Accept    application/json
// @Produce   application/json
// @Success 200 {object} response.Response{code=int,msg=string,data=types.UserInfoResponse}
// @Router /user/info [get]
func (u *UserApi) GetUserInfo(c *gin.Context) {
	userUID := utils.GetUserUID(c)
	if userUID == "" {
		response.FailWithMsg("请求数据错误", c)
		return
	}

	user, err := userServce.GetUserInfo(userUID)
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}

	response.OkWithDetail(types.UserInfoResponse{User: *user}, "查询成功", c)
}

// @Tags 用户模块
// @Summary 用户退出登录
// @Description 退出登录
// @accept    application/json
// @Produce   application/json
// @Success 200 {object} response.Response{code=int,msg=string}
// @Router /user/logout [delete]
func (u *UserApi) UserLogout(c *gin.Context) {
	token := utils.GetToken(c)
	if token == "" {
		response.FailWithMsg("退出登录失败", c)
		return
	}

	// TODO: 将用户token放入到黑名单中

	response.OkWithMsg("退出登录成功", c)
}

// @Tags 用户模块
// @Summary 用户注销
// @Description 注销
// @accept    application/json
// @Produce   application/json
// @Success 200 {object} response.Response{code=int,msg=string}
// @Router /user/logoff [delete]
func (u *UserApi) UserLogoff(c *gin.Context) {
	// TODO: 将用户token放入到黑名单中
	// 删除用户信息
	claims, found := c.Get("claims")
	if !found {
		response.FailWithMsg("无效用户", c)
		return
	}
	custom := claims.(types.CustomClaims)
	userServce.UserLogout(custom.UserUID)
	response.OkWithMsg("注销成功", c)
}
