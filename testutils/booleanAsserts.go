package testutils

import "testing"

func AssertFalse(t *testing.T, decision bool, failureMessage string) {
	AssertTrue(t, !decision, failureMessage)
}

func AssertTrue(t *testing.T, decision bool, failureMessage string) {
	if decision {
		return
	} else {
		t.Fatal(failureMessage)
	}
}
