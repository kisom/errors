package errors

import "runtime/debug"

type actionableError struct {
	M      string
	Action string
	Frame  string
	Next   error
	s      string
}

func (err *actionableError) String() string {
	if err.s != "" {
		return err.s
	}
	err.s = "Error: " + err.M + "\n" + "\tRemediation: " + err.Action + "\n" + err.Frame
	return err.s
}

func (err *actionableError) Error() string {
	s := err.String()
	if err.Next != nil {
		s += "\ncaused by:\n" + err.Next.Error()
	}
	return s
}

func new(msg, action string) *actionableError {
	err := &actionableError{
		M:      msg,
		Action: action,
	}

	err.Frame = string(debug.Stack())
	return err
}

// New creates a new error.
func New(msg, action string) error { return new(msg, action) }

// Wrap encapsulates an existing error as an actionable error.
func Wrap(err error, msg, action string) error {
	aerr := new(msg, action)
	aerr.Next = err
	return aerr
}

// From creates an actionable error using an existing error.
func From(err error, action string) error {
	return New(err.Error(), action)
}

// Action is used to preset the action field for convenience.
type Action string

// New creates a new error.
func (a Action) New(msg string) error {
	return New(msg, string(a))
}

// Wrap encapsulates an existing error as an actionable error.
func (a Action) Wrap(err error, msg string) error {
	return Wrap(err, msg, string(a))
}

// From creates an actionable error using an existing error.
func (a Action) From(err error) error {
	return From(err, string(a))
}
