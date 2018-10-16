package gbatis

import (
	"database/sql"
)

type sqlSessionFactory struct {
	connPools map[string]*sql.DB
}

var sessionFactory = &sqlSessionFactory{
	connPools: make(map[string]*sql.DB),
}

/*
var (
	sessionFactory *SQLSessionFactory
	once           sync.Once
)

// Instance means SQLSession Factory
func Instance() *SQLSessionFactory {
	once.Do(func() {
		sessionFactory = &SQLSessionFactory{}
	})
	return sessionFactory
}
*/
