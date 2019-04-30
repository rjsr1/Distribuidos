package main
import(
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"strconv"
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


func mdc(a int,b int) int{
	if a<b{
	  tempa:=a
	  a=b
	  b=tempa
	}
	for b!=0 {    
	  r := a%b 
	  a=b
	  b=r
	  // fmt.Printf("no mdc -> a = %d , b= %d , r=%d '\n' ",a,b,r)
	}
	return a;
  }
  func mmc(a int,b int) int{
	// fmt.Printf("no mmc -> a = %d , b= %d '\n' ",a,b)
	return a*(b/mdc(a,b))
  }

func demultiplexer(request messageBody) string{
	log.Println("invoker - iniciando demultiplaxer")	
	response:= ""
	log.Println("args do request - "+request.RequestBody.Parameters[0])
	
	if request.RequestHeader.Operation == "MMC" {
		log.Println("invoker - operacao foi MMC")	
		mmcArray := request.RequestBody.Parameters[0]
		values:=strings.Trim(mmcArray, "\r\n")     
		numbersReceived := strings.Split(values,",")    
    	arraySize:=len(numbersReceived)    
    	numbers :=make([]int, arraySize);

    	for a:=0;a<arraySize;a++ {    
     		 i,_ := strconv.Atoi(numbersReceived[a])     
      		numbers[a] = i     
    	}
    	mmcTotal:=1
    	if len(numbers)>1{
       		for i:=0;i<len(numbers);i++ {
        		mmcTotal = mmc(mmcTotal,numbers[i])      
      		}
    	}  
		response=strconv.Itoa(mmcTotal)
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