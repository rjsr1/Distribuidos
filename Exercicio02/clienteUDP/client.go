package main

import "net"
import "fmt"
import "bufio"
import "os"


func main() {
  LocalAddr, _ := net.ResolveUDPAddr("udp", "127.0.0.1:55950")
  conn,_ := net.DialUDP("udp",LocalAddr,&net.UDPAddr{IP:[]byte{127,0,0,1},Port:6081,Zone:""})
 
  buffer := make([]byte, 1024)
  for { 
    
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Text to send: ")
    text, _ := reader.ReadString('\n')
    
    fmt.Fprintf(conn, text + "\n")	
    
    n,_,_ := conn.ReadFrom(buffer) 
    message := string(buffer[0:n]) 
    fmt.Println("Message from server: "+message)
  }
}