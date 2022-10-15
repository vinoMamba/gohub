package config

import "github.com/vinoMamba/gohub/pkg/config"

func init() {
	config.Add("app", func() map[string]interface{} {
		return map[string]interface{}{
			"name": config.Env("APP_NAME", "gohub"),
			"port": config.Env("APP_PORT", "3000"),
		}
	})
}
