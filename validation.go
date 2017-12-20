package c

import "errors"

func ValidateExistInMap(m map[string]struct{}, v string, msg string, errs []error) []error {
	isValueInMap := func(m map[string]struct{}, val string) bool {
		_, ok := m[val]
		return ok
	}

	if ok := isValueInMap(m, v); !ok {
		errs = append(errs, errors.New(msg))
	}
	return errs
}

func ValidateRangeCondition(min int, max int, val int, msg string, errs []error) []error {
	isValidRange := func(min int, max int, val int) bool {
		return !(val < min || val > max)
	}
	if ok := isValidRange(min, max, val); !ok {
		errs = append(errs, errors.New(msg))
	}
	return errs
}

func ValidatePositiveNumber(n int, msg string, errs []error) []error {
	if n < 0 {
		errs = append(errs, errors.New(msg))
	}
	return errs
}

func ValidateCondition(fn func() bool, msg string, errs []error) []error {
	if ok := fn(); !ok {
		errs = append(errs, errors.New(msg))
	}
	return errs
}
