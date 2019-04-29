package main

import(
	"net"
	"log"
	"bufio"
)


var connection net.Conn

var port string
var ip string
func setAddress(ipp string,portp string){
	ip=ipp
	port=portp
}

func send(bytes []byte){
	log.Println("CRH - iniciando escuta no server request handler")

	connection,err := net.Dial("tcp", ip+":"+port)

	failOnError(err,"CRH - falha ao iniciar conexao")
	_,err = connection.Write(bytes)
}

func receive()[]byte{
	//espera uma quebra de linha ao fim da mensagem
	bytesReply, err := bufio.NewReader(connection).ReadBytes('\n')
	failOnError(err,"CRH - falha ao iniciar conexao")
	log.Println("CRH - retornando bytes")
	return bytesReply
}