package main

import(
	"encoding/json"	
	"log"
)

type invocation struct {
	Proxy clientProxy
	Operation string
	Args []string
}
type replyInvocation struct{
	result string
}

type requestHeader struct{
	Context string
	RequestID int
	ResponseExpect string
	ObjectKey string
	Operation string
}


type requestBody struct{
	Parameters []string
}

type replyHeader struct{
	ID string
	Context string
	Status string

}

type replyBody struct{
	OperationResult string
}

type messageBody struct{
	RequestHeader requestHeader
	RequestBody requestBody
	ReplyHeader replyHeader
	ReplyBody replyBody
}
type operationRequest struct {
	Operation string
    Args []arg
}
type arg struct{
	Key string
	Value string
}



func invoke(inv invocation) replyInvocation{

	//cria a mensagem a ser transmitida
	
	reqBody:=requestBody{}
	reqBody.Parameters=inv.Args //no caso array de strings
	
	reqHeader:=requestHeader{}
	reqHeader.Operation=inv.Operation

	msgBody:=messageBody{}
	msgBody.RequestBody=reqBody
	msgBody.RequestHeader=reqHeader
	////
	msg,err:=json.Marshal(msgBody)
	log.Println("Requestor Client - fazendo marshall da mensagem")
	failOnError(err,"falha ao fazer marshal do msgbody")
	
	log.Println("Requestor Client - enviando mensagem ao CRH")
	setAddress(inv.Proxy.IP,inv.Proxy.Port)
	send(msg)


	msgResponse:=receive()
	
	log.Println("Requestor Client - Recebida mensagem do CRH")

	msgReplied:=messageBody{}
	err=json.Unmarshal(msgResponse,&msgReplied)
	failOnError(err,"falha ao fazer unmarshal do msgbody")
	replyInv:=replyInvocation{}

	if msgReplied.ReplyHeader.Status == "200"{
		replyInv.result=msgReplied.ReplyBody.OperationResult
	}else{
		replyInv.result="erro ao fazer solicitação" 
	}
	return replyInv

}