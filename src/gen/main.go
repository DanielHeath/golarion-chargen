package main

import (
	"fmt"
	"gen/types"
	"math/rand"
	"time"
)

func main() {
	seed := time.Now().Unix()
	fmt.Printf("Seed: %d\n", seed)
	rand.Seed(seed)

	fmt.Println(types.NewCharacter())
	fmt.Println(types.NewCharacter())
	fmt.Println(types.NewCharacter())
	fmt.Println(types.NewCharacter())
	fmt.Println(types.NewCharacter())
	fmt.Println(types.NewCharacter())
	fmt.Println(types.NewCharacter())
	fmt.Println(types.NewCharacter())
	fmt.Println(types.NewCharacter())
	fmt.Println(types.NewCharacter())
	fmt.Println(types.NewCharacter())
	fmt.Println(types.NewCharacter())
	fmt.Println(types.NewCharacter())
	fmt.Println(types.NewCharacter())
}
