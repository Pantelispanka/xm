package domain

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator"
)

type CompanyType int

const (
	COORPORATIONS CompanyType = iota
	NONPROFIT
	COOPERATIVE
	SOLEPROPRIETORSHIP
)

type Company struct {
	ID          string `json:"id" validate:"required" bson:"_id"`
	Name        string `json:"name"  validate:"required,min=5,max=15" bson:"name"`
	Description string `json:"description" validate:"min=0,max=3000" bson:"description"`
	Employees   int64  `json:"employees" validate:"required" bson:"employees"`
	Registered  bool   `json:"registered" validate:"required" bson:"registered"`
	CompanyType string `json:"company_type" validate:"required,oneof=coorporation nonprofit cooperative sole" bson:"company_type"`
}

// If the validation fails and error will be returned with the details of the error.
func (p *Company) Validate() error {
	validate := validator.New()
	err := validate.Struct(p)
	if err != nil {
		var validationErrors []string
		for _, vErr := range err.(validator.ValidationErrors) {
			errMessage := fmt.Sprintf("'%s' has a value of '%v' which does not satisfy '%s'.\n", vErr.Field(), vErr.Value(), vErr.Tag())
			validationErrors = append(validationErrors, errMessage)
		}
		return fmt.Errorf("error in struct: %+v. details: '%v'", *p, strings.Join(validationErrors, ", "))
	}
	return nil
}
