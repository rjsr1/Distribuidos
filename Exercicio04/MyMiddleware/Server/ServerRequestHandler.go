package main

import (
	"bufio"	
	"net"
	"log"
	
)
var connection net.Conn




func startServer(){
	log.Println("SRH - iniciando escuta no server request handler")
	 
	//ouvindo na porta 8081
	ln, err := net.Listen("tcp", "127.0.0.1:8082")
	log.Println("SRH - apos net listen")
	failOnError(err,"SRH - falha ao iniciar escuta")
	
	conn, err := ln.Accept()
	log.Println("SRH - apos net accept")
	connection = conn
	failOnError(err,"SRH - falha ao receber conexão")

}

func receive() [] byte{	

		message, err := bufio.NewReader(connection).ReadBytes('\n')
		log.Println("SRH - seguinte mesagem recebida - ")
		log.Println(message)
		failOnError(err,"falha ao receber dados")
		return message		
}
func send(sendData [] byte){
	log.Println("SRH - Enviando seguinte mensagem"+string(sendData))
	log.Println("SHR - enviando sendData"+string(sendData))
	log.Println(connection.RemoteAddr())
	_,err := connection.Write(sendData)
	failOnError(err,"falha ao enviar resposta")
}