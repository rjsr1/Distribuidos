package main
import(
	"log"
	"encoding/json"	
)

var namingServerAdress="127.0.0.1:8081"
func lookup(name string) clientProxy{
	operationReq:= operationRequest{}
	operationReq.Operation="lookup"
	args:=arg{"ServiceName","MMC_Calculator"}
	operationReq.Args = [] arg {args}
	requestLookup,err := json.Marshal(operationReq)
	failOnError(err,"Requestor - falha ao fazer marshal do operationRequest")
	setAddress("127.0.0.1","8081")	
	send(requestLookup)
	
	lookupResponse:=receive()
	log.Println(lookupResponse)

	clientProxyResult:=clientProxy{}
	err= json.Unmarshal(lookupResponse,&clientProxyResult)
	log.Println("Requestor - unmarhaller de "+clientProxyResult.ID)
	return clientProxyResult
}