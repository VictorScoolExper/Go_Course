// chap1/manual-parse/validate_args_test.go
package main

import (
	"errors"
	"testing"
)

func TestValidateArgs(t *testing.T) {
	// TODO Insert definition tests[] slice as above
	for _, tc := range tests {
		err := validateArgs(tc.c)
		if tc.err != nil && err.Error() != tc.err.Error() {
			t.Errorf("Expected error to be: %v, got: %v\n",
				tc.err, err)
		}
		if tc.err == nil && err != nil {
			t.Errorf("Expected nil error, got: %v\n", err)
		}
	}
}
