package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5"
	"pycrs.cz/what-it-do/internal/apiserver"
)

func DbConfigStructLevelValidation(sl validator.StructLevel) {
	cfg := sl.Current().Interface().(apiserver.DatabaseConfiguration)

	if cfg.URL == "" {
		if cfg.Host == "" {
			sl.ReportError(cfg.Host, "host", "Host", "required_without_url", "")
		}
		if cfg.User == "" {
			sl.ReportError(cfg.User, "user", "User", "required_without_url", "")
		}
		if cfg.Password == "" {
			sl.ReportError(cfg.Password, "password", "Password", "required_without_url", "")
		}
		if cfg.Name == "" {
			sl.ReportError(cfg.Name, "name", "Name", "required_without_url", "")
		}
		if cfg.Port == 0 {
			sl.ReportError(cfg.Port, "port", "Port", "required_without_url", "")
		}
	} else {
		if _, err := pgx.ParseConfig(cfg.URL); err != nil {
			sl.ReportError(cfg.URL, "url", "URL", "invalid_postgres_url", err.Error())
		}
	}
}
