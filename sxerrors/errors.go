package sxerrors

import (
    "fmt"
)

// sample error
type MissingEnvironmentVariable struct {
    Message string
}

func (e MissingEnvironmentVariable) Error() string {
    return fmt.Sprintf("[strix error] [MissingEnvironmentVariable] %v", e.Message)
}
