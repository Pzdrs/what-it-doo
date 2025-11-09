package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"pycrs.cz/what-it-doo/internal/apiserver/problem"
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

func DecodeAndValidate[T any](w http.ResponseWriter, r *http.Request) (*T, bool) {
	var req T

	// Decode JSON
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		problem.Write(w, problem.New(
			r, http.StatusBadRequest,
			"Invalid JSON",
			"Request body contains invalid JSON",
			"invalid-json",
		))
		return nil, false
	}

	// Validate struct
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(req); err != nil {
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
		return nil, false
	}

	return &req, true
}

func WriteJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
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
