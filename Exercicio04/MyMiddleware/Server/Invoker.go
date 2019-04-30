package main
import(
	"encoding/json"
	"fmt"
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



func demultiplexer(request messageBody) string{
	log.Println("invoker - iniciando demultiplaxer")	
	response:= ""
	log.Println("args do request - "+request.RequestBody.Parameters[0])
	
	if request.RequestHeader.Operation == "MMC" {
		log.Println("invoker - operacao foi MMC")	
		mmcArray := request.RequestBody.Parameters[0]
		
		//calculo mmc
		response="12345"
	}

	
	return response
}

func invoker(){	
	log.Println("invoker - iniciando start Naming Server")	
	for {		
		request:= receive()
		log.Println("Invoker - recebido request")
		message := messageBody{}
		fmt.Println(request)
		err:=json.Unmarshal(request,&message)
		failOnError(err,"falha ao fazer unmarshal do request")
		log.Println("Invoker - operation eh : ")
		response:=demultiplexer(message)
		result,err:=json.Marshal(response+"\n")//adiciona delimitador de mensagem
		failOnError(err,"falha ao fazer marshal do response")
		send(result)
	}
}