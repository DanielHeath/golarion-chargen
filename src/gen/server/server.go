package main

import (
	"fmt"
	"gen/server/character_json"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	seed := time.Now().Unix()
	fmt.Printf("Seed: %d\n", seed)
	rand.Seed(seed)

	http.HandleFunc(
		"/character.json",
		character_json.CreateCharacter,
	)
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.ListenAndServe(":8080", nil)
}
