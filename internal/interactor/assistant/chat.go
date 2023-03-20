package assistant

func (assistant *Assistant) Chat(statement string) <-chan string {
	text := make(chan string)

	go func() {
		defer close(text)

		text <- "I'm an AI language model"
	}()

	return text
}
