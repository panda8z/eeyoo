package test

import (
	"testing"

	"github.com/panda8z/eeyoo/utils"
)

func init() {

}

func TestScryptPassword(t *testing.T) {
	ret := utils.ScryptPassword("panda")
	t.Log(ret)
}
