package validator

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestHandler(t *testing.T) {
	t.Run("Get Message", func(t *testing.T) {
		expect := []string{
			"Reは必須です。",
		}

		type validationTarget struct {
			Re interface{} `validate:"required"`
		}
		target := validationTarget{}

		validation := Validate(&target)
		result := validation.GetMessages()

		if diff := cmp.Diff(expect, result); diff != "" {
			t.Errorf("(-expect +result):\n%s", diff)
		}
	})
}
