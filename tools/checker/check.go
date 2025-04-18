package checker

import (
	"github.com/eos/tools/errs"
)

type Checker interface {
	Check() error
}
// only use check interface 
func Validate(args interface{}) error {
	if checker, ok := args.(Checker); ok {
		if err := checker.Check(); err != nil {
			if _, ok := errs.Unwrap(err).(errs.CodeError); ok {
				return err
			}
			return errs.ErrArgs.WrapMsg(err.Error())
		}
	}
	return nil
}
