package problem

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ProblemDetails struct {
	Status   int    `json:"status,omitempty"`
	Title    string `json:"title,omitempty"`
	Detail   string `json:"detail,omitempty"`
	Type     string `json:"type,omitempty"`
	Instance string `json:"instance,omitempty"`
} //	@name	ProblemDetails

type ValidationError struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

type ValidationProblemDetails struct {
	ProblemDetails
	Errors map[string][]ValidationError `json:"errors,omitempty"`
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

func WriteProblemDetails(w http.ResponseWriter, p Problem) error {

	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(p.GetStatus())

	return json.NewEncoder(w).Encode(p)
}

func getDefaultType(statusCode int) string {
	return fmt.Sprintf("https://httpstatuses.io/%d", statusCode)
}
