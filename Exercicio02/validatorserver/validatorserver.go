package main

import (
	  "net"
    "fmt"
  	"bufio"
    "strings"
    "strconv"
)

func validate(A string) bool{
    A=strings.Trim(A, "\r\n")  
    strs := strings.Split(A, ",")
    ary := make([]int, len(strs))
    for i := range ary {
      ary[i], _ = strconv.Atoi(strs[i])      
		if ary[i] <= 0 {
		    return false
		}
    }
	
	return true
}

func main() {	
   
    ln, _ := net.Listen("tcp", ":8082")
    
    conn, _ := ln.Accept()
   
    for {
      
      message, _ := bufio.NewReader(conn).ReadString('\n')
     
      fmt.Print("Messagem Recebida:", string(message))
    
	  if last := len(message) - 1; last >= 0{
        message = message[:last]
		}
	  if validate(message) == true{    
    conn.Write([]byte(message + "\n"))	
	  }else{
	    newmessage := "Formato invalido"
		  conn.Write([]byte(newmessage + "\n"))
	    }
    }
	
}
