package gbatis

type sqlAction int8

const (
	selectAction    sqlAction = 0
	insertAction    sqlAction = 1
	updateAction    sqlAction = 2
	deleteAction    sqlAction = 3
	procStoreAction sqlAction = 3
)

// NamedSQLInfo means NamedSQL infomation
type namedSQL struct {
	action sqlAction
	sql    string
}

// NamedSQL include all namedSQL
var NamedSQL = make(map[string]namedSQL)
