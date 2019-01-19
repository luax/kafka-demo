package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	kafka "github.com/segmentio/kafka-go"
)

const topic = "my-topic"
const kafkaURL = "kafka:9092"

var kafkaReader *kafka.Reader

func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Waiting for message")
	m, err := kafkaReader.ReadMessage(context.Background())
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, "Error reading message")
		return
	}
	s := fmt.Sprintf(
		"%d: %s = %s\n",
		m.Offset,
		string(m.Key),
		string(m.Value),
	)
	fmt.Printf("Got message: %s\n", s)
	fmt.Fprintf(w, s)
}

func main() {
	// Kafka
	config := kafka.ReaderConfig{
		Brokers:  []string{kafkaURL},
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
		MaxWait:  1 * time.Second,
	}
	kafkaReader = kafka.NewReader(config)

	// Router
	r := mux.NewRouter()
	r.HandleFunc("/", handler)

	// Http
	http.Handle("/", r)
	fmt.Println("Listening")
	log.Fatal(http.ListenAndServe(":8080", nil))

	kafkaReader.Close()
}
