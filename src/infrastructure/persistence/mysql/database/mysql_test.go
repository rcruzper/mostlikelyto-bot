package database

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestMysql(t *testing.T) {
    t.Run("should create mysql data", func(t *testing.T) {
        mysql := New("username", "password", "host", "port", "schema")

        expectedMysql := &Mysql{
            username: "username",
            password: "password",
            host: "host",
            port: "port",
            schema: "schema",
        }

        assert.Equal(t, mysql, expectedMysql)
    })

    t.Run("should set default port", func(t *testing.T) {
        mysql := New("username", "password", "host", "", "schema")

        expectedMysql := &Mysql{
            username: "username",
            password: "password",
            host: "host",
            port: "3306",
            schema: "schema",
        }

        assert.Equal(t, mysql, expectedMysql)
    })

    t.Run("should create a correct connection string", func(t *testing.T) {
        mysql := New("username", "password", "host", "port", "schema")

        connectionString := mysql.connectionString()

        assert.Equal(t, connectionString, "username:password@tcp(host:port)/schema")
    })

}
