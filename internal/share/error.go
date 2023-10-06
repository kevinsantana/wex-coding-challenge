package share

import (
	"errors"
	"fmt"
	"strings"
)

// internal infra errors
var (
	// database
	ErrUpdateNoneOrMany = errors.New("error in quantity of rows affected")
	ErrDatabase         = errors.New("error database")
	ErrDuplicate        = errors.New("error duplicate")
	ErrOrderNotCreated  = errors.New("purchase not created, please try again")
	// rest
	ErrTimeout     = errors.New("error timeout")
	ErrContentType = errors.New("content-type error")
	ErrValidation  = errors.New("validation error")
)

type DomainError struct {
	Domain      string
	Module      string
	Err         string
	Description string
}

func (err DomainError) Error() string {
	return fmt.Sprintf(
		"%s|%s|%s",
		strings.ToUpper(err.Domain),
		strings.ToUpper(err.Module),
		strings.ToUpper(err.Err),
	)
}

type ClientError struct {
	Domain      string
	Module      string
	Err         string
	Description string
}

func (err ClientError) Error() string {
	return fmt.Sprintf(
		"%s|%s|%s",
		strings.ToUpper(err.Domain),
		strings.ToUpper(err.Module),
		strings.ToUpper(err.Err),
	)
}
