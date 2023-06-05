package system

import "github.com/gin-gonic/gin"

type BasicApi struct{}

// @Tags 基本模块
// @Summary 发送短信验证码
// @Description 短信验证码
// @Success 200 {string} health
// @Router /sms [get]
func (b *BasicApi) Sms(c *gin.Context) {
	// 不同类型发送短信,  对于手机号注册更滑你手机号,要验证手机号是否已经存在
}
