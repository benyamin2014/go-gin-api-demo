package route

import (
	"github.com/gin-gonic/gin"
	"jtapi/api/test"
	"net/http"
)
// 路由初始化
// 引入路由
func Setup() *gin.Engine {
	engin := gin.Default()

	engin.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "项目已启动")
	})

	//定义路由组 /test post/put/get/delete
	//[GIN-debug] GET    /                         --> jtapi/route.Setup.func1 (3 handlers)
	//[GIN-debug] POST   /test                     --> jtapi/api/test.Add (3 handlers)
	//[GIN-debug] PUT    /test/:id                 --> jtapi/api/test.Update (3 handlers)
	//[GIN-debug] GET    /test/:id                 --> jtapi/api/test.Get (3 handlers)
	//[GIN-debug] DELETE /test/:id                 --> jtapi/api/test.Delete (3 handlers)
	//[GIN-debug] GET    /tests                    --> jtapi/api/test.Page (3 handlers)
	TestRoute := engin.Group("/test")
	TestRoute.POST("", test.Add)
	TestRoute.PUT("/:id", test.Update)
	TestRoute.GET("/:id", test.Get)
	TestRoute.DELETE("/:id", test.Delete)
	// 列表查
	engin.GET("/tests", test.Page)

	return engin
}
