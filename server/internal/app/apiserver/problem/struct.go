package problem

type ProblemDetails struct {
	Status   int    `json:"status" validate:"required"`
	Title    string `json:"title" validate:"required"`
	Detail   string `json:"detail" validate:"required"`
	Type     string `json:"type" validate:"required"`
	Instance string `json:"instance" validate:"required"`
} //	@name	ProblemDetails

type FieldValidationError struct {
	Type    string `json:"type" validate:"required"`
	Message string `json:"message" validate:"required"`
} //	@name	FieldValidationError

type ValidationProblemDetails struct {
	ProblemDetails
	Errors map[string]FieldValidationError `json:"errors" validate:"required"`
}
