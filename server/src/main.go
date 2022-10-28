package main

import (
	"log"

	"github.com/kuma0328/web3hackathon/infrastructure/database"
	"github.com/kuma0328/web3hackathon/presentation/router"
)

func main() {
	conn, err := database.NewConn()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := conn.DB.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// Routerの初期化
	r := router.NewRouter()

	r.Health()
	// r.NewTaskRouter(conn)

	// Routerの起動
	r.Serve()
}
