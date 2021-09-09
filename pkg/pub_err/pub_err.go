package pub_err

import "errors"

var (
	ErrLackParams = errors.New("参数缺失或非法")
)
