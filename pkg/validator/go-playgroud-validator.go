package validator

import (
	"encoding/json"
	"net/http"
	"reflect"
	"strings"
	"time"

	v10 "github.com/go-playground/validator/v10"
	"github.com/patrickchagastavares/rinha-backend/internal/entities"
)

const (
	InvalidPayload = "invalid_payload"
)

type validator struct {
	validate *v10.Validate
}

func New() Validator {
	validate := v10.New()

	// Registre uma função de validação personalizada para o formato de data.
	validate.RegisterValidation("dateformat", func(fl v10.FieldLevel) bool {
		dateStr := fl.Field().String()

		_, err := time.Parse("2006-01-02", dateStr)
		return err == nil
	})

	return &validator{
		validate: validate,
	}
}

// Validate is used to validate a struct using rules defined in the 'validate' tag.
// This method return the struct *ValidationError that contains details of the rules violation.
// ValidationError is compatible with 'error' interface and can be returned as error.
func (v *validator) Validate(val any) *ValidationError {
	err := v.validate.Struct(val)
	if err == nil {
		return nil
	}

	if violationEntries, ok := err.(v10.ValidationErrors); ok {
		violations := make([]Violation, len(violationEntries))
		for i, err := range violationEntries {
			violations[i] = Violation{
				Namespace: err.Namespace(),
				Field:     err.Field(),
				FieldJSON: getJSONTag(val, err.StructField()),
				Tag:       err.Tag(),
				Value:     err.Value(),
			}
		}
		return &ValidationError{
			OriginalMessage: err.Error(),
			Message:         InvalidPayload,
			Violations:      violations,
		}
	}
	return &ValidationError{
		OriginalMessage: err.Error(),
		Message:         err.Error(),
	}
}

// Error return error data as json string
func (v *ValidationError) Error() string {
	dt, err := json.Marshal(v)
	if err != nil {
		return err.Error()
	}
	return string(dt)
}

func (v *ValidationError) ToHttpErr() error {
	return entities.NewHttpErr(http.StatusBadRequest, v.Message, v.Violations)
}

// getJSONTag is user to get a json field name using reflection.
func getJSONTag(val any, fieldName string) string {
	field, ok := reflect.TypeOf(val).FieldByName(fieldName)
	if !ok {
		return ""
	}

	jsonTag, hasJsonTag := field.Tag.Lookup("json")
	if !hasJsonTag {
		return fieldName
	}

	return strings.Split(jsonTag, ",")[0]
}
