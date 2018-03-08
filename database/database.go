package database

//Read returns a value from the database
func Read() ([]byte, error) {
	return []byte("hello from db"), nil
}
