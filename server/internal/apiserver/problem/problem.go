package problem

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"pycrs.cz/what-it-do/internal"
)

type ProblemDetails struct {
	Status   int    `json:"status,omitempty"`
	Title    string `json:"title,omitempty"`
	Detail   string `json:"detail,omitempty"`
	Type     string `json:"type,omitempty"`
	Instance string `json:"instance,omitempty"`
} //	@name	ProblemDetails

type FieldValidationError struct {
	Type    string `json:"type"`
	Message string `json:"message"`
} //	@name	FieldValidationError

type ValidationProblemDetails struct {
	ProblemDetails
	Errors map[string]FieldValidationError `json:"errors,omitempty"`
}

type Problem interface {
	GetStatus() int
	GetTitle() string
	GetDetail() string
	GetType() string
	GetInstance() string
}

func (p ProblemDetails) GetStatus() int      { return p.Status }
func (p ProblemDetails) GetTitle() string    { return p.Title }
func (p ProblemDetails) GetDetail() string   { return p.Detail }
func (p ProblemDetails) GetType() string     { return p.Type }
func (p ProblemDetails) GetInstance() string { return p.Instance }

func NewProblemDetails(r *http.Request, status int, title, detail, problemType string) ProblemDetails {
	return ProblemDetails{
		Status:   status,
		Title:    title,
		Detail:   detail,
		Type:     wrapType(problemType),
		Instance: r.URL.Path,
	}
}

func NewInternalServerError(r *http.Request, err error) ProblemDetails {
	return NewProblemDetails(
		r, http.StatusInternalServerError,
		"Internal Server Error",
		internal.FirstUpper(strings.TrimSpace(err.Error())),
		"internal-server-error",
	)
}

func WriteProblemDetails(w http.ResponseWriter, p Problem) error {

	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(p.GetStatus())

	return json.NewEncoder(w).Encode(p)
}

func wrapType(problem string) string {
	return fmt.Sprintf("https://wid.pycrs.cz/probs/%s", problem)
}
