package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/AryanMishra09/Bank/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB
var testStore *Store

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("can not load configs: ", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to DB")
	}

	testQueries = New(testDB)
	testStore = NewStore(testDB)

	os.Exit(m.Run())
}
