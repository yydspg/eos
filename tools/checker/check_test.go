package checker_test

import (
	"testing"

	"github.com/eos/tools/errs"
	"github.com/stretchr/testify/assert"
	"github.com/eos/tools/checker"
)
type testChecker struct {
	err error
}
// 
func(t testChecker) Check() error {
	return t.err
}
func TestValidate(t * testing.T ) {
	cases := []struct{
		name string
		args interface{}
		wantError error
	} {	
		{
		name : "non-check args",
		args : "non-check",
		wantError : nil,
	},
	{
		name : "check no error",
		args : &testChecker{nil},
		wantError : nil,
	},
	{
		name : "check with generic error",
		args : &testChecker{errs.ErrArgs},
		wantError : errs.ErrArgs ,
	},
	{
		name : "check with code error",
		args : &testChecker{errs.NewCodeError(400,"bad args")},
		wantError : errs.NewCodeError(400,"bad args"),
	},
	}

	for _,c := range cases {
		t.Run(c.name,func(t *testing.T) {
			err := checker.Validate(c.args)
			if c.wantError != nil {
				assert.ErrorIs(t,err,c.wantError)
			}else {
				assert.NoError(t,err)
			}
		})
	}
}
