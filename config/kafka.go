package config

import "os"

type kafka struct {
	BootstrapServers string
	GroupId          string
	Topics           topic
}

type topic struct {
	LockUser          string
	UserDeleteAccount string
}

var singleKafka *kafka

func InitKafka() *kafka {
	if singleKafka != nil {
		return singleKafka
	}

	lock.Lock()
	defer lock.Unlock()

	singleKafka = &kafka{
		BootstrapServers: os.Getenv("KAFKA_SERVERS"),
		GroupId:          os.Getenv("KAFKA_CLIENT_ID"),
		Topics: topic{
			LockUser:          "block-user",
			UserDeleteAccount: "user-delete-account",
		},
	}

	return singleKafka
}
