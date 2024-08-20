package utils

import (
	"testing"
)

func Test_DecolorString(t *testing.T) {
    given := "^7Flir^1oo^7w"
	want := "Fliroow"
	res := DecolorString(given)

	if want != res {
		t.Errorf("DecolorString result wasn't correct, got: %s, want: %s", res, want);
	}
}

func Test_IsDigitOnly_ShouldReturnTrue(t *testing.T) {
    given := "1234"
	want := true
	res := IsDigitOnly(given)

	if want != res {
		t.Errorf("IsDigitOnly (%s) result wasn't correct, got: %v, want: %v", given, res, want);
	}
}

func Test_IsDigitOnly_ShouldReturnFalse(t *testing.T) {
    given := "12a34"
	want := false
	res := IsDigitOnly(given)

	if want != res {
		t.Errorf("IsDigitOnly (%s) result wasn't correct, got: %v, want: %v", given, res, want);
	}
}

func Test_CleanEmptyElements(t *testing.T) {
	given := []string{"", "a","b", "", "c", ""}
	want := []string{"a","b","c"}
	res := CleanEmptyElements(given)

	if len(want) != len(res) || want[0] != "a" || want[1] != "b" || want[2] != "c" {
		t.Errorf("CleanEmptyElements (%s) result wasn't correct, got: %v, want: %v", given, res, want);
	}
}

func Test_CleanDuplicateElements(t *testing.T) {
	given := []string{"a", "b", "a", "c", "c", "b", "d"}
	want := []string{"a","b","c","d"}
	res := CleanDuplicateElements(given)

	if len(want) != len(res) || want[0] != "a" || want[1] != "b" || want[2] != "c" || want[3] != "d"{
		t.Errorf("CleanDuplicateElements (%s) result wasn't correct, got: %v, want: %v", given, res, want);
	}
}

func Test_GetColorRun(t *testing.T) {
	given := 1
	res := GetColorRun(given)
	want := "^2"

	if (res != want) {
		t.Errorf("Color isn't correct for given index: %d. Expected: (%s), got: (%s)", given, want, res)
	}
} 

func Test_IsVoteCommand(t *testing.T) {
	given := "+"

	if IsVoteCommand(given) != true {
		t.Errorf("(%s) is a vote and should return true", given)
	}

	given = "-"

	if IsVoteCommand(given) != true {
		t.Errorf("(%s) is a vote and should return true", given)
	}

	given = "v"

	if IsVoteCommand(given) != true {
		t.Errorf("(%s) is a vote and should return true", given)
	}

	given = "other"

	if IsVoteCommand(given) != false {
		t.Errorf("(%s) is not a vote and should return false", given)
	}
}