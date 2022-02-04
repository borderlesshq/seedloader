package seedloader

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetSeed(t *testing.T) {

	scenarios := []struct {
		scenarioName  string
		packageName   string
		inputSeed     string
		expectedError error
	}{
		{
			"Initialize Seeder",
			"ms.fish",
			"",
			ErrInitializingSeeder,
		},
		{
			"Correct Payload",
			"github.com/borderlesshq/seedloader",
			"user.json",
			nil,
		},
		{
			"Cannot Open File",
			"github.com/borderlesshq/seedloader",
			"non-existing",
			ErrOpeningSeed,
		},
	}

	type User struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
	}

	for _, scenario := range scenarios {
		t.Run(scenario.scenarioName, func(t *testing.T) {
			seeder, err := NewSeedLoader(scenario.packageName, "seeds")
			if err != nil {
				assert.Equal(t, err, scenario.expectedError)
				return
			}

			b, err := seeder.GetSeed(scenario.inputSeed)
			assert.Equal(t, err, scenario.expectedError)
			if err == nil {
				assert.NotNil(t, b)
				assert.NotEmpty(t, b)
			}

			var user User
			err = seeder.ParseSeed("user.json", &user)
			assert.Nil(t, err)

			assert.Equal(t, user.FirstName, "Borderless")
			assert.Equal(t, user.LastName, "HQ Inc.")
		})
	}

}
