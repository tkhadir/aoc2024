package utils

import (
    "os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}


func ReadFile(filename string) string {
	dat, err := os.ReadFile(filename)
    check(err)
    return string(dat)
}