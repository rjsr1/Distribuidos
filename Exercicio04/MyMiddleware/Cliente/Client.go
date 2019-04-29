package main

import(
	"log"
)
const namingServerAddress string="127.0.0.1:8081"

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)		
	}
}

type clientProxy struct{
	Port string
	IP string
	ID string
}


// func lookup(serviceName string) clientProxy{
// 	//faz chamada ao requestor para o servico de nomes.
// 	return 
// }

func main(){
	resultTest:=calculateMMC([] string {"1,2,3,4,5,6,7,8,9,10"})
	log.Println("Client - recebendo este result")
	log.Println(resultTest)
}