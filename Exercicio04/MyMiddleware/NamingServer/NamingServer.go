package main

import (
	"errors"	
	"log"
)

type clientProxy struct{
	Port string
	IP string
	ID string
}
type arg struct{
	Key string
	Value string
}
type operationRequest struct {
	Operation string
    Args []arg
}


type nameRegistry struct{
	 ServiceName string
	 Proxy clientProxy
}
type lookupRegistry struct{
	Name string
}

var activeServices []nameRegistry;

func lookup(name string) (clientProxy,error){
	response:=clientProxy{}
	log.Println("NamingServer - buscando por "+name)
	for _,n := range activeServices{
		log.Println("NamingServer - loop nome : "+ n.ServiceName)
		if n.ServiceName == name {
			response = n.Proxy
			return response,nil	
		}
	}
	return response,errors.New("Serviço nao encontrado")	
}

func bind(registry nameRegistry){
	activeServices = append(activeServices,registry)
}

func list() []nameRegistry{
	return activeServices
}
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)		
	}
}

func main(){
	log.Println("iniciando servidor com chamada ao invoker")
	invoker()
}