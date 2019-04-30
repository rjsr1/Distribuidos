package main

//faz busca pelo serviço no servidor de nomes

var mmcProxy clientProxy


func calculateMMC(args [] string) string{

	
	inv:=invocation{}
	inv.Operation="MMC"
	inv.Args=args
	inv.Proxy=mmcProxy
	reply := invoke(inv)

	//faz chamada ao requestor
	return reply.result
}