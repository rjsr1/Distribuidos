package main

import (
	"fmt"
	"net/rpc"
	"strings"
	"strconv"
	"net"
	"time"
	"log"
)

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

// Service is our RPC functions return type
type Service int

// Details is our exposed RPC function
func (l *Service) MMC(args *string, reply *int) error {
  fmt.Printf("Recebido: %+v\n", &args)  
  message:=strings.Trim(*args, "\r\n")
  
  conn, err := net.DialTimeout("tcp", "localhost:8081", time.Minute)
  if err != nil {
    log.Fatal("dialing:", err)
  }

  client := rpc.NewClient(conn)

  var reply1 bool  
    
  e := client.Call("Service.Valid", &message, &reply1)
  if e != nil {
		log.Fatalf("Algo deu errado: %v", e.Error())
  }else{
	  fmt.Printf("Recebido: %t\n", reply1)
	  if reply1 == true{
	  numbersReceived := strings.Split(message,",")
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
		}else{
		   mmcTotal = numbers[0]
		}
		fmt.Println("MMC calculado = ", mmcTotal)
		*reply = mmcTotal
		}else{
			*reply = -1
		}
	}
  return nil
}

func main() {
  service := new(Service)

  rpc.Register(service)

  listener, err := net.Listen("tcp", ":8080")
  if err != nil {
    // handle error
  }

  for {
    conn, err := listener.Accept()
    if err != nil {
      // handle error
    }

    go rpc.ServeConn(conn)
  }
}
