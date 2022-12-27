package main

import (
	"fmt"
)

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

func main() {
	// Capture connection properties.

	varConfig := MyLoadJSON()

	fmt.Println("Apps Log path:     " + varConfig.AppLogPath)
	fmt.Println("DB Username:       " + varConfig.MySQLDBConn.DbUsername)
	fmt.Println("Cloud API Binary:  " + varConfig.MachineConfig.CloudAPIBin) 
	
	for i:=0; i< len(varConfig.CloudProviders) ; i++ {
		fmt.Println("NAME:            " + varConfig.CloudProviders[i].Name)
	  fmt.Println("Resource Group:    " + varConfig.CloudProviders[i].ResourceGroup)
	}
	
	for j:=0; j< len(varConfig.Consumers) ; j++ {
		fmt.Print("Consumer:            " + varConfig.Consumers[j] + ", ")
	}
	
	return
}