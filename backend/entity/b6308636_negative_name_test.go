package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
)

func TestNegativeName(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	t.Run("Negative Name Case", func(t *testing.T) {
		e := Employee{
			Name:       "",
			Email:      "nattavutnon@gmail.com",
			EmployeeID: "J12345678",
		}

		ok, err := govalidator.ValidateStruct(e)

		g.Expect(ok).NotTo(gomega.BeTrue())

		g.Expect(err).NotTo(gomega.BeNil())

		g.Expect(err.Error()).To(gomega.Equal("Name cannot Blank"))
	})
}
