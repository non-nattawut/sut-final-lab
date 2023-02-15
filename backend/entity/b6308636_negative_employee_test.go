package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
)

func TestNegativeEmployeeID(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	t.Run("Negative Employee ID Case 1", func(t *testing.T) {
		e := Employee{
			Name:       "nattawut Rodthong",
			Email:      "nattavutnon@gmail.com",
			EmployeeID: "12345678",
		}

		ok, err := govalidator.ValidateStruct(e)

		g.Expect(ok).NotTo(gomega.BeTrue())

		g.Expect(err).NotTo(gomega.BeNil())

		g.Expect(err.Error()).To(gomega.Equal("Bad employee ID format"))
	})

	t.Run("Negative Employee ID Case 2", func(t *testing.T) {
		e := Employee{
			Name:       "nattawut Rodthong",
			Email:      "nattavutnon@gmail.com",
			EmployeeID: "A12345678",
		}

		ok, err := govalidator.ValidateStruct(e)

		g.Expect(ok).NotTo(gomega.BeTrue())

		g.Expect(err).NotTo(gomega.BeNil())

		g.Expect(err.Error()).To(gomega.Equal("Bad employee ID format"))
	})

	t.Run("Negative Employee ID Case 3", func(t *testing.T) {
		e := Employee{
			Name:       "nattawut Rodthong",
			Email:      "nattavutnon@gmail.com",
			EmployeeID: "M12345678123123",
		}

		ok, err := govalidator.ValidateStruct(e)

		g.Expect(ok).NotTo(gomega.BeTrue())

		g.Expect(err).NotTo(gomega.BeNil())

		g.Expect(err.Error()).To(gomega.Equal("Bad employee ID format"))
	})

	t.Run("Negative Employee ID Case 4", func(t *testing.T) {
		e := Employee{
			Name:       "nattawut Rodthong",
			Email:      "nattavutnon@gmail.com",
			EmployeeID: "S12",
		}

		ok, err := govalidator.ValidateStruct(e)

		g.Expect(ok).NotTo(gomega.BeTrue())

		g.Expect(err).NotTo(gomega.BeNil())

		g.Expect(err.Error()).To(gomega.Equal("Bad employee ID format"))
	})

	t.Run("Negative Employee ID Case 5", func(t *testing.T) {
		e := Employee{
			Name:       "nattawut Rodthong",
			Email:      "nattavutnon@gmail.com",
			EmployeeID: "S12345678SD",
		}

		ok, err := govalidator.ValidateStruct(e)

		g.Expect(ok).NotTo(gomega.BeTrue())

		g.Expect(err).NotTo(gomega.BeNil())

		g.Expect(err.Error()).To(gomega.Equal("Bad employee ID format"))
	})

	t.Run("Negative Employee ID Case 6", func(t *testing.T) {
		e := Employee{
			Name:       "nattawut Rodthong",
			Email:      "nattavutnon@gmail.com",
			EmployeeID: "S12345s67",
		}

		ok, err := govalidator.ValidateStruct(e)

		g.Expect(ok).NotTo(gomega.BeTrue())

		g.Expect(err).NotTo(gomega.BeNil())

		g.Expect(err.Error()).To(gomega.Equal("Bad employee ID format"))
	})

	t.Run("Negative Employee ID Case 7", func(t *testing.T) {
		e := Employee{
			Name:       "nattawut Rodthong",
			Email:      "nattavutnon@gmail.com",
			EmployeeID: "S12sssss345ssss67",
		}

		ok, err := govalidator.ValidateStruct(e)

		g.Expect(ok).NotTo(gomega.BeTrue())

		g.Expect(err).NotTo(gomega.BeNil())

		g.Expect(err.Error()).To(gomega.Equal("Bad employee ID format"))
	})
}
