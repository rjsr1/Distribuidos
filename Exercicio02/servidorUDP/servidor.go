package main

import "net"
import "fmt"
import "strings" 
import "strconv"

type UDPServer struct{
  endereco string
  server *net.UDPConn
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



func main() {
 
  fmt.Println("Launching server...")
  
  pc,_:= net.ListenPacket("udp",":6081") 
  
  buffer := make([]byte, 1024)

  
  for {
    
    
    n,address,_ := pc.ReadFrom(buffer)  

    message := string(buffer[0:n]) 
    
    message=strings.Trim(message, "\r\n")      
    
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
    }      
    fmt.Println("MMC calculado = ", mmcTotal)   
    newmessage := strconv.Itoa(mmcTotal)
   
    fmt.Println(address)
    pc.WriteTo([]byte(newmessage),address) 
  }
}