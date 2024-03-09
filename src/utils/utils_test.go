package utils

import (
	"testing"
)

func TestDecolorString(t *testing.T) {
    given := "^7Flir^1oo^7w"
	want := "Fliroow"
	res := DecolorString(given)

	if want != res {
		t.Errorf("DecolorString result wasn't correct, got: %s, want: %s", res, want);
	}
}

func TestIsDigitOnly_ShouldReturnTrue(t *testing.T) {
    given := "1234"
	want := true
	res := IsDigitOnly(given)

	if want != res {
		t.Errorf("IsDigitOnly (%s) result wasn't correct, got: %v, want: %v", given, res, want);
	}
}

func TestIsDigitOnly_ShouldReturnFalse(t *testing.T) {
    given := "12a34"
	want := false
	res := IsDigitOnly(given)

	if want != res {
		t.Errorf("IsDigitOnly (%s) result wasn't correct, got: %v, want: %v", given, res, want);
	}
}