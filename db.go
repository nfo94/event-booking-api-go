var DB *sql.DB

func InitDB(){
    var err error
    DB, err = sql.Open(sqlite3, "api.db")
    if err != nil {
        panic("Failed to connect to database")
    }

    DB.SetMaxOpenConns(10)
    DB.SetMaxIdleConns(5)

    createTables()
}