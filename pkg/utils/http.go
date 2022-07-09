package utils

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

func GetValidatMsg(err error) string {
	var verr validator.ValidationErrors
	msg := ""
	if errors.As(err, &verr) {
		for i, f := range verr {
			err := f.ActualTag()
			if f.Param() != "" {
				err = fmt.Sprintf("%s=%s", err, f.Param())
			}
			msg += fmt.Sprintf("%s: %s", f.Field(), err)
			if i != len(verr)-1 {
				msg += ", "
			}

		}
		return msg
	}

	return "bad params"

}
