package db

import (
    "database/sql"
    "os"

    _ "github.com/lib/pq"
)

func DBConnection() (*sql.DB, error) {
    db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
    if err != nil {
        return nil, err
    }

    // Create the table if it doesn't exist
    _, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name TEXT, email TEXT)")
    if err != nil {
        return nil, err
    }

    return db, nil
}
