package main

import (
	"bufio"	
	"net"
	"log"
	"io"
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

func receive() [] byte{	

		message, err := bufio.NewReader(connection).ReadBytes('\n')
		log.Println("SRH - seguinte mesagem recebida - ")
		log.Println(message)
		failOnError(err,"falha ao receber dados")
		return message		
}
func send(response [] byte){
	log.Println("SRH - Enviando seguinte mensagem como resposta")
	log.Println(response)
	_,err := connection.Write(response)
	failOnError(err,"falha ao enviar resposta")
}