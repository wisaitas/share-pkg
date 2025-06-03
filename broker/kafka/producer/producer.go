package producer

type Producer interface {
	SendMessage(topic string, message []byte) error
}

type producer struct {
}

func NewProducer() Producer {
	return &producer{}
}

func (p *producer) SendMessage(topic string, message []byte) error {
	return nil
}
