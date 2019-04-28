package main

import(
	"encoding/json"
	"fmt"
	"log"
)

func demultiplexer(request operationRequest) string{
	log.Println("invoker - iniciando demultiplaxer")	
	response:= ""
	log.Println("args do request - "+request.Args[0].Key)
	args:=request.Args
	if request.Operation == "lookup" {
		log.Println("invoker - operacao foi lookup")	
		
		for _,a := range args{
			if a.Key == "ServiceName"{
				lookupResponse,err := lookup(a.Value)
				failOnError(err,"falha ao fazer lookup do serviço ")
				responsebytes, err := json.Marshal(lookupResponse)
				failOnError(err,"falha ao fazer marshall do lookupResponse")
				response= string(responsebytes)
			}else{
				response="Arg com formato nao conhecido, falha ao buscar arg serviceName"
			}
		}
	}

	if request.Operation == "bind" {
		log.Println("invoker - operacao foi bind")
		proxy:= clientProxy{}	
		registry:= nameRegistry{} 
		registry.Proxy=proxy
		for _,a := range args{
			if a.Key == "ServiceName"{
				registry.ServiceName = a.Value
			}
			if a.Key =="Port" {
				proxy.Port = a.Value
			}
			if a.Key == "IP"{
				proxy.IP=a.Value
			}
			if a.Key == "ID"{
				proxy.ID=a.Value
			}			
		}
		registry.Proxy=proxy
		bind(registry)
		response="serviço registrado"		
	}
	log.Println("invoker - response enviado será"+response)	
	return response
}

func invoker(){
	startNamingServer()
	log.Println("invoker - iniciando start Naming Server")	
	for {		
		request:= receiveRequest()
		log.Println("Invoker - recebido request")
		operation := operationRequest{}
		fmt.Println(request)
		err:=json.Unmarshal([]byte(request),&operation)
		failOnError(err,"falha ao fazer marshal do request")
		log.Println("Invoker - operation eh : "+operation.Operation)
		response:=demultiplexer(operation)
		sendResponse(response)
	}
}