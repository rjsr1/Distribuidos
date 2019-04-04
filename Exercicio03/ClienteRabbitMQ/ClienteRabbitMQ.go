package main

import (
  //"encoding/json"
 	"fmt"
 	"os"
  	"log"
  	"bufio"
	"github.com/streadway/amqp"
	"math/rand"
	"strconv"
	"strings"
	// "time"
)
func qtdNumbers() int{
	return rand.Intn(10)
}

func joinMmcArgs(mmcArgs []string) string {
	var sb strings.Builder
	for _, r := range mmcArgs {
		sb.WriteString(r)
		sb.WriteString(",")
	}
	mmcArgsGenerated:=strings.Trim(sb.String(),",")
	return mmcArgsGenerated
}

//funcao para gerar uma string de entrada ex: "1,2,3,4..."
func mmcArgGenerator() string {
	mmcArgs:= make([] string,1000)	
	for i:=0;i<len(mmcArgs);i++{
		mmcArgs[i] = strconv.Itoa(i+1)
	}
	return joinMmcArgs(mmcArgs)
}

func deleteAllQueues(){
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()	  
	ch.QueueDelete("MMCArgs",false,false,false)
	ch.QueueDelete("ValidationArgs",false,false,false)
	ch.QueueDelete("validationResult",false,false,false)
	ch.QueueDelete("MMCResult",false,false,false)
}

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
	requestMMC, err := ch.QueueDeclare("MMCArgs", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)

	responseMMC, err := ch.QueueDeclare(
		"MMCResult", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgsFromServer, err := ch.Consume(
		responseMMC.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	go func(){
		for d := range msgsFromServer {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	//codigo para analise de desempenho
	//  fmt.Print("iniciando teste com 5000 entradas")
	//  totalTime:= 0.0
	//  for i := 0; i <5000; i++{
	// 		mmcArg := mmcArgGenerator()
	// 		//fmt.Println("Argumento ===============")
	// 		//fmt.Println(mmcArg)
	// 		t1 := time.Now()	
	// 		err = ch.Publish(
	// 					"",     // exchange
	// 					requestMMC.Name, // routing key
	// 					false,  // mandatory
	// 					false,  // immediate
	// 					amqp.Publishing{
	// 						ContentType: "text/plain",
	// 						Body:        []byte(mmcArg),
	// 					})

				
	// 	failOnError(err, "Falha ao publicar mmc args")  
					
	// 	<- msgsFromServer
	// 	t2 := time.Now()
	// 	x := float64(t2.Sub(t1).Nanoseconds()) / 1000000
	// 	totalTime= totalTime + x
	// 	fmt.Println(x)
	//  }
	//  	fmt.Print("Tempo Total = ")
	// 	fmt.Println(totalTime)

	 
    reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		
		text, _ := reader.ReadString('\n')  
		
    err = ch.Publish(
			"",     // exchange
			requestMMC.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(text),
			})
		failOnError(err, "Failed to publish mmc args")    
  }