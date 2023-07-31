package types

type AppError interface {
	Status() int
	Message() string
	MessageId() string
	Error() string
}
