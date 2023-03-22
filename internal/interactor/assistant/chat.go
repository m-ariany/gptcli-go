package assistant

import (
	"context"
	"fmt"
	"time"
)

func (assistant *Assistant) Chat(statement string) <-chan string {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	text := make(chan string)
	go func() {
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
