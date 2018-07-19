package pkg1

import (
	"fmt"

	"github.com/8tomat8/sketchs/initFuncs/pkg2"
)

func init() {
	fmt.Println("pkg1 init()")
}

func F1() {
	pkg2.F2()
	fmt.Println("pkg1 F1()")
}
