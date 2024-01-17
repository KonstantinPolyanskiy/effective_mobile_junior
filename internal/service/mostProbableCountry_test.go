package service

import (
	"effective_mobile_junior/external/nationalize"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMostProbableCountry(t *testing.T) {
	testCases := []struct {
		name                string
		expectedCode        string
		expectedProbability float64
		input               nationalize.Result
	}{
		{
			name:                "good",
			expectedCode:        "AU",
			expectedProbability: 0.5,
			input: nationalize.Result{
				Count: 10,
				Name:  "TestUserName",
				Country: []nationalize.Country{
					{"AU", 0.5}, {"BB", 0.4}, {"CB", 0.3}, {"OO", 0},
				},
				Code: 200,
			},
		},
		{
			name:                "equal probability",
			expectedCode:        "QW",
			expectedProbability: 0.1,
			input: nationalize.Result{
				Count: 10,
				Name:  "TestUserName",
				Country: []nationalize.Country{
					{"QW", 0.1}, {"QE", 0.1}, {"QR", 0.1}, {"QT", 0.1},
				},
				Code: 200,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c, p := mostProbableCountry(tc.input)

			assert.Equal(t, tc.expectedCode, c)
			assert.Equal(t, tc.expectedProbability, p)
		})
	}
}
