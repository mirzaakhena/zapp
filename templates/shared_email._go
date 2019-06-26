package email

import "fmt"

// IEmail is
type IEmail interface {
	Send(from, to, subject, message string) error
}

type EmailImpl struct {
}

func NewEmailImpl() *EmailImpl {
	return &EmailImpl{}
}

func (e *EmailImpl) Send(from, to, subject, message string) error {
	fmt.Printf(">>>>> email send to %s with message %s", to, message)
	return nil
}
