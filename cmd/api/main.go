package main

import (
	"github.com/DiegoUrrego4/Wheeper/cmd/server"
)

func main() {
	srv := server.NewServer()
	srv.Run()
}
