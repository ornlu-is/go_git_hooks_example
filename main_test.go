package main

import "testing"

func TestIsEven(t *testing.T) {
	for _, tc := range []struct {
		name           string
		givenNumber    int
		expectedResult bool
	}{
		{
			name:           "given an even number returns true",
			givenNumber:    6661,
			expectedResult: true,
		},
		{
			name:           "given an odd number returns false",
			givenNumber:    333,
			expectedResult: false,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			res := isEven(tc.givenNumber)

			if res != tc.expectedResult {
				t.Errorf("expected result %t, but got %t", tc.expectedResult, res)
			}
		})
	}
}
