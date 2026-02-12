package main

import "testing"

func TestParseIntSafe(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		fallback int
		want     int
		wantErr  bool
	}{
		{"valid positive", "42", 0, 42, false},
		{"valid negative", "-7", 0, -7, false},
		{"valid zero", "0", 99, 0, false},
		{"empty string", "", -1, -1, true},
		{"letters", "abc", -1, -1, true},
		{"float string", "3.14", 0, 0, true},
		{"overflow", "99999999999999999999", 0, 0, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ParseIntSafe(tc.input, tc.fallback)
			if (err != nil) != tc.wantErr {
				t.Fatalf("ParseIntSafe(%q, %d) error = %v, wantErr %v",
					tc.input, tc.fallback, err, tc.wantErr)
			}
			if got != tc.want {
				t.Errorf("ParseIntSafe(%q, %d) = %d, want %d",
					tc.input, tc.fallback, got, tc.want)
			}
		})
	}
}

func TestClamp(t *testing.T) {
	tests := []struct {
		name       string
		n, lo, hi  int
		want       int
	}{
		{"in range", 5, 0, 10, 5},
		{"below min", -3, 0, 10, 0},
		{"above max", 15, 0, 10, 10},
		{"at min", 0, 0, 10, 0},
		{"at max", 10, 0, 10, 10},
		{"negative range", -5, -10, -1, -5},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Clamp(tc.n, tc.lo, tc.hi)
			if got != tc.want {
				t.Errorf("Clamp(%d, %d, %d) = %d, want %d",
					tc.n, tc.lo, tc.hi, got, tc.want)
			}
		})
	}
}
