package repositories

import (
	"flag"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type SendMsgQueue struct {
	Name              string
	MessageAttributes map[string]*sqs.MessageAttributeValue
	MessageBody       string
	MessageGroupId    *string
}

func (p *SendMsgQueue) SendMsgQueue() error {

	flag.Parse()

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := sqs.New(sess)
	result, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: &p.Name,
	})
	if err != nil {
		return err
	}
	_, err = svc.SendMessage(&sqs.SendMessageInput{
		DelaySeconds:      aws.Int64(10),
		MessageAttributes: p.MessageAttributes,
		MessageBody:       aws.String(p.MessageBody),
		QueueUrl:          result.QueueUrl,
	})
	if err != nil {
		return err
	}
	return nil
}
