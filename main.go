package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/tester/api"
	db "github.com/tester/db/sqlc"
	"github.com/tester/util"
)

func main() {
	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal("an erros has detected when loading configuration")
	}

	conx, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("we couldn't connect to the database", err)
	}

	store := db.NewStore(conx)

	server, err := api.NewServer(store, config.ServerAddress, config)

	if err != nil {
		log.Fatal("we couldn't create a Server", err)
	}

	err = server.Start()

	if err != nil {
		log.Fatal("we couldn't start a Server", err)
	}

	/* var strBuffer bytes.Buffer
	strBuffer.WriteString("Ranjan")
	strBuffer.WriteString("Kumar")
	fmt.Println("The string buffer output is", strBuffer.String())
	fmt.Println("The string buffer output is", strBuffer.Bytes())
	fmt.Println("The string buffer output is", string(strBuffer.Bytes())) */

	/* var buffer bytes.Buffer
	buffer.Write([]byte("Hello, "))
	buffer.Write([]byte("World!"))
	result := buffer.Bytes()
	fmt.Println(buffer)
	fmt.Println(result) */

}
