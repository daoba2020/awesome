package nsq

import (
	"awesome/config"
	"github.com/nsqio/go-nsq"
	"unsafe"
)

type Message nsq.Message
type HandlerFunc func(message *Message) error

type Consumer struct {
	conn    *nsq.Consumer
	channel string
	Topic   string
}

var (
	addr          = config.DbServer.NSQIP
	defaultConfig = nsq.NewConfig()
)

func NewConsumer(topic string, channel string, handlerFunc HandlerFunc) (*Consumer, error) {
	consumer := &Consumer{}
	consumer.Topic = topic
	c, err := nsq.NewConsumer(topic, channel, defaultConfig)
	if err != nil {
		return nil, err
	}

	consumer.conn = c
	c.AddHandler(convertHandlerFunc(handlerFunc))
	if err := c.ConnectToNSQD(addr); err != nil {
		return nil, err
	}

	//建立连接
	return consumer, nil
}

func (consumer *Consumer) Stop() {
	consumer.conn.Stop()
}

type Producer struct {
	conn  *nsq.Producer
	Topic string
}

func NewProducer(topic string) (*Producer, error) {
	producer := &Producer{}
	producer.Topic = topic

	// 新建生产者
	p, err := nsq.NewProducer(addr, defaultConfig)
	if err != nil {
		return nil, err
	}

	// 发布消息
	producer.conn = p
	producer.Topic = topic

	return producer, nil
}

func (producer *Producer) Publish(bytes []byte) error {
	return producer.conn.Publish(producer.Topic, bytes)
}

func (producer *Producer) Stop() {
	producer.conn.Stop()
}

func convertHandlerFunc(handlerFuncLocal HandlerFunc) nsq.HandlerFunc {
	return func(nsqMessage *nsq.Message) error {
		convertMessage := (*Message)(unsafe.Pointer(nsqMessage))
		return handlerFuncLocal(convertMessage)
	}
}
