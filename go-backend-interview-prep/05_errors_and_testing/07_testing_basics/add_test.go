package main

import "testing"

// assertEqual is a test helper that reports the caller's line on failure.
func assertEqual(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestAdd(t *testing.T) {
	assertEqual(t, Add(2, 3), 5)
	assertEqual(t, Add(0, 0), 0)
	assertEqual(t, Add(-1, 1), 0)
	assertEqual(t, Add(-3, -7), -10)
}

func TestAbs(t *testing.T) {
	assertEqual(t, Abs(-7), 7)
	assertEqual(t, Abs(5), 5)
	assertEqual(t, Abs(0), 0)
}

func TestIsEven(t *testing.T) {
	tests := []struct {
		input int
		want  bool
	}{
		{0, true},
		{1, false},
		{2, true},
		{-3, false},
		{-4, true},
	}

	for _, tc := range tests {
		got := IsEven(tc.input)
		if got != tc.want {
			t.Errorf("IsEven(%d) = %v, want %v", tc.input, got, tc.want)
		}
	}
}
