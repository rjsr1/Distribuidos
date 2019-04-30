package main
import(	
	"encoding/json"
	"log"
	"net"
)

var namingServerAdress="127.0.0.1:8081"
func bind(registry nameRegistry){

	connec,err := net.Dial("tcp",namingServerAdress)
	connection=connec
	log.Println("NamingProxy - fazendo marshal do registry")
	bytesToSend,err:=json.Marshal(registry)
	bytesToSend=append(bytesToSend,byte('\n'))
	log.Println("Naming Proxy - objeto serializado foi : "+string(bytesToSend))
	failOnError(err,"NamingProxi Server - erro ao fazer marshall de registry")
	send(bytesToSend)
	bytesReceived:=receive()
	response:=string(bytesReceived)
	log.Println("NamingProxi Server - response : "+response)
	
}


