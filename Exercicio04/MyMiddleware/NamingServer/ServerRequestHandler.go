package main

import (
	"bufio"	
	"net"
	"log"
)
var connection net.Conn

func startNamingServer(){
	log.Println("SRH - iniciando escuta no server request handler")
	 
	//ouvindo na porta 8081
	ln, err := net.Listen("tcp", "127.0.0.1:8081")
	log.Println("SRH - apos net listen")
	failOnError(err,"SRH - falha ao iniciar escuta")
	
	conn, err := ln.Accept()
	log.Println("SRH - apos net accept")
	connection = conn
	failOnError(err,"SRH - falha ao receber conexão")

}

func receiveRequest() string{	

		message, err := bufio.NewReader(connection).ReadString('\n') 
		log.Println("SRH - seguinte mesagem recebida - ")
		log.Println(message)
		failOnError(err,"falha ao receber dados")
		return message		
}
func sendResponse(response string){
	log.Println("SRH - Enviando seguinte mensagem como resposta")
	log.Println(response)
	_,err := connection.Write([]byte(response + "\n"))
	failOnError(err,"falha ao enviar resposta")
}