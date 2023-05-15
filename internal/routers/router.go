package routers

import (
	"blog_service/global"
	"blog_service/internal/middleware"
	v1 "blog_service/internal/routers/v1"
	"net/http"

	_ "blog_service/docs"

	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Translations())
	url := ginSwagger.URL("http://127.0.0.1:8000/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, url))
	r.GET("/auth", v1.GetAuth)

	upload := NewUpload()
	r.POST("/upload/file", upload.UploadFile)

	//訪問http://127.0.0.1:8000/static/ 查看所有資源
	//長傳加密後檔名
	//curl -X POST http://127.0.0.1:8000/upload/file \
	//-F file=@/Users/terry_hsiesh/Side\ Project/blog-service/storage/uploads/go.png \
	//-F type=1
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))

	tag := v1.NewTag()
	article := v1.NewArticle()
	apiv1 := r.Group("/api/v1")
	//jwt 驗證中介
	apiv1.Use(middleware.JWT())
	{
		apiv1.POST("/tags", tag.Create)
		apiv1.DELETE("/tags/:id", tag.Delete)
		apiv1.PUT("/tags/:id", tag.Update)
		apiv1.GET("/tags/:id", tag.Get)
		apiv1.GET("/tags", tag.List)

		apiv1.POST("/articles", article.Create)
		apiv1.DELETE("/articles/:id", article.Delete)
		apiv1.PUT("/articles/:id", article.Update)
		apiv1.GET("/articles/:id", article.Get)
		apiv1.GET("/articles", article.List)

	}

	return r
}
