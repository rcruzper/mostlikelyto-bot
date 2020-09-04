package database

import (
    "database/sql"
    "fmt"
)

type Mysql struct {
    username string
    password string
    host string
    port string
    schema string
}

func New(username string, password string, host string, port string, schema string) *Mysql {
    return &Mysql{
        username: username,
        password: password,
        host:     host,
        port:     validatePort(port),
        schema:   schema,
    }
}
func validatePort(port string) string{
    if len(port) == 0 {
        return "3306"
    }
    return port

}

func (m *Mysql) connectionString() string {
    return fmt.Sprintf(
        "%s:%s@tcp(%s:%s)/%s", m.username, m.password, m.host, m.port, m.schema,
    )
}

func (m *Mysql) Connect() *sql.DB {
    db, err := sql.Open("mysql", m.connectionString())

    if err != nil {
        panic(err.Error())
    }

    return db
}
