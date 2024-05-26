package data

import (
	"database/sql"
	"fmt"

	"github.com/joaodrts/Drts.Ads.GO.API/config"
	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) {

	fmt.Printf("antes conectar banco")
	conf := config.GetDb()

	stringConnection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Password, conf.Database)

	conn, err := sql.Open("postgres", stringConnection)

	fmt.Printf("antes conectar banco -panic")

	if err != nil {
		panic(err)
	}

	err = conn.Ping()

	fmt.Printf("antes conectar banco - depoois panic")

	return conn, err
}
