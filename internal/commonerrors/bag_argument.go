package commonerrors

import (
	"fmt"
	"strings"
)

const (
	badArgumentMessage = "bad argument"
)

func IsBadArgument(err error) bool {
	return strings.Contains(err.Error(), badArgumentMessage)
}

func BadArgument(err error) error {
	return fmt.Errorf("%s: %w", badArgumentMessage, err)
}
