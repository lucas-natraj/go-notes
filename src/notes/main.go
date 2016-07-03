package main

import (
	"net/http"
	_ "notes/app"
	"os"
)

func getenv(key string, fallback string) string {
	v := os.Getenv(key)
	if len(v) != 0 {
		return v
	}
	return fallback
}

// used for local testing / debugging
func main() {
	http.ListenAndServe(":" + getenv("PORT", "3000"), nil)
}