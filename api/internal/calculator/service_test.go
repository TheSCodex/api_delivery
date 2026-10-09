package calculator

import "testing"

func TestCalculate(t *testing.T) {
	tests := []struct {
		name      string
		a         float64
		b         float64
		operation string
		expected  float64
		wantErr   bool
	}{
		{
			name:      "addition",
			a:         9,
			b:         4,
			operation: "+",
			expected:  13,
		},
		{
			name:      "subtraction",
			a:         9,
			b:         4,
			operation: "-",
			expected:  5,
		},
		{
			name:      "multiplication",
			a:         9,
			b:         4,
			operation: "*",
			expected:  36,
		},
		{
			name:      "division",
			a:         8,
			b:         4,
			operation: "/",
			expected:  2,
		},
		{
			name:      "negative operand",
			a:         -5,
			b:         3,
			operation: "+",
			expected:  -2,
		},
		{
			name:      "two negative operands",
			a:         -5,
			b:         -3,
			operation: "*",
			expected:  15,
		},
		{
			name:      "negative result",
			a:         3,
			b:         5,
			operation: "-",
			expected:  -2,
		},
		{
			name:      "decimal addition",
			a:         2.5,
			b:         1.5,
			operation: "+",
			expected:  4,
		},
		{
			name:      "decimal division",
			a:         5,
			b:         2,
			operation: "/",
			expected:  2.5,
		},
		{
			name:      "addition with zero",
			a:         5,
			b:         0,
			operation: "+",
			expected:  5,
		},
		{
			name:      "multiplication by zero",
			a:         5,
			b:         0,
			operation: "*",
			expected:  0,
		},
		{
			name:      "division by zero",
			a:         5,
			b:         0,
			operation: "/",
			wantErr:   true,
		},
		{
			name:      "unsupported operation",
			a:         5,
			b:         3,
			operation: "%",
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Calculate(tt.a, tt.b, tt.operation)

			if tt.wantErr {
				if err == nil {
					t.Errorf("Calculate(%v, %v, %q) expected an error, got nil",
						tt.a, tt.b, tt.operation)
				}
				return
			}

			if err != nil {
				t.Fatalf("Calculate(%v, %v, %q) returned unexpected error: %v",
					tt.a, tt.b, tt.operation, err)
			}

			if result != tt.expected {
				t.Errorf("Calculate(%v, %v, %q) = %v; expected %v",
					tt.a, tt.b, tt.operation, result, tt.expected)
			}
		})
	}
}