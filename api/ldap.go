package api

import (
	"github.com/gin-gonic/gin"
	"watchAlert/internal/middleware"
	"watchAlert/internal/models"
	"watchAlert/internal/services"
)

type LdapController struct{}

func (lc LdapController) API(gin *gin.RouterGroup) {
	ldapA := gin.Group("ldap")
	ldapA.Use(
		middleware.Auth(),
		middleware.Permission(),
		middleware.ParseTenant(),
		middleware.AuditingLog(),
	)
	{
		ldapA.POST("ldapCreate", lc.Create)
		ldapA.POST("ldapUpdate", lc.Update)
	}
	ldapB := gin.Group("ldap")
	ldapB.Use(
		middleware.Auth(),
		middleware.Permission(),
		middleware.ParseTenant(),
	)
	{
		ldapB.GET("ldapInfo", lc.Get)
		ldapB.POST("ldapPing", lc.Ping)
	}
}

// 获取ldap信息

func (lc LdapController) Get(ctx *gin.Context) {
	Service(ctx, func() (interface{}, interface{}) {
		return services.LdapService.Get()
	})
}

// 创建ldap

func (lc LdapController) Create(ctx *gin.Context) {
	r := new(models.Ldap)
	BindJson(ctx, r)
	Service(ctx, func() (interface{}, interface{}) {
		return services.LdapService.Create(r)
	})
}

// 更新ldap

func (lc LdapController) Update(ctx *gin.Context) {
	r := new(models.Ldap)
	BindJson(ctx, r)
	Service(ctx, func() (interface{}, interface{}) {
		return services.LdapService.Update(r)
	})
}

// ldap测试

func (lc LdapController) Ping(ctx *gin.Context) {
	r := new(models.Ldap)
	BindJson(ctx, r)
	Service(ctx, func() (interface{}, interface{}) {
		return services.LdapService.Ping(r)
	})
}

// 登录

func (lc LdapController) Login(ctx *gin.Context) {
	r := new(models.Member)
	BindJson(ctx, r)
	Service(ctx, func() (interface{}, interface{}) {
		return services.LdapService.Login(r)
	})
}
