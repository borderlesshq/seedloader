package seedloader

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetSeed(t *testing.T) {

	scenarios := []struct {
		scenarioName  string
		serviceName   string
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
			"ms.admin",
			"create-admin/valid.json",
			nil,
		},
		{
			"Cannot Open File",
			"ms.admin",
			"non-existing",
			ErrOpeningSeed,
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.scenarioName, func(t *testing.T) {
			seeder, err := NewSeedLoader(scenario.serviceName, "/fish")
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
		})
	}

}
