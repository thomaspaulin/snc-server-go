package snc

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

var TestDB *sql.DB

// TODO write the insert queries to set up the database and determine the expected values
// TODO use database migrations to create and drop the databases. e.g. goose https://github.com/pressly/goose

// TestMain runs all the tests for the rink file. It requires that a database named "snc_test" exists.
func TestMain(m *testing.M) {
	// do the setup
	var retCode int
	// run the tests
	retCode = m.Run()
	// do the tear down

	os.Exit(retCode)
}

func TestFetchAllRinks(t *testing.T) {
	//rinks, err := FetchRinks(DB)
	//exp := []
	//ok(t, err)
	//equals(t, exp, rinks)
	log.Println("Fetch All")
}

// helper functions thanks to https://github.com/benbjohnson/testing
// assert fails the test if the condition is false.
func assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: "+msg+"\033[39m\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	}
}

// ok fails the test if an err is not nil.
func ok(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
		tb.FailNow()
	}
}

// equals fails the test if exp is not equal to act.
func equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}
