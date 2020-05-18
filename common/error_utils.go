package common

import "errors"

const ErrorNilString = "nil"

func ErrorToString(err error) string{
	if err == nil {
		return ErrorNilString
	}
	return err.Error()
}


func ErrorFromString(err string) error{
	if err == ErrorNilString {
		return nil
	}
	return errors.New(err)
}
