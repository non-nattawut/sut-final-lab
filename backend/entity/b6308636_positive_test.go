package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
)

func TestPositive(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	t.Run("Positive Case", func(t *testing.T) {
		e := Employee{
			Name:       "Nattawut Rodthong",
			Email:      "nattavutnon@gmail.com",
			EmployeeID: "J12345678",
		}

		ok, err := govalidator.ValidateStruct(e)

		g.Expect(ok).To(gomega.BeTrue())

		g.Expect(err).To(gomega.BeNil())
	})
}
