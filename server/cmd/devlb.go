package main

import (
	h "github.com/kradnoel/DEVLOGBOOK/server/internal/http"
)

func main() {
	server := h.New()
	server.Execute()
}
