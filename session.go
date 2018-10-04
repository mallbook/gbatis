package gbatis

// SQLSession means sql connection
type SQLSession struct {
}

// SelectOne means select one record from database
func (s SQLSession) SelectOne(sqlID string, parameter interface{}) (interface{}, error) {
	return nil, nil
}

// SelectList means select multi records from database
func (s SQLSession) SelectList(sqlID string, parameter interface{}) ([]interface{}, error) {
	return nil, nil
}

// Execute means write operate
func (s SQLSession) Execute(sqlID string, parameter interface{}) error {
	return nil
}

// ExecuteMany means bulk write
func (s SQLSession) ExecuteMany(sqlID string, parameter []interface{}) error {
	return nil
}
