package pack

import (
	"errors"

	"MyDouyin/pkg/errno"

	"MyDouyin/cmd/user/kitex_gen/user"
)

// BuilduserRegisterResp build userRegisterResp from error
func BuilduserRegisterResp(err error) *user.DouyinUserRegisterResponse {
	if err == nil {
		return userRegisterResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return userRegisterResp(e)
	}

	s := errno.ErrUnknown.WithMessage(err.Error())
	return userRegisterResp(s)
}

func userRegisterResp(err errno.ErrNo) *user.DouyinUserRegisterResponse {
	return &user.DouyinUserRegisterResponse{StatusCode: int32(err.ErrCode), StatusMsg: &err.ErrMsg}
}

// BuilduserResp build userResp from error
func BuilduserUserResp(err error) *user.DouyinUserResponse {
	if err == nil {
		return userResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return userResp(e)
	}

	s := errno.ErrUnknown.WithMessage(err.Error())
	return userResp(s)
}

func userResp(err errno.ErrNo) *user.DouyinUserResponse {
	return &user.DouyinUserResponse{StatusCode: int32(err.ErrCode), StatusMsg: &err.ErrMsg}
}
