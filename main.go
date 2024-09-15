package main

import (
	"context"
	"github.com/quinn71/todolist/internal/config"
	"github.com/quinn71/todolist/internal/listener"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	l := listener.New(cfg)
	if err := l.Run(context.Background()); err != nil {
		panic(err)
	}
}
