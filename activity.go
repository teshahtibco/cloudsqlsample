package CloudSqlConnection

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {
	// Get the activity data from the context
	host := context.GetInput("hostname").(string)
	port := context.GetInput("port").(string)
	user := context.GetInput("username").(string)
	pwd := context.GetInput("password").(string)
	instance := context.GetInput("instance").(string)
	s := []string{user, ":", pwd, "@tcp(", host, ":", port, ")/", instance}
	url := strings.Join(s, "")
	// do eval
	db, err := sql.Open("mysql", url)
	if err != nil {
		fmt.Printf("hello, world inside Error\n")
		log.Fatal(err)
		return false, err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Printf("hello, world inside DB Error\n")
		return false, err
	} else {
		fmt.Printf("hello, world inside DB Success\n")
	}
	return true, nil
}
