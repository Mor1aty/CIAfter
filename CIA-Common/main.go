package main

import "moriaty.com/cia/cia-common/auth"

func main() {
	auth := auth.NewAuth("memory")
	auth.RemoveToken("1")
}
