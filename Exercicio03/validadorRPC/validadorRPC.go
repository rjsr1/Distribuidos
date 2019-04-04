package main

import (
	"fmt"
	"net/rpc"
	"strings"
	"strconv"
	"net"
)

type Service bool

func (l *Service) Valid(args *string, reply *bool) error {
  fmt.Printf("Recebido: %+v\n", &args)
  message:=strings.Trim(*args, "\r\n")  
  numbersReceived := strings.Split(message,",")
  ary := make([]int, len(numbersReceived))
  for i := range ary {
	ary[i], _ = strconv.Atoi(numbersReceived[i])
	if ary[i] <= 0 {
		*reply = false
		break		
		}else{
		*reply = true
	}
  }
  return nil
}

func main() {
  service := new(Service)

  rpc.Register(service)

  listener, err := net.Listen("tcp", ":8081")
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
