package main

import (
		"log"
	
)
type clientProxy struct{
	Port string
	IP string
	ID string
}

type nameRegistry struct{
	ServiceName string
	Proxy clientProxy
}
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)		
	}
}
func main(){
	registryInfo:=nameRegistry{}
	registryInfo.Proxy.Port="8082"
	registryInfo.Proxy.IP="127.0.0.1"
	registryInfo.Proxy.ID="123"
	registryInfo.ServiceName="MMC_Calculator"	
	bind(registryInfo)

	

	


}