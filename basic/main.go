package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background() // Empty Context

	ctx1 := context.WithValue(ctx, "key1", "value1") // parent: ctx
	ctx2 := context.WithValue(ctx, "key2", "value2") // parent: ctx
	ctx3 := context.WithValue(ctx, "key3", "value3") // parent: ctx

	ctx4 := context.WithValue(ctx1, "key4", "value4") // parent: ctx1

	fmt.Println(ctx1.Value("key1")) // value1
	fmt.Println(ctx2.Value("key2")) // value2
	fmt.Println(ctx3.Value("key3")) // value3

	fmt.Println(ctx4.Value("key4")) // value4
	fmt.Println(ctx4.Value("key1")) // value1

	fmt.Println(ctx3.Value("key1")) // nil
}
