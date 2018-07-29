package main

import (
	"fmt"
)

/*
'go help' 							-> help
'godoc -http=:6060' 		-> deixando a doc offline
'go vet `file.go`' 			-> verifica falhas no codigo
'go build `file.go`'		-> compila o programa
'go run `file.go`'			-> compila e executa o programa
'go get -u github...'		-> instalando dependência, ex: go get -u github.com/go-sql-driver/mysql
*/
func main() {
	fmt.Printf("Outro programa em %s!", "Go")
}
