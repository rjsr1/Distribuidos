package main

import (
	"log"
	"strings" 
	"strconv"
	"github.com/streadway/amqp"
)
func mdc(a int,b int) int{
	if a<b{
	  tempa:=a
	  a=b
	  b=tempa
	}
	for b!=0 {    
	  r := a%b 
	  a=b
	  b=r
	  // fmt.Printf("no mdc -> a = %d , b= %d , r=%d '\n' ",a,b,r)
	}
	return a;
  }
  func mmc(a int,b int) int{
	// fmt.Printf("no mmc -> a = %d , b= %d '\n' ",a,b)
	return a*(b/mdc(a,b))
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

	//declara queue para arg mmc
	q, err := ch.QueueDeclare(
		"MMCArgs", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue for MMCArgs")
	
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	validationArgsQueue, err := ch.QueueDeclare(
		"ValidationArgs", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue for ValidationArgs")

	validationResultQueue, err := ch.QueueDeclare(
		"validationResult", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue for validationResult")

	msgsValidation, err := ch.Consume(
		validationResultQueue.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")


	r, err := ch.QueueDeclare(
		"MMCResult", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	forever := make(chan bool)
	
	failOnError(err, "Failed to publish mmc args")    


	go func(){
		for valid := range msgsValidation{

			validateResult := string(valid.Body)
			log.Printf("Received a message: %s", valid.Body)
			mmcResult:=""
			if strings.Contains(validateResult,"Formato invalido"){
				mmcResult="formato enviado invalido."
			}else{

			
			message := strings.Trim(validateResult, "\r\n")
			numbersReceived := strings.Split(message,",")  
			arraySize:=len(numbersReceived) 
			numbers :=make([]int, arraySize);
			for a:=0;a<arraySize;a++ {    
				i,_ := strconv.Atoi(numbersReceived[a])     
				numbers[a] = i     
			  }

			  mmcTotal:=1
    		if len(numbers)>1{
     		  for i:=0;i<len(numbers);i++ {
     		mmcTotal = mmc(mmcTotal,numbers[i])      
				}	
			}
			mmcResult=strconv.Itoa(mmcTotal)
			}
			err = ch.Publish(
				"",     // exchange
				r.Name, // routing key
				false,  // mandatory
				false,  // immediate
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        []byte(mmcResult),
				})
			failOnError(err, "Failed to publish mmc args")  
		}
	}()

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)	

			err = ch.Publish(
				"",     // exchange
				validationArgsQueue.Name, // routing key
				false,  // mandatory
				false,  // immediate
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        d.Body,
				})
			log.Printf("Enviando ao servidor de validacao : %s", d.Body)
			failOnError(err, "Failed to publish mmc args")    
		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

}