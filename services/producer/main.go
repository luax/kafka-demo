package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	kafka "github.com/segmentio/kafka-go"
)

const topic = "my-topic"
const kafkaURL = "kafka:9092"

var kafkaWriter *kafka.Writer

func handler(resp http.ResponseWriter, r *http.Request) {
	kafkaWriter.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("Key-A"),
			Value: []byte("Beep"),
		},
		kafka.Message{
			Key:   []byte("Key-B"),
			Value: []byte("Boop"),
		},
	)
	fmt.Fprintf(resp, "Produced messages")
}

func main() {
	// Kafka
	config := kafka.WriterConfig{
		Brokers:  []string{kafkaURL},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
	kafkaWriter = kafka.NewWriter(config)

	// Router
	r := mux.NewRouter()
	r.HandleFunc("/", handler)

	// Http
	http.Handle("/", r)
	fmt.Println("Listening")
	log.Fatal(http.ListenAndServe(":8080", nil))

	kafkaWriter.Close()
}
