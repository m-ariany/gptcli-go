package assistant

import (
	"context"
	"time"
)

func (assistant *Assistant) Chat(statement string) <-chan string {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	_ = ctx

	text := make(chan string)

	go func() {
		defer close(text)

		text <- "I'm an AI language model"
	}()

	return text
}
