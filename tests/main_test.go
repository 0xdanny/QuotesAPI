package tests

// import (
// 	"log"
// 	"os"
// 	"testing"

// 	"github.com/0xdanny/QuotesAPI/server"
// 	"github.com/0xdanny/QuotesAPI/server/database"
// )

// const (
// 	DB_URL    = "postgres://zvlvdtht:Zb-Y6Dub_6cGmUL7voZh_Ehd3XXftF6q@rogue.db.elephantsql.com:5432/zvlvdtht"
// 	DB_DRIVER = "postgres"
// )

// const tableCreationQuery = `CREATE TABLE IF NOT EXISTS plines
// (
//     id SERIAL,
//     line TEXT NOT NULL,
//     CONSTRAINT plines_pkey PRIMARY KEY (id)
// )`

// var s *server.Server

// func TestMain(m *testing.M) {
// 	s = server.NewServer()
// 	s.DB = &database.DB{}
// 	err := s.DB.Open()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer s.DB.Close()

// 	ensureTableExists()
// 	exitCode := m.Run()
// 	clearTable()
// 	os.Exit(exitCode)
// }

// func ensureTableExists() {
// 	if _, err := s.DB.Exec(tableCreationQuery); err != nil {
// 		log.Fatal(err)
// 	}
// }

// func clearTable() {

// 	s.DB.Exec("DELETE FROM plines")
// 	s.DB.Exec("ALTER SEQUENCE plines_id_seq RESTART WITH 1")
// }
