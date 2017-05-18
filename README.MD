## struct binding for cassandra query

```go
package main

import (
    "fmt"
    . "github.com/duythinht/fury"
}

// Sold is a struct, which will be binding to cassandra row query
type Sold struct {
	UserID    int       `cass:"user_id"`
	Adv_ID    int       `cass:"ad_id"`
	Title     string    `cass:"title"`
	CreatedAt time.Time `cass:"created_at"`
}

func main() {
  cluster := gocql.NewCluster("10.60.3.13")
	cluster.Keyspace = "history"

	if sess, err := cluster.CreateSession(); err == nil {
		defer sess.Close()
		rows := QueryUsing(sess).CQL("SELECT user_id, ad_id, created_at, title FROM sold").Rows()

		for rows.Next() {

			sold := new(Sold)
			rows.Scan(sold)
			fmt.Println(sold.UserID, sold.Title, sold.CreatedAt)
		}
	}
}
```
