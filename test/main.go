package main

import (
	"fmt"
	"reflect"

	utils "github.com/2637309949/bulrush-utils"
)

func main() {
	// 1.array
	a := 1
	arr := utils.Append(1, []int{1, 2, 3})
	fmt.Println(arr)
	arr = utils.Append(1, &[]int{1, 2, 3})
	fmt.Println(arr)
	arr = utils.Append(&a, reflect.New(reflect.TypeOf([]int{})).Interface())
	fmt.Println(arr)
}
