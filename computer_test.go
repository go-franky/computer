package computer_test

import (
	"testing"

	"github.com/frankywahl/computer"
)

func TestString(t *testing.T) {
	for _, tc := range []struct {
		given    computer.Size
		expected string
	}{
		{
			1 * computer.Kibibyte,
			"1KiB",
		},
		{
			5 * computer.Mebibyte,
			"5MiB",
		},
		{
			5*computer.Mebibyte + 300*computer.Kibibyte,
			"5MiB300KiB",
		},
		{
			5*computer.Exbibyte + 353*computer.Byte,
			"5EiB353B",
		},
	} {
		if tc.given.String() != tc.expected {
			t.Errorf("Expected %s, got %s", tc.expected, tc.given)
		}
	}
}

func TestKibibytes(t *testing.T) {
	for _, tc := range []struct {
		given    computer.Size
		expected float64
	}{
		{
			1 * computer.Mebibyte,
			1024.0,
		},
		{
			512 * computer.Byte,
			0.5,
		},
	} {
		if tc.given.Kibibytes() != tc.expected {
			t.Errorf("Expected %f, got %f", tc.expected, tc.given.Kibibytes())
		}
	}
}
