package utils_test

import (
	"testing"

	"github.com/user0608/mujeresapi/utils"
)

func TestDecimalValidator(t *testing.T) {
	datos := []struct {
		name string
		got  string
		want error
	}{
		{name: "valid1", got: "-2832.3232423", want: nil},
		{name: "valid2", got: "22.322423", want: nil},
		{name: "valid3", got: "1", want: nil},
		{name: "valid4", got: "-1", want: nil},
		{name: "invalid1", got: "-.3232423", want: utils.ErrInvalidData},
		{name: "invalid2", got: "....", want: utils.ErrInvalidData},
		{name: "invalid3", got: "----3.423", want: utils.ErrInvalidData},
		{name: "invalid4", got: "-323.2.4.23", want: utils.ErrInvalidData},
		{name: "invalid5", got: "-323-232423", want: utils.ErrInvalidData},
	}

	for _, d := range datos {
		t.Run(d.name, func(t *testing.T) {
			if err := utils.ValidarDecimal(d.got); err != d.want {
				t.Errorf("Se esperaba un error: %v", d.want)
			}
		})
	}
}
func BenchmarkDecimalValid(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = utils.ValidarDecimal("-2832.3232423")
	}
}
