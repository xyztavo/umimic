package utils

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func BindAndValidate(r *http.Request, body interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		return err
	}
	validate := validator.New()
	return validate.Struct(body)
}
