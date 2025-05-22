package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name        string
		header      http.Header
		expectedKey string
		expectedErr string
	}{
		{
			name: "valid authorization header",
			header: http.Header{
				"Authorization": []string{"ApiKey my-key"},
			},
			expectedKey: "my-key",
			expectedErr: "",
		},
		{
			name:        "no authorization header",
			header:      http.Header{},
			expectedKey: "",
			expectedErr: ErrNoAuthHeaderIncluded.Error(),
		},
		{
			name: "missing ApiKey prefix",
			header: http.Header{
				"Authorization": []string{"Key my-key"},
			},
			expectedKey: "",
			expectedErr: "malformed authorization header",
		},
		{
			name: "missing Key",
			header: http.Header{
				"Authorization": []string{"ApiKey"},
			},
			expectedKey: "",
			expectedErr: "malformed authorization heade",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, err := GetAPIKey(tt.header)
			if err != nil && err.Error() != tt.expectedErr {
				t.Errorf("GetAPIKey() error = %v, want %v", err, tt.expectedErr)
			}
			if key != tt.expectedKey {
				t.Errorf("GetAPIKey() gotKey = %v, want %v", key, tt.expectedKey)
			}
		})
	}
}
