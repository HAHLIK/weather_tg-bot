package pkg

import "fmt"

func ErrorWrap(message string, errMessage any) error {
	var err error = nil

	if errMessage != nil {
		err = fmt.Errorf("%s : %v", message, errMessage)
	}

	return err
}
