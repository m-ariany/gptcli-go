package contract

type AssistantInteractor interface {
	Chat(string) <-chan string
}
