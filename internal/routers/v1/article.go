package v1

import (
	"blog_service/pkg/app"
	"blog_service/pkg/errorcode"

	"github.com/gin-gonic/gin"
)

type Article struct {
}

func NewArticle() Article {
	return Article{}
}

func (a Article) Get(c *gin.Context) {
	app.NewResponse(c).ToErrorResponse(errorcode.ServerError)

}

func (a Article) List(c *gin.Context) {

}
func (a Article) Create(c *gin.Context) {

}
func (a Article) Update(c *gin.Context) {

}
func (a Article) Delete(c *gin.Context) {

}
