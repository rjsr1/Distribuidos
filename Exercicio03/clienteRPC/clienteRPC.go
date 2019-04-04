package main

import (
  "fmt"
  "math/rand"
  "strconv"
  "log"
  "net"
  "net/rpc"
  "time"
  //"bufio"
  // "os"
	"strings"
	
)

func qtdNumbers() int{
	return rand.Intn(10)
}

func joinMmcArgs(mmcArgs []string) string {
	var sb strings.Builder
	for _, r := range mmcArgs {
		sb.WriteString(r)
		sb.WriteString(",")
	}
	mmcArgsGenerated:=strings.Trim(sb.String(),",")
	return mmcArgsGenerated
}

//funcao para gerar uma string de entrada ex: "1,2,3,4..."
func mmcArgGenerator() string {
	mmcArgs:= make([] string,1000)	
	for i:=0;i<len(mmcArgs);i++{
		mmcArgs[i] = strconv.Itoa(i+1)
	}
	return joinMmcArgs(mmcArgs)
}

func main() {
  conn, err := net.DialTimeout("tcp", "localhost:8080", time.Minute)
  if err != nil {
    log.Fatal("dialing:", err)
  }

  client := rpc.NewClient(conn)

  var reply int
  
	totalTime:= 0.0

  for i := 0; i <5000; i++{
	  //fmt.Printf("Digite sua lista de numeros separados por virgula: ")
	  
	  //message, _ := bufio.NewReader(os.Stdin).ReadString('\n') 
	  message := mmcArgGenerator()
	  //fmt.Printf(message)
	  //fmt.Printf("\n")
	  
	  // if strings.Trim(message, "\r\n") == "exit" {
	  
		// os.Exit(1)
		// }
		t1 := time.Now()
		e := client.Call("Service.MMC", &message, &reply)		
	  if e != nil {
		log.Fatalf("Algo deu errado: %v", e.Error())
	  }else{
		if reply == -1{
			fmt.Printf("Lista de valores inválida. \n")
		}else{

			t2 := time.Now()		
			x := float64(t2.Sub(t1).Nanoseconds()) / 1000000
			totalTime= totalTime + x
			fmt.Println("totalTime")
			fmt.Println(totalTime)
			//fmt.Printf("Seu MMC é: %d \n", reply)}			
	  	}
		}	
	}	
	fmt.Println("tempo total acima...")
	fmt.Println(totalTime)

}
