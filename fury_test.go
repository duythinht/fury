package fury

import (
	"testing"
	"time"

	"github.com/gocql/gocql"
)

type Sold struct {
	UserID    int       `cass:"user_id"`
	Adv_ID    int       `cass:"ad_id"`
	Title     string    `cass:"title"`
	CreatedAt time.Time `cass:"created_at"`
}

func TestQueryUsingSession(t *testing.T) {

	cluster := gocql.NewCluster("10.60.3.13")
	cluster.Keyspace = "history"

	if sess, err := cluster.CreateSession(); err == nil {
		defer sess.Close()
		rows := QueryUsing(sess).CQL("SELECT user_id, ad_id, created_at, title FROM sold").Rows()

		solds := []Sold{}

		for rows.Next() {
			sold := new(Sold)
			rows.Scan(sold)
			solds = append(solds, *sold)
		}

		if solds[0].UserID != 1 {
			t.Fail()
		}
	}
}
