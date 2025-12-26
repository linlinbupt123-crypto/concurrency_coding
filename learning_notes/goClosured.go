package learningnotes

import (
	"fmt"
	"time"
)

func goclosurederr() {
	i := 0
	for ; i < 3; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
}

// no problem after go1.22 官方做了修复，for循环中的i，每次都是一个新的变量
func goclosured() {
	for i := 0; i <= 200; i++ {
		go func() {
			fmt.Print(i, " ")
		}()
	}
	time.Sleep(time.Second)
	fmt.Println()
}
