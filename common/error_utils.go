package common

import "errors"

const errorString  = "nil"

func ErrorToString(err error) string{
	if err == nil {
		return errorString
	}
	return err.Error()
}


func ErrorFromString(err string) error{
	if err == errorString {
		return nil
	}
	return errors.New(err)
}
