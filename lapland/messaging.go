package lapland

import (
	"os"
	"sync"
)

type Attachment os.File

type Message struct {
	Text       []byte
	Attachment Attachment
}

// FIXME: WIP interface
type Sender interface {
	Send(msg Message) error
}

type Office struct {
	msgPool chan Message

	wg sync.WaitGroup
}

func NewOffice(sender Sender, workers, poolSize int) *Office {
	pool := make(chan Message, 1000)
	wg := sync.WaitGroup{}
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go func() {
			for msg := range pool {
				sender.Send(msg)
			}

			wg.Done()
		}()
	}

	return &Office{}
}

func (o *Office) Send(msg Message) {
	o.msgPool <- msg
}
