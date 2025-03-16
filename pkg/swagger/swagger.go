package swagger

import "app/docs"

type SwaggerSpec struct {
	Version     string   `yaml:"version"`
	Host        string   `yaml:"host"`
	BasePath    string   `yaml:"base_path"`
	Schemes     []string `yaml:"schemes"`
	Title       string   `yaml:"title"`
	Description string   `yaml:"description"`
}

func SetSwaggerConfig(cfg SwaggerSpec) {
	info := docs.SwaggerInfo

	info.Version = cfg.Version
	info.Host = cfg.Host
	info.BasePath = cfg.BasePath
	info.Schemes = cfg.Schemes
	info.Title = cfg.Title
	info.Description = cfg.Description
}
