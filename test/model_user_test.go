package test

import (
	"testing"

	"gitee.com/panda8xy/gin-blog/utils"
)

func init() {

}

func TestScryptPassword(t *testing.T) {
	ret := utils.ScryptPassword("panda")
	t.Log(ret)
}
