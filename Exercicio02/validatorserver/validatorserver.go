package main

import (
	"net"
    "fmt"
	"bufio"
    "strings"
    "strconv"
	"reflect"
)

func validate(A string) bool{
	
	var V bool
	V = true

    strs := strings.Split(A, ",")
    ary := make([]int, len(strs))
    for i := range ary {
	    ary[i], _ = strconv.Atoi(strs[i])
		if ary[i] < 0 {
		    V = false
		}else if reflect.TypeOf(ary[i]).String() !=  "int" {
			V = false
		}
    }
	return V
}

func main() {
	
    // listen on all interfaces
    ln, _ := net.Listen("tcp", ":8081")
    // accept connection on port
    conn, _ := ln.Accept()
    // run loop forever (or until ctrl-c)
    for {
      // will listen for message to process ending in newline (\n)
      message, _ := bufio.NewReader(conn).ReadString('\n')
      // output message received
      fmt.Print("Messagem Recebida:", string(message))
      // sample process for string received
	  if last := len(message) - 1; last >= 0{
        message = message[:last]
		}
	  if validate(message) == true{
		conn,_ := net.Dial("tcp", "127.0.0.1:8082")
		fmt.Fprintf(conn, message + "\n")
	  }else{
	    newmessage := "Formato invalido"
		conn.Write([]byte(newmessage + "\n"))
	  }
    }
	
}