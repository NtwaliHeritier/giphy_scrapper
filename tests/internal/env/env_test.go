package env

import (
	"os"
	"testing"

	env "github.com/ntwaliheritier/giphy_scrapper/internal/env"
)

func TestGetString(t *testing.T) {
		t.Run("returns value when key exists", func(t *testing.T) {
			key := "API_KEY"
			want := "123"

			os.Setenv(key, want)
			defer os.Unsetenv(key)

			got, err := env.GetString(key)

			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}

			if got != want {
				t.Errorf("expected %s; got %s", want, got)
			}
	})

	t.Run("returns error when the key does not exists", func(t *testing.T) {
		key := "API_KEY"
		want := "you need to save an api key in your .env file as API_KEY"
		_, err := env.GetString(key)

		if err == nil {
			t.Errorf("expected an error %s got nil", want)
		}
	})
}