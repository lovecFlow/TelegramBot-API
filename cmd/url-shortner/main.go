package main

import "telegrambot-api/internal/config"

func main() {
	//TODO init config: cleanenv
	cfg := config.Load()

	//TODO init logger: slog

	//TODO init storage

	//TODO init router

	//TODO run server
}
