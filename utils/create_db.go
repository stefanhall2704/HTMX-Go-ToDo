package main

import (
    "database/sql"
    "log"

    _ "github.com/mattn/go-sqlite3"
)

// TaskStatus represents the status of a task.
type TaskStatus struct {
    StatusID   int
    TaskStatus string
}

func main() {
    // Create or open the SQLite database file
    dbFile := "Tasks.db"
    db, err := sql.Open("sqlite3", dbFile)
    if err != nil {
        log.Fatal("Error creating or opening database:", err)
    }
    defer db.Close()

    // Create the Status table
    _, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS Status (
            ID INTEGER PRIMARY KEY AUTOINCREMENT,
            TaskStatus TEXT
        )
    `)
    if err != nil {
        log.Fatal("Error creating Status table:", err)
    }

    // Create the Task table with a foreign key constraint on StatusID
    _, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS Task (
            ID INTEGER PRIMARY KEY AUTOINCREMENT,
            Description TEXT,
            StatusID INTEGER,
            FOREIGN KEY (StatusID) REFERENCES Status(ID)
        )
    `)
    if err != nil {
        log.Fatal("Error creating Task table:", err)
    }

    log.Println("SQLite database with Status and Task tables created successfully.")

    // List of TaskStatus values to add
    statuses := []TaskStatus{
        {TaskStatus: "Not Started"},
        {TaskStatus: "In Progress"},
        {TaskStatus: "Complete"},
    }

    // Add each status to the Status table
    for _, status := range statuses {
        statusID, err := addStatus(db, status)
        if err != nil {
            log.Fatal("Error adding status:", err)
        }
        log.Printf("Status added with ID: %d\n", statusID)
    }
}

// addStatus inserts a new status into the Status table and returns the ID of the inserted status.
func addStatus(db *sql.DB, status TaskStatus) (int64, error) {
    result, err := db.Exec("INSERT INTO Status (TaskStatus) VALUES (?)", status.TaskStatus)
    if err != nil {
        return 0, err
    }
    statusID, err := result.LastInsertId()
    if err != nil {
        return 0, err
    }
    return statusID, nil
}
