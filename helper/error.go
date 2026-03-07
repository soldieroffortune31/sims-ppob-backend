package helper

import "fmt"

func PanicIfError(err error) {
	if err != nil {
		fmt.Print(err)
		panic(err)
	}
}
