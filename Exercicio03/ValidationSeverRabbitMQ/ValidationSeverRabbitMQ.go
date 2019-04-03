package main

import (
	"log"	
    "strings"
    "strconv"
	"github.com/streadway/amqp"
)

func validate(A string) bool{
    A=strings.Trim(A, "\r\n")  
    strs := strings.Split(A, ",")
    ary := make([]int, len(strs))
    for i := range ary {
      ary[i], _ = strconv.Atoi(strs[i])      
		if ary[i] <= 0 {
		    return false
		}
    }
	
	return true
}
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

	qValidationArgs, err := ch.QueueDeclare(
		"ValidationArgs", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	qValidationResult, err := ch.QueueDeclare(
		"validationResult", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgsValidationArgs, err := ch.Consume(
		qValidationArgs.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")
	forever := make(chan bool)
	
	go func() {
		for valid := range msgsValidationArgs{
			log.Printf("Received a message: %s", valid.Body)
			receivedArgs:= string(valid.Body)
			validateResult:=""
			isValid:=validate(receivedArgs)
			if isValid {
				validateResult=receivedArgs
			}else{
				validateResult="Formato invalido"
			}
			err = ch.Publish(
				"",     // exchange
				qValidationResult.Name, // routing key
				false,  // mandatory
				false,  // immediate
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        []byte(validateResult),
				})
			failOnError(err, "Failed to publish mmc args")  

		}
	}()
	<-forever
}