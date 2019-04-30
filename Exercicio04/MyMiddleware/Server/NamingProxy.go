package main
import(	
	"encoding/json"
	"log"
)

var namingServerAdress="127.0.0.1:8081"
func bind(registry nameRegistry){
	bytesToSend,err:=json.Marshal(registry)
	failOnError(err,"NamingProxi Server - erro ao fazer marshall de registry")
	send(bytesToSend)
	bytesReceived:=receive()
	response:=string(bytesReceived)
	log.Println("NamingProxi Server - response : "+response)
	
}
