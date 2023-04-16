package repositories_test

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/ory/dockertest/v3"

	_ "github.com/go-sql-driver/mysql"
)

var (
	pool       *dockertest.Pool
	testDB     *sql.DB
	dbResource *dockertest.Resource
)

var (
	dbUser     = "testuser"
	dbPassword = "testpassword"
	dbName     = "local_video_station"
)

func TestMain(m *testing.M) {
	err := setup()
	if err != nil {
		os.Exit(1)
	}

	m.Run()

	teardown()
}

func setup() error {
	err := connectDB()
	if err != nil {
		log.Fatalf("could not connect database: %v", err)

		return err
	}

	if err := setupTestData(); err != nil {
		log.Fatalf("could not setup data: %v", err)

		return err
	}

	os.Setenv("IPADDRESS", "localhost")
	os.Setenv("MINIO_PORT", "9000")

	return nil
}

func connectDB() error {
	var err error
	pool, err = dockertest.NewPool("")

	if err != nil {
		log.Fatalf("could not start resource: %v", err)

		return err
	}

	dbResource, err = pool.Run(
		"mysql",
		"8.0.32",
		[]string{
			"MYSQL_ROOT_PASSWORD=admin",
			"MYSQL_DATABASE=" + dbName,
			"MYSQL_USER=" + dbUser,
			"MYSQL_PASSWORD=" + dbPassword,
		})

	if err != nil {
		log.Fatalf("could not start resource: %v", err)

		return err
	}

	if err = pool.Retry(func() error {
		var err error
		testDB, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, "localhost", dbResource.GetPort("3306/tcp"), dbName))
		if err != nil {
			return err
		}

		return testDB.Ping()
	}); err != nil {
		return err
	}

	return nil
}

func teardown() {
	if err := testDB.Close(); err != nil {
		log.Fatalf("could not close DB connection: %v", err)
	}

	if err := pool.Purge(dbResource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
}

func setupTestData() error {
	if err := executeSQLFile("./testdata/createDB.sql"); err != nil {
		return err
	}

	if err := executeSQLFile("./testdata/insertDB.sql"); err != nil {
		return err
	}

	return nil
}

func executeSQLFile(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("could not read file: %v", err)

		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var query string

	for scanner.Scan() {
		line := scanner.Text()

		query += line
		if strings.Contains(line, ";") {
			query = strings.TrimSpace(query)

			if _, err = testDB.Exec(query); err != nil {
				log.Fatalf("could not exec sql: %v", err)
			}

			query = ""
		}
	}

	return nil
}
