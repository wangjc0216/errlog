package errlog

import "fmt"

func init() {
	fmt.Println("init v1.0")
}

func Errorf(a ...any) {
	fmt.Printf("errlog v1.0:%v", a)
}
