package http

import (
	"encoding/json"
	nethttp "net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCalculateHandler(t *testing.T) {
	tests := []struct {
		name           string
		body           string
		expectedStatus int
		expectedResult float64
		checkResult    bool
	}{
		{
			name:           "successful calculation",
			body:           `{"a":5,"b":6,"operation":"+"}`,
			expectedStatus: nethttp.StatusOK,
			expectedResult: 11,
			checkResult:    true,
		},
		{
			name:           "invalid JSON",
			body:           `{"a":5,"b":}`,
			expectedStatus: nethttp.StatusBadRequest,
			checkResult:    false,
		},
		{
			name:           "division by zero",
			body:           `{"a":10,"b":0,"operation":"/"}`,
			expectedStatus: nethttp.StatusBadRequest,
			checkResult:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(
				nethttp.MethodPost,
				"/calculate",
				strings.NewReader(tt.body),
			)

			recorder := httptest.NewRecorder()

			Calculate(recorder, req)

			if recorder.Code != tt.expectedStatus {
				t.Fatalf(
					"expected status %d, got %d",
					tt.expectedStatus,
					recorder.Code,
				)
			}

			if tt.checkResult {
				var response CalculateResponse

				err := json.NewDecoder(recorder.Body).Decode(&response)
				if err != nil {
					t.Fatalf("failed to decode response: %v", err)
				}

				if response.Result != tt.expectedResult {
					t.Errorf(
						"expected result %v, got %v",
						tt.expectedResult,
						response.Result,
					)
				}
			}
		})
	}
}