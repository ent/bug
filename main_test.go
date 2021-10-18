package bug

import (
	"context"
	"entgo.io/bug/ent/enttest"
	"entgo.io/ent/dialect"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"net"
	"strconv"
	"testing"
)

func TestBugSQLite(t *testing.T) {
	ctx := context.Background()
	client := enttest.Open(t, dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	// Run schema migration.
	if err := client.Schema.Create(ctx); err != nil {
		t.Fatal(err)
	}
	client.User.Create().SetName("Ariel").SetAge(30).ExecX(ctx)
	if n := client.User.Query().CountX(ctx); n != 1 {
		t.Errorf("unexpected number of users: %d", n)
	}
}

func TestBugMySQL(t *testing.T) {
	ctx := context.Background()
	for version, port := range map[string]int{"56": 3306, "57": 3307, "8": 3308} {
		addr := net.JoinHostPort("localhost", strconv.Itoa(port))
		t.Run(version, func(t *testing.T) {
			client := enttest.Open(t, dialect.MySQL, fmt.Sprintf("root:pass@tcp(%s)/test?parseTime=True", addr))
			defer client.Close()
			// Run schema migration.
			if err := client.Schema.Create(ctx); err != nil {
				t.Fatal(err)
			}
			client.User.Create().SetName("Ariel").SetAge(30).ExecX(ctx)
			if n := client.User.Query().CountX(ctx); n != 1 {
				t.Errorf("unexpected number of users: %d", n)

			}
		})
	}
}

func TestBugPostrgreSQL(t *testing.T) {
	ctx := context.Background()
	for version, port := range map[string]int{"10": 5430, "11": 5431, "12": 5432, "13": 5433, "14": 5434} {
		t.Run(version, func(t *testing.T) {
			client := enttest.Open(t, dialect.Postgres, fmt.Sprintf("host=localhost port=%d user=postgres dbname=test password=pass sslmode=disable", port))
			defer client.Close()
			// Run schema migration.
			if err := client.Schema.Create(ctx); err != nil {
				t.Fatal(err)
			}
			client.User.Create().SetName("Ariel").SetAge(30).ExecX(ctx)
			if n := client.User.Query().CountX(ctx); n != 1 {
				t.Errorf("unexpected number of users: %d", n)

			}
		})
	}
}

func TestBugMaria(t *testing.T) {
	ctx := context.Background()
	for version, port := range map[string]int{"10.5": 4306, "10.2": 4307, "10.3": 4308} {
		t.Run(version, func(t *testing.T) {
			addr := net.JoinHostPort("localhost", strconv.Itoa(port))
			client := enttest.Open(t, dialect.MySQL, fmt.Sprintf("root:pass@tcp(%s)/test?parseTime=True", addr))
			defer client.Close()
			// Run schema migration.
			if err := client.Schema.Create(ctx); err != nil {
				t.Fatal(err)
			}
			client.User.Create().SetName("Ariel").SetAge(30).ExecX(ctx)
			if n := client.User.Query().CountX(ctx); n != 1 {
				t.Errorf("unexpected number of users: %d", n)
			}
		})
	}
}
