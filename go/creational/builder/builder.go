package builder

import (
	"errors"
)

type Database struct {
	name string
	host string
	port int
}

type DatabaseBuilder struct {
	database *Database
}

func NewDatabaseBuilder() *DatabaseBuilder {
	return &DatabaseBuilder{database: &Database{}}
}

func (b *DatabaseBuilder) WithName(name string) *DatabaseBuilder {
	b.database.name = name
	return b
}

func (b *DatabaseBuilder) WithUrl(host string, port int) *DatabaseBuilder {
	b.database.host = host
	b.database.port = port
	return b
}

func (b *DatabaseBuilder) Build() (*Database, error) {
	if !b.isValid() {
		return nil, errors.New("invalid database configuration")
	}
	return b.database, nil
}

func (b *DatabaseBuilder) isValid() bool {
	return b.database.name != ""
}
