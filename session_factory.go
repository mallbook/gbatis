package gbatis

import (
	"sync"
)

// OpenSession means Open a SQL Session
func OpenSession() *SQLSession {
	return nil
}

type SQLSessionFactory struct{}

var (
	sessionFactory *SQLSessionFactory
	once           sync.Once
)

// SessionFactory means SQLSession Factory
func SessionFactory() *SQLSessionFactory {
	once.Do(func() {
		sessionFactory = &SQLSessionFactory{}
	})
	return sessionFactory
}
