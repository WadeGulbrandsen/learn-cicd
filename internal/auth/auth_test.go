package auth

import (
	"net/http"
	"strings"
	"testing"
)

func TestAuth(t *testing.T) {
	tests := map[string]struct {
		key       string
		value     string
		expect    string
		expectErr string
	}{
		"valid":        {key: "Authorization", value: "ApiKey TheSecretKey", expect: "TheSecretKey", expectErr: "not expecting an error"},
		"missing key":  {key: "Authorization", value: "ApiKey", expectErr: "malformed authorization header"},
		"wrong auth":   {key: "Authorization", value: "Bearer TheSecretKey", expectErr: "malformed authorization header"},
		"wrong header": {key: "auth", value: "ApiKey TheSecretKey", expectErr: "no authorization header"},
		"empty":        {expectErr: "no authorization header"},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			header := http.Header{}
			header.Add(test.key, test.value)

			got, err := GetAPIKey(header)
			if err != nil {
				if strings.Contains(err.Error(), test.expectErr) {
					return
				}
				t.Fatalf("Unexpected error: %v\n", err)
			}

			if got != test.expect {
				t.Fatalf("Expected: %v, Got: %v\n", test.expect, got)
			}
		})
	}
}
