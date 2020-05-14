package digicert

import (
	"strings"
	"testing"
)

func testExpectedErrorChecker(t *testing.T, expectedError, actualError error) {
	t.Helper()

	if expectedError != nil {
		if actualError == nil || !strings.Contains(actualError.Error(), expectedError.Error()) {
			t.Errorf("Expected error %s, but got %s", expectedError, actualError)
		}
	} else if actualError != nil {
		t.Errorf("Expected no error, but got %s", actualError)
	}
}
