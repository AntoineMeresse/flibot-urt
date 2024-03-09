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