package db

import (
	"database/sql"
	"log/slog"

	_ "github.com/mattn/go-sqlite3"
)

var sqliteConn *sql.DB

func MustSetup(dsn string, logger *slog.Logger) {
	var err error
	sqliteConn, err = sql.Open("sqlite3", dsn)
	sqliteConn.SetMaxOpenConns(1)

	if err != nil {
		panic(err)
	}
	if err = sqliteConn.Ping(); err != nil {
		panic(err)
	}

	err = create_tables(logger)
	if err != nil {
		panic("Error configuring database")
	}

	logger.Info("Database is ready")
}

func create_table(create_table_sql string, logger *slog.Logger) error {
	_, err := sqliteConn.Exec(create_table_sql)
	if err != nil {
		logger.Error("Error when trying to prepare statement during creating tables")
		logger.Error(err.Error())
		return err
	}

	return nil
}

func create_tables(logger *slog.Logger) error {
	create_user_table_sql := `CREATE TABLE IF NOT EXISTS user (
		id VARCHAR PRIMARY KEY,
		tgid INTEGER,
		name VARCHAR,
		tgusername VARCHAR,
		chatid VARCHAR,
		createdat VARCHAR,
		updatedat VARCHAR,

		UNIQUE(tgid)
	);`

	create_event_table_sql := `CREATE TABLE IF NOT EXISTS event (
		id VARCHAR PRIMARY KEY,
		chatid VARCHAR,
		ownerid INTEGER,
		text VARCHAR,
		notifyat VARCHAR,
		delta VARCHAR,
		createdat VARCHAR,
		updatedat VARCHAR
	);`

	for _, table := range []string{
		create_user_table_sql,
		create_event_table_sql,
	} {
		err := create_table(table, logger)
		if err != nil {
			return err
		}
	}

	return nil
}
