package db_helpers

// DBInterface abstracts common database operations.
type DBInterface interface {
	Connect(connectionString string) error
	Close() error
}

// DBManager implements the DBInterface
type DBManager struct{}

// Connect establishes a connection to the database.
func (d *DBManager) Connect(connectionString string) error {
	return ConnectPostgres(connectionString)
}

// Close closes the database connection.
func (d *DBManager) Close() error {
	return ClosePostgres()
}
