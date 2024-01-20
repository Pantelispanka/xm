package domain_test

import (
	"testing"
	"xm-challenge/internal/domain"

	"github.com/stretchr/testify/assert"
)

func TestCompanyValidation(t *testing.T) {

	t.Run("non valid", func(t *testing.T) {

		errMessage := "''Name' has a value of 'Aadsfasdfasdfasdfasdfas' which does not satisfy"
		port := domain.Company{
			ID:          "AEAUH",
			Name:        "Aadsfasdfasdfasdfasdfas",
			Description: "asfdasdf",
			Employees:   5,
			Registered:  true,
			CompanyType: "coorporation",
		}
		err := port.Validate()
		assert.ErrorContains(t, err, errMessage)
	})

	t.Run("valid", func(t *testing.T) {

		port := domain.Company{
			ID:          "AEAUH",
			Name:        "asdfg",
			Description: "asfdasdf",
			Employees:   5,
			Registered:  true,
			CompanyType: "coorporation",
		}
		err := port.Validate()
		assert.NoError(t, err)
	})

}
