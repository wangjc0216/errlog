package errlog

import "fmt"

func init() {
	fmt.Println("init v1.1")
}

func Errorf(a ...any) {
	fmt.Printf("errlog v1.1:%v", a)
}
