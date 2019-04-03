package main

import "net"
import "fmt"
import "strings" 
import "strconv"


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

  LocalAddr, _ := net.ResolveUDPAddr("udp", "127.0.0.1:55951")
  connv,_ := net.DialUDP("udp",LocalAddr,&net.UDPAddr{IP:[]byte{127,0,0,1},Port:6082,Zone:""})
  
  buffer := make([]byte, 1024)
  bufferValidation :=make([]byte,1024)

  
  for {
    n,clientAddress,_ := pc.ReadFrom(buffer)  
    
    message := string(buffer[0:n]) 
    fmt.Fprintf(connv, message + "\n")	
    
    va,_,_ := connv.ReadFrom(bufferValidation) 
    validateMessage := string(bufferValidation[0:va])     
    if strings.Contains(validateMessage,"Formato invalido"){
      pc.WriteTo([]byte("Formato invalido"),clientAddress) 
    }else{
    
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
      
    pc.WriteTo([]byte(newmessage),clientAddress) 
  }
  }
}