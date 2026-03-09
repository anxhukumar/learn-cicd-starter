package auth

import (
	"net/http"
	"testing"
)

func TestGetAPI(t *testing.T) {

	tests := []struct {
		name    string
		headers http.Header
		value   string
		wantErr bool
	}{
		{
			name: "Valid header",
			headers: http.Header{
				"Authorization": []string{"ApiKey djfojewojoadsjofjwe"},
			},
			value:   "djfojewojoadsjofjwe",
			wantErr: false,
		},
		{
			name:    "Authorization header missing",
			headers: http.Header{},
			value:   "djfojewojoadsjofjwe",
			wantErr: true,
		},
		{
			name: "Api key missing",
			headers: http.Header{
				"Authorization": []string{"Token djsflajldfjlasjd;lfjals"},
			},
			value:   "djfojewojoadsjofjwe",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			val, err := GetAPIKey(tt.headers)

			// Check if errors when expected
			if (err != nil) != tt.wantErr {
				t.Fatalf("GetAPIKey() | error = %v, wantErr = %v", err, tt.wantErr)
			}

			if !tt.wantErr && val != tt.value {
				t.Fatalf("GetAPIKey() | Expected value doesn't match")
			}

		})
	}
}
