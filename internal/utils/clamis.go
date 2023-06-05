package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/toma-photo/internal/global"
	"github.com/toma-photo/internal/models/system/types"
)

func GetToken(c *gin.Context) string {
	return c.Request.Header.Get("x-token")
}

func GetClaims(c *gin.Context) (*types.CustomClaims, error) {
	token := c.Request.Header.Get("x-token")
	j := NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		global.ZAP_LOG.Error("从Gin的Context中获取从jwt解析信息失败, 请检查请求头是否存在x-token且claims是否为规定结构")
	}
	return claims, err
}

// GetUserID 从Gin的Context中获取从jwt解析出来的用户ID
func GetUserID(c *gin.Context) uint {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return 0
		} else {
			return cl.BaseClaims.ID
		}
	} else {
		waitUse := claims.(*types.CustomClaims)
		return waitUse.BaseClaims.ID
	}
}

// GetUserUuid 从Gin的Context中获取从jwt解析出来的用户UUID
func GetUserUID(c *gin.Context) string {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return ""
		} else {
			return cl.UserUID
		}
	} else {
		waitUse := claims.(*types.CustomClaims)
		return waitUse.UserUID
	}
}

// GetUserInfo 从Gin的Context中获取从jwt解析出来的用户角色id
func GetUserInfo(c *gin.Context) *types.CustomClaims {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return nil
		} else {
			return cl
		}
	} else {
		waitUse := claims.(*types.CustomClaims)
		return waitUse
	}
}
