// Package swag provides ...
package swag

import (
	_ "github.com/eiblog/eiblog/v2/pkg/core/blog/docs" // docs

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// RegisterRoutes register routes
func RegisterRoutes(group gin.IRoutes) {
	group.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}