package common

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"pycrs.cz/what-it-doo/internal/app/apiserver/problem"
	"pycrs.cz/what-it-doo/internal/config"
	"pycrs.cz/what-it-doo/internal/domain/model"
)

func validationErrorsToMap(ve validator.ValidationErrors) map[string]problem.FieldValidationError {
	errors := make(map[string]problem.FieldValidationError)
	for _, fe := range ve {
		field := strings.ToLower(fe.Field()) // normalize field name
		errors[field] = problem.FieldValidationError{
			Message: fmt.Sprintf("failed validation on '%s'", fe.Tag()),
			Type:    fe.Tag(),
		}
	}
	return errors
}

func DecodeValidate[T any](w http.ResponseWriter, r *http.Request) (T, bool) {
	var v T

	// Decode JSON
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&v); err != nil {
		problem.Write(w, problem.New(
			r, http.StatusBadRequest,
			"Invalid JSON",
			"Request body contains invalid JSON",
			"invalid-json",
		))
		return v, false
	}

	// Validate struct
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(v); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			errMap := validationErrorsToMap(ve)

			problem.Write(w, problem.ValidationProblemDetails{
				ProblemDetails: problem.New(
					r, http.StatusBadRequest,
					"Validation Error",
					"One or more fields are invalid",
					"validation-error",
				),
				Errors: errMap,
			})
		}
		return v, false
	}

	return v, true
}

func Encode[T any](w http.ResponseWriter, r *http.Request, status int, v T) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return fmt.Errorf("encode json: %w", err)
	}
	return nil
}

func ParseQueryInt[T int | int8 | int16 | int32 | int64](r *http.Request, name string, defaultValue T) (T, error) {
	val := r.URL.Query().Get(name)
	if val == "" {
		return defaultValue, nil
	}

	// Parse as int64 first
	i, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		var zero T
		return zero, err
	}

	return T(i), nil
}

func ParseQueryTime(r *http.Request, name string, defaultValue time.Time) (time.Time, error) {
	val := r.URL.Query().Get(name)
	if val == "" {
		return defaultValue, nil // return default if missing
	}
	parsedTime, err := time.Parse(time.RFC3339, val)
	if err != nil {
		return time.Time{}, err
	}
	return parsedTime, nil
}

func GetAvatarUrl(user model.User, config config.GravatarConfig) string {
	if user.AvatarUrl != "" {
		return user.AvatarUrl
	}

	if config.Enabled {
		hash := md5.Sum([]byte(strings.ToLower(strings.TrimSpace(user.Email))))
		return strings.NewReplacer(
			"{{hash}}", hex.EncodeToString(hash[:]),
			"{{size}}", strconv.Itoa(80),
		).Replace(config.Url)
	}

	return ""
}
