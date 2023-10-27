package errlog

import (
	"context"
	"fmt"
)

func init() {
	fmt.Println("init v1.1")
}

func Errorf(ctx context.Context, a ...any) {
	fmt.Printf("errlog v1.1:%v", a)
}
