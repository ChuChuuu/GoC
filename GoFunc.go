package GoC

import "C"
import "fmt"

//export GoPrint
func GoPrint(ptr *C.char) {
	fmt.Println("This is Go Func:",C.GoString(ptr))
}