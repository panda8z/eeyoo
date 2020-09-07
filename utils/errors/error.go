package errors

const (
	SUCCESS = 200
	ERROR   = 500

	// code = 1000 - 1999 user errs
	ERROR_USERNAME_USED    = 1001
	ERROR_PASSWORD_WRONG   = 1002
	ERROR_USER_NOT_FOUND   = 1003
	ERROR_TOKEN_EXIST      = 1004
	ERROR_TOKEN_OUTTIME    = 1005
	ERROR_TOKEN_WRONG      = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007

	// code = 2000 - 2999 article errs

	// code = 3000 - 3999 category errs
)

var codeMap = map[int]string{
	SUCCESS:                "OK",
	ERROR:                  "FIAL",
	ERROR_USERNAME_USED:    "用户名已存在",
	ERROR_PASSWORD_WRONG:   "密码错误",
	ERROR_USER_NOT_FOUND:   "用户未找到",
	ERROR_TOKEN_EXIST:      "TOKEN不存在",
	ERROR_TOKEN_OUTTIME:    "TOKEN已过期",
	ERROR_TOKEN_WRONG:      "TOKEN不正确",
	ERROR_TOKEN_TYPE_WRONG: "TOKEN格式错误",
}

// Code getter for errors
func Code(codeNum int) string {
	return codeMap[codeNum]
}
