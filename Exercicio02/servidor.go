package main

import "net"
import "fmt"
import "bufio"
import "strings" 
import "strconv"// only needed below for sample processing


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
    fmt.Printf("no mdc -> a = %d , b= %d , r=%d '\n' ",a,b,r)
  }
  return a;
}
func mmc(a int,b int) int{
  fmt.Printf("no mmc -> a = %d , b= %d '\n' ",a,b)
  return a*(b/mdc(a,b))
}



func main() {
 // conn2,_ := net.Dial("tcp","127.0.0.1:8082")
  fmt.Println("Launching server...")

  // listen on all interfaces
  ln, _ := net.Listen("tcp", ":8081")

  // accept connection on port
  conn, _ := ln.Accept()

  // run loop forever (or until ctrl-c)
  for {
    // will listen for message to process ending in newline (\n)
    message, _ := bufio.NewReader(conn).ReadString('\n') 
    fmt.Println("mensagem : ")
    fmt.Println(message)
    message=strings.Trim(message, "\r\n")      
    fmt.Println("apos trim : ")
    fmt.Println(message)
    numbersReceived := strings.Split(message,",")
    fmt.Println("numbers Received : ")
    fmt.Println(numbersReceived)
    fmt.Println("tamanho array : ")
    arraySize:=len(numbersReceived)
    fmt.Println(arraySize)
    numbers :=make([]int, arraySize);
    for a:=0;a<arraySize;a++ {
      fmt.Println("numero recebido na posicao a")
      fmt.Println(numbersReceived[a])
      i,_ := strconv.Atoi(numbersReceived[a])
      fmt.Println("numero apos conversao")
      fmt.Println(i)
      numbers[a] = i
      fmt.Println(numbers)
    }
    mmcTotal:=1
    if len(numbers)>2{
       for i:=0;i<len(numbers);i++ {
        mmcTotal = mmc(mmcTotal,numbers[i])
        fmt.Printf("no main -> i = %d,numero = %d , mmcTotal= %d '\n' ",i,numbers[i],mmcTotal)
      }
    }  
    // output message received
    fmt.Println("MMC calculado = ", mmcTotal)
    //fmt.Fprintf(conn2," enviando mensagem ao servidor 02 "+string(message)+ "\n")
    // sample process for string received
    newmessage := strings.ToUpper(message)
    // send new string back to client
    conn.Write([]byte(newmessage + "\n"))
  }
}