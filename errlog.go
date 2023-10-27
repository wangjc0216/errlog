package errlog

import (
	"context"
	"fmt"
)

func init() {
	fmt.Println("init v2.1")
}

func Errorf(ctx context.Context, a ...any) {
	fmt.Printf("errlog v2.1:%v", a)
}
