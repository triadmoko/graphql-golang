package user

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/triadmoko/grahpql-golang/graph/request"
	"gorm.io/gorm"
)

func (s *user_service) VerifyEmail(ctx context.Context, req request.NewVerify) (string, error) {
	verify, err := s.userRepository.GetVerifyByEmail(ctx, req.Email)
	if err != nil && err != gorm.ErrRecordNotFound {
		s.loggger.Error("verify not found ", err.Error())
		return "", errors.New("verify not found")
	}
	fmt.Println(time.Now().UTC().Unix() - verify.Expired)
	if time.Now().UTC().Unix() >= verify.Expired {
		return "", errors.New("Verification expired")
	}

	if verify.Code != req.Code {
		return "", errors.New("Code incorrect")
	}
	err = s.userRepository.UpdateStatusUser(ctx, verify.UserID, "active")
	if err != nil {
		s.loggger.Error("Error failed update status user", err.Error())
		return "", errors.New("verify not found")
	}
	return "active", nil
}
