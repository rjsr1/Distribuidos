package main

import (
  "fmt"
  "log"
  "net"
  "net/rpc"
  "time"
  "bufio"
  "os"
  "strings"
)

func main() {
  conn, err := net.DialTimeout("tcp", "localhost:8080", time.Minute)
  if err != nil {
    log.Fatal("dialing:", err)
  }

  client := rpc.NewClient(conn)

  var reply int
  for{
	  fmt.Printf("Digite sua lista de numeros separados por virgula: ")
	  
	  message, _ := bufio.NewReader(os.Stdin).ReadString('\n') 
	  
	  if strings.Trim(message, "\r\n") == "exit" {
	  
		os.Exit(1)
		}

	  e := client.Call("Service.MMC", &message, &reply)
	  if e != nil {
		log.Fatalf("Algo deu errado: %v", e.Error())
	  }else{
		if reply == -1{
			fmt.Printf("Lista de valores inválida. \n")
		}else{
	  fmt.Printf("Seu MMC é: %d \n", reply)}
	  }
	}
}