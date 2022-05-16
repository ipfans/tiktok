package main

import (
	"github.com/ipfans/tiktok/cmd/api"
	_ "github.com/ipfans/tiktok/cmd/docs"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title Swagger Product API
// @version v0.0.1
// @description This is product api swagger doc.
// @contact.name yan.jinbin
// @contact.email yan.jinbin@shipper.id
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host 127.0.0.1:10000
// @BasePath /v1
func main() {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/v1")
	{
		v1.GET("/product/category", api.GetCategoryTree)
		v1.GET("/product/band", api.GetBrandList)
		v1.GET("/product/attribute", api.GetCategoryAttrs)
		v1.POST("/product/add", api.CreateProduct)
		v1.PATCH("/product/edit", api.EditProductAttributes)
		v1.PUT("/product/price", api.UpdateProductPrice)
		v1.PUT("/product/stock", api.UpdateProductStock)
		v1.DELETE("/product/delete", api.DeleteProduct)
		v1.GET("/product/info", api.QueryProductInfoList)
		v1.PUT("/product/shelf/up-down", api.UpDown)
	}

	r.Run(":10000")
}
