package kafka

import ckafka "github.com/confluentinc/confluent-kafka-go/kafka"

type KafkaProducer struct {
	Producer *ckafka.Producer
}


func (kp KafkaProducer) PublishMessage(msg, topic string) error {

	message := &ckafka.Message{
		TopicPartition: ckafka.TopicPartition{Topic: &topic, Partition: ckafka.PartitionAny},
		Value: []byte(msg),
	}


	if err := kp.Producer.Produce(message, nil); err != nil {
		return err
	}

	return nil
}

func (kp *KafkaProducer) SetupProducer(bootstrapServers string) {
	configMap := &ckafka.ConfigMap{
		"bootstrapServers": bootstrapServers,
	}

	kp.Producer, _ = ckafka.NewProducer(configMap)
}

func NewKafkaProducer() KafkaProducer {
	return KafkaProducer{}
}
