package interviews

import (
	"fmt"
)
func checkValue(val int) {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()

    if val < 0 {
        panic(fmt.Sprintf("Invalid value: %d", val))
    }
}

func runCheckValue() {
    checkValue(-1)
    fmt.Println("Program continues after recovery")
}