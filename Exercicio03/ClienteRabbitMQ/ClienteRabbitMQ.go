package main

import (
  //"encoding/json"
 	"fmt"
 	"os"
  "log"
 "bufio"
  "github.com/streadway/amqp"
)

// type MMCArgs struct{
// 	numbers string
// }

func failOnError(err error, msg string) {
	if err != nil {
	  log.Fatalf("%s: %s", msg, err)
		}
  }

  func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"MMCArgs", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)

	r, err := ch.QueueDeclare(
		"MMCResult", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		r.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	go func(){
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	for { 
   
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')  
		
    err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(text),
			})
		failOnError(err, "Failed to publish mmc args")    
  }
}