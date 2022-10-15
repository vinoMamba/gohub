package config

import "github.com/vinoMamba/gohub/pkg/config"

func init() {
	config.Add("db", func() map[string]interface{} {

		return map[string]interface{}{
			"type": config.Env("DB_TYPE", "postgres"),
			"postgres": map[string]interface{}{
				"host":     config.Env("DB_HOST", "loaclhost"),
				"port":     config.Env("DB_PORT", "5432"),
				"name":     config.Env("DB_NAME", "postgres"),
				"username": config.Env("DB_USER", "postgres"),
				"password": config.Env("DB_PASSWORD", "123456"),
			},
		}
	})
}
