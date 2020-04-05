package pgutils

import (
	"strings"

	"github.com/lib/pq"
	"github.com/nanoTitan/analytics-users-api/utils/errors"
)

const (
	errorNoRows = "no rows in result set"
)

// ParseError - parse a postgres REST service error given an error object
func ParseError(err error) *errors.RestErr {
	pgErr, ok := err.(*pq.Error)
	if !ok || pgErr.Message == "" {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError("no records found")
		}
		return errors.NewInternalServerError("error parsing database response")
	}

	return errors.NewDbError(pgErr.Message, pgErr.Code.Name())
}
