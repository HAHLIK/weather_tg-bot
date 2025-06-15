package pkg

import "fmt"

func ErrorWrap(message string, errMessage any) error {
	var err error = nil

	if errMessage != nil {
		err = fmt.Errorf("%s : %v", message, errMessage)
	}

	return err
}

func CelsiumFromKelvin(temp float64) float64 {
	return (temp - 273.15)
}
