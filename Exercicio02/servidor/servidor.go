package main

import "net"
import "fmt"
import "bufio"
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
  ln, _ := net.Listen("tcp", "127.0.0.1:8081")

  conn, _ := ln.Accept()
  connv,_ := net.Dial("tcp","127.0.0.1:8082")
  for {
    
    message, _ := bufio.NewReader(conn).ReadString('\n') 
    
    connv.Write([]byte(message))
    validateMessage, _ := bufio.NewReader(connv).ReadString('\n') 
    fmt.Println(validateMessage)
    if strings.Contains(validateMessage,"Formato invalido"){
      conn.Write([]byte("Formato invalido" + "\n"))
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
    
    conn.Write([]byte(newmessage + "\n"))
    }
  }
}