package main

import (
	"testing"
)

func TestInitMigratesDB(t *testing.T) {
    db, mock := newMockDatabase()
    mock.ExpectExec("CREATE TABLE users(.*)")
    mock.ExpectCommit()
    
    // fails here
    initDB(db)
}