package helpers_test

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"code.cloudfoundry.org/diego-db-helpers/sqldb/helpers"
	"code.cloudfoundry.org/lager/v3"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"testing"
)

func TestHelpers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Helpers Suite")
}

var (
	db                     *sql.DB
	ctx                    context.Context
	dbName                 string
	dbDriverName           string
	dbBaseConnectionString string
	dbFlavor               string
	tableName              string
	dbParams               *helpers.ConnectParams
)

func usePostgres() bool {
	return dbDriver() == "postgres"
}

func useMySQL() bool {
	d := dbDriver()
	return d == "mysql" || d == "mysql8"
}

func dbDriver() string {
	flavor := os.Getenv("DB")
	if flavor == "" {
		flavor = "postgres"
	}
	return flavor
}

var _ = BeforeEach(func() {
	dbName = fmt.Sprintf("diego_%d", GinkgoParallelProcess())

	if usePostgres() {
		dbDriverName = "postgres"
		user, ok := os.LookupEnv("DB_USER")
		if !ok {
			user = "diego"
		}
		password, ok := os.LookupEnv("DB_PASSWORD")
		if !ok {
			password = "diego_pw"
		}
		dbBaseConnectionString = fmt.Sprintf("postgres://%s:%s@localhost/", user, password)
		dbFlavor = helpers.Postgres
	} else if useMySQL() {
		dbDriverName = "mysql"
		user, ok := os.LookupEnv("DB_USER")
		if !ok {
			user = "diego"
		}
		password, ok := os.LookupEnv("DB_PASSWORD")
		if !ok {
			password = "diego_password"
		}
		dbBaseConnectionString = fmt.Sprintf("%s:%s@/", user, password)
		dbFlavor = helpers.MySQL
	} else {
		panic(fmt.Sprintf("unsupported DB driver: %q — set DB=postgres, mysql, or mysql8", dbDriver()))
	}

	logger := lager.NewLogger("helper-suite-test")

	var err error
	dbParams = &helpers.ConnectParams{
		DriverName:                    dbDriverName,
		DatabaseConnectionString:      dbBaseConnectionString,
		SqlCACertFile:                 "",
		SqlEnableIdentityVerification: false,
	}
	db, err = helpers.Connect(logger, dbParams)
	Expect(err).NotTo(HaveOccurred())
	Expect(db.Ping()).NotTo(HaveOccurred())

	ctx = context.Background()

	// Ensure that if another test failed to clean up we can still proceed
	db.ExecContext(ctx, fmt.Sprintf("DROP DATABASE %s", dbName))

	_, err = db.ExecContext(ctx, fmt.Sprintf("CREATE DATABASE %s", dbName))
	Expect(err).NotTo(HaveOccurred())

	Expect(db.Close()).To(Succeed())

	connStringWithDB := fmt.Sprintf("%s%s", dbBaseConnectionString, dbName)
	dbParams = &helpers.ConnectParams{
		DriverName:                    dbDriverName,
		DatabaseConnectionString:      connStringWithDB,
		SqlCACertFile:                 "",
		SqlEnableIdentityVerification: false,
	}
	db, err = helpers.Connect(logger, dbParams)
	Expect(err).NotTo(HaveOccurred())
	Expect(db.Ping()).NotTo(HaveOccurred())
})

var _ = AfterEach(func() {
	logger := lager.NewLogger("helper-suite-test")
	baseParams := &helpers.ConnectParams{
		DriverName:                    dbDriverName,
		DatabaseConnectionString:      dbBaseConnectionString,
		SqlCACertFile:                 "",
		SqlEnableIdentityVerification: false,
	}
	Expect(db.Close()).NotTo(HaveOccurred())
	db, err := helpers.Connect(logger, baseParams)
	Expect(err).NotTo(HaveOccurred())
	Expect(db.Ping()).NotTo(HaveOccurred())
	_, err = db.Exec(fmt.Sprintf("DROP DATABASE diego_%d", GinkgoParallelProcess()))
	Expect(err).NotTo(HaveOccurred())
	Expect(db.Close()).NotTo(HaveOccurred())
})
