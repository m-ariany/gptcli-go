package assistant

import (
	"context"
	"fmt"
	"time"
)

func (assistant *Assistant) Chat(statement string) <-chan string {
	text := make(chan string)

	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
		defer cancel()

		defer close(text)

		for message := range assistant.server.Chat(ctx, statement) {
			if message.Error != nil {
				text <- fmt.Sprintf("%s", message.Error)

				return
			}

			text <- message.Data
		}
	}()

	return text
}
