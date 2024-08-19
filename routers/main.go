package routers

import "github.com/gin-gonic/gin"

func RouterCombine(r *gin.Engine) {
	UserRouter(r.Group("/users"))
	AuthRouter(r.Group("/auth"))
	EventRouter(r.Group("/event"))
	ProfileRouter(r.Group("/profile"))
	CategoriesRouter(r.Group("/categories"))
	TransactionRouter(r.Group("/transaction"))
	PartnerRouter(r.Group("/partner"))
	LocationRouter(r.Group("/location"))
}