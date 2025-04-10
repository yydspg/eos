package auth

 import (
	"errors"
	"github.com/eos/protocol/constant"
 )
func (q *GetAdminTokenReq) Check() error {
	if q.UserID == "" {
		return errors.New("UserID is empty")
	}
	return nil
}

func (q *ForceLogoutReq) Check() error {
	if q.UserID == "" {
		return errors.New("userID is empty")
	}
	if q.PlatformID > constant.AdminPlatformID  || q.PlatformID < constant.IOSPlatformID {
		return errors.New("platformID is invalidate")
	}
	return nil
}

func (q *ParseTokenReq) Check() error {
	if q.Token == "" {
		return errors.New("userID is empty")
	}
	return nil
}
func (q *GetUserTokenReq) Check() error {
	if q.UserID == "" {
		return errors.New("userID is empty")
	}
	if q.PlatformID == 0  {
		return errors.New("platformID is empty")
	}
	return nil
}