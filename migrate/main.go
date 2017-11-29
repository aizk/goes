package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mattes/migrate/database/mysql"
	"github.com/mattes/migrate"
	_ "github.com/mattes/migrate/source/file"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "user:password@tcp(host:port)/dbname?multiStatements=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatal(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		// file:///absolute/path
		// file://relative/path
		"file://.",
		// 数据库名
		"mysql",
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(m.Version())
	//m.Steps(2)

	// Migrate all the way up ...
	if err := m.Up(); err != nil {
	    log.Fatal()
	}
}