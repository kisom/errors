package errors

import (
	"fmt"
	"os"
	"testing"
)

var testAction Action = "pass the tests"

func TestFrom(t *testing.T) {
	err := fmt.Errorf("test error")
	aerr := From(err, "dummy action")
	t.Log(aerr)
}

func TestAction(t *testing.T) {
	err := testAction.New("this is a test of an actionable error")
	fmt.Fprintln(os.Stderr, err)
}
