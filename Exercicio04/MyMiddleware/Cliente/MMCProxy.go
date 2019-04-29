package main

//faz busca pelo serviço no servidor de nomes

func calculateMMC(args [] string) string{


	proxy :=clientProxy {"127.0.0.1","8082","123"}
	inv:=invocation{}
	inv.Operation="MMC"
	inv.Args=args
	inv.Proxy=proxy
	reply := invoke(inv)

	//faz chamada ao requestor
	return reply.result
}