package mod1

import (
	"fmt"
	"github.com/sesky4/worktest/common"
)

func GetVersion() string {
	return "0.0.1"
}

func PrintVersion() {
	fmt.Printf("mod1: ver=%s common_ver=%s\n", GetVersion(), common.GetVersion())
}
