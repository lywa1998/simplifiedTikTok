package errno

import (
    "errors"
    "fmt"
)

const (
    SuccessCode = 0
    ServiceErrCode = iota + 1000
    ParamErrCode
    AuthorizationFailedErrCode

    UserAlreadyExistErrCode
    UserIsNotExistErrCode

    MessageAddFailedErrCode
)

const (
    SuccessMsg = "Success"
    ServerErrMsg = "Service is unable to start successfully"

    UserAlreadyExistErrMsg = "user already exist"
    UserIsNotExistErrMsg = "user is not exist"

    MessageAddFailedErrMsg = "message add failed"
)

var (
    Success = NewDouYinErr(SuccessCode, SuccessMsg)
    ServiceErr = NewDouYinErr(ServiceErrCode, ServerErrMsg)
    UserAlreadyExistErr = NewDouYinErr(UserAlreadyExistErrCode, UserAlreadyExistErrMsg)
    UserIsNotExistErr = NewDouYinErr(UserIsNotExistErrCode, UserIsNotExistErrMsg)
    MessageAddFailedErr = NewDouYinErr(MessageAddFailedErrCode, MessageAddFailedErrMsg)
)

type DouYinErr struct {
    ErrCode int32
    ErrMsg string
}

func (e DouYinErr) Error() string {
    return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

func NewDouYinErr(code int32, msg string) DouYinErr {
    return DouYinErr{code, msg}
}

func (e DouYinErr) WithMessage(msg string) DouYinErr {
    e.ErrMsg = msg
    return e
}

func ConvertErr(err error) DouYinErr {
    Err := DouYinErr{}
    if errors.As(err, &Err) {
        return Err
    }

    s := ServiceErr
    s.ErrMsg = err.Error()
    return s
}
