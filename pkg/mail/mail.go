package mail

import "sync"

type From struct {
	Address string
	Name    string
}

type Email struct {
	From    From
	To      []string
	Subject string
	Text    []byte
	HTML    []byte
}

type Mailer struct {
	Driver Driver
}

var once sync.Once
var internalMailer *Mailer

func NewEmailer(driver Driver) *Mailer {
	once.Do(func() {
		internalMailer = &Mailer{
			Driver: driver,
		}
	})
	return internalMailer
}

func (m *Mailer) Send(email Email, config map[string]string) bool {
	return m.Driver.Send(email, config)
}
