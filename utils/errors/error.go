package errors

const (
	SUCCESS = 200
	ERROR   = 500

	// code = 1001 - 1999 user errs

	ERROR_USERNAME_USED    = 1001
	ERROR_PASSWORD_WRONG   = 1002
	ERROR_USER_NOT_FOUND   = 1003
	ERROR_TOKEN_NOT_EXIST  = 1004
	ERROR_TOKEN_OUTTIME    = 1005
	ERROR_TOKEN_WRONG      = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007
	ERROR_USER_NO_RIGHTS   = 1008 //  no right login backened
	ERROR_FILE_UPLOAD      = 1009

	// code = 2001 - 2999 article errs

	ERROR_ARTICLE_USED      = 2001
	ERROR_ARTICLE_NOT_EXIST = 2002

	// code = 3001 - 3999 category errs

	ERROR_CATEGORY_USED = 3001
)

var codeMap = map[int]string{
	SUCCESS:                 "OK",
	ERROR:                   "FIAL",
	ERROR_USERNAME_USED:     "用户名已存在",
	ERROR_PASSWORD_WRONG:    "密码错误",
	ERROR_USER_NOT_FOUND:    "用户未找到",
	ERROR_TOKEN_NOT_EXIST:   "TOKEN不存在",
	ERROR_TOKEN_OUTTIME:     "TOKEN已过期",
	ERROR_TOKEN_WRONG:       "TOKEN不正确",
	ERROR_TOKEN_TYPE_WRONG:  "TOKEN格式错误",
	ERROR_CATEGORY_USED:     "Cate已存在",
	ERROR_ARTICLE_USED:      "文章已存在",
	ERROR_ARTICLE_NOT_EXIST: "文章不存在",
	ERROR_USER_NO_RIGHTS:    "用户没有权限",
	ERROR_FILE_UPLOAD:       "文件上传出错",
}

// Msg getter for errors
func Msg(codeNum int) string {
	return codeMap[codeNum]
}
