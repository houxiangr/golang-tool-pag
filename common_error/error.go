package common_error

import (
	"encoding/json"
)

type Error struct {
	ErrNo  int    `json:"errno"`
	ErrMsgMap map[int]string `json:"errmsg"`
}

func NewError(err Error, msg string) Error {
	errMap := make(map[int]string)
	errMap[1]=msg
	return Error{err.ErrNo, errMap}
}

func (e Error) Error() string {
	errMsg,_ := json.Marshal(e.ErrMsgMap)
	return string(errMsg)
}

func (e Error) GetErrNo() int {
	return e.ErrNo
}

func (e Error) SetExtraMsg(s string) Error {
	errLevel := len(e.ErrMsgMap)
	e.ErrMsgMap[errLevel+1] = s
	return e
}

var (
	Success = Error{0, map[int]string{1:"success"}}
)
