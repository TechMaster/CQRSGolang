package service

import (
	"github.com/go-pg/pg"
	proto "github.com/TechMaster/microKafka/RestGRPCKafka/proto"
	"github.com/go-pg/pg/orm"
)

var Db *pg.DB

//TODO: get info from config
func ConnectDb() error {
	Db = pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "@123-",
		Database: "bank",
		Addr:     "localhost:5432",
	})


	var n int
	_, err := Db.QueryOne(pg.Scan(&n), "SELECT 1")
	return err
}

func InitSchema() {
	var account proto.Account
	// Tạo bảng
	for _, model := range []interface{}{&account} {
		err := Db.CreateTable(model, &orm.CreateTableOptions{
			Temp:          false,
			FKConstraints: true,
			IfNotExists:   true,
		})
		if err != nil {
			panic(err)
		}
	}


}
