package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	fmt.Println(ctx)
	fmt.Println(ctx.Err())
	fmt.Println(ctx.Value("key"))
	fmt.Println(ctx.Deadline())
	fmt.Println(ctx.Done())

}
