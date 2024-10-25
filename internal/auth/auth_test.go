package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	type test struct {
		name   string
		input  http.Header
		output string
		err    error
	}

	tests := []test{
		{
			name: "ApiKey Valid",
			input: http.Header{
				"Authorization": []string{"ApiKey 1234"},
			},
			output: "1234",
			err:    nil,
		},
		{
			name: "ApiKey Invalid",
			input: http.Header{
				"Authorization": []string{"ApiKey"},
			},
			output: "",
			err:    errors.New("alformed authorization header"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := GetAPIKey(tc.input)
			if tc.err != nil && err != nil {

				if err.Error() != tc.err.Error() {
					t.Errorf("expected error %v, got %v", tc.err, err)
				}
			}
			if got != tc.output {
				t.Errorf("expected %v, got %v", tc.output, got)
			}
		})
	}

}
