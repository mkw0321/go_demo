package Connection

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

// 发送消息
func writeByConn() {
	topic := "my-topic"
	partition := 0
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition) //连接至kafka集群的Learder结点
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}
	conn.SetWriteDeadline(time.Now().Add(10 * time.Second)) //设置消息发送的超时时间即截止时间
	//发送消息
	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte("one")},
		kafka.Message{Value: []byte("two")},
		kafka.Message{Value: []byte("three")},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}
	//关闭连接
	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}
