package v1

import (
	"app/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func (r *Router) initSwaggerRoute(h *gin.RouterGroup) *gin.RouterGroup {
	router := h.Group("swagger")
	{
		router.GET("/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}

	return router
}

type SwaggerCfg struct {
	Version     string   `confy:"version"     yaml:"version"`
	Host        string   `confy:"host"        yaml:"host"`
	BasePath    string   `confy:"base_path"   yaml:"base_path"`
	Schemes     []string `confy:"schemes"     yaml:"schemes"`
	Title       string   `confy:"title"       yaml:"title"`
	Description string   `confy:"description" yaml:"description"`
}

func setSwaggerConfig(cfg SwaggerCfg) {
	info := docs.SwaggerInfo

	info.Version = cfg.Version
	info.Host = cfg.Host
	info.BasePath = cfg.BasePath
	info.Schemes = cfg.Schemes
	info.Title = cfg.Title
	info.Description = cfg.Description
}
