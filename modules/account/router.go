package account

import (
	"github.com/ahmad20/bri-mini-project/modules/auth"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine, h HandlerInterface, loginM, statusM gin.HandlerFunc) {
	r.POST("/login", h.Login)
	r.POST("/admin", h.AdminRegister)

	// Login Account
	LoginGroup := r.Group("/admin", loginM)
	{
		// Login and role:admin, superadmin
		AllRoleGroup := LoginGroup.Group("/",
			auth.RoleAuthorization("admin", "superadmin"), statusM)
		{
			AllRoleGroup.GET(":id", h.ReadById)
			AllRoleGroup.GET("admins", h.GetAllAdmin)
			AllRoleGroup.GET("admins?search={keyword}&page={page}&limit={limit}", h.GetAllAdmin)
			AllRoleGroup.GET("approvals", h.GetAdminApproval)

			AllRoleGroup.POST("customer", h.CustomerRegister)
			AllRoleGroup.GET("customers", h.GetAllCustomer)
			AllRoleGroup.GET("customers?search={keyword}&page={page}&limit={limit}", h.GetAllCustomer)
			AllRoleGroup.DELETE("customer/:id", h.CustomerDelete)
		}
		// Login and role:superadmin
		SuperAdminGroup := LoginGroup.Group("/", auth.RoleAuthorization("superadmin"))
		{
			SuperAdminGroup.PATCH(":id/approval", h.UpdateAdminApproval)
			SuperAdminGroup.DELETE(":id", h.AdminDelete)
			SuperAdminGroup.PATCH(":id/status", h.UpdateAdminStatus)
		}

	}
}
