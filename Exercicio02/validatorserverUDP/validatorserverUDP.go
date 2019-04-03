package main

import (
	  "net"
    "fmt"  	
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
   
    
    pc,_:= net.ListenPacket("udp",":6082") 
    buffer := make([]byte, 1024)
    
    for {
      n,address,_ := pc.ReadFrom(buffer)  
      message := string(buffer[0:n]) 
     
      fmt.Print("Messagem Recebida:", string(message))
    
	  if last := len(message) - 1; last >= 0{
        message = message[:last]
    } 
    
	  if validate(message) == true {         
      pc.WriteTo([]byte(message + "\n"),address)	
	  }else{      
      pc.WriteTo([]byte("Formato invalido" + "\n"),address)	
	    }
    }
	
}
