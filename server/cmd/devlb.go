package main

import (
	h "github.com/kradnoel/teka/internal/http"
)

func main() {
	server := h.New()
	server.Execute()
}
