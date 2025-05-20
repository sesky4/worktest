package mod2

import (
	"fmt"
	"github.com/sesky4/worktest/common"
)

func GetVersion() string {
	return "0.0.3"
}

func PrintVersion() {
	fmt.Printf("mod2: ver=%s common_ver=%s\n", GetVersion(), common.GetVersion())
}
