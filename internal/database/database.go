package database

import (
	"database/sql"
	"time"
	_ "github.com/lib/pq"
)

type Log struct {
	Id int
	Time time.Time
	Username string
	Request string
	Answer float64
}

type Database interface {
	Close()
	AddLog(log Log)
	GetLogsByUsername(username string) ([]Log, error)
}

type postgresDatabase struct{
	wrappee *sql.DB
}

func NewPostgresDatabase(user string, password string, dbname string, sslmodeEnable bool) (postgresDatabase, error) {
	sslmode := ""
	if sslmodeEnable {
		sslmode = "enable"
	} else {
		sslmode = "disable"
	}

	db, err := sql.Open("postgres", "user="+user+" "+"password="+password+" dbname="+dbname+" sslmode="+sslmode)

	if err != nil {
		return postgresDatabase{}, err
	}

	postgresDb := postgresDatabase{db}
	postgresDb.createLogTableIfNotExists()

	return postgresDb, nil
}

func (db postgresDatabase) createLogTableIfNotExists() {
	_, err := db.wrappee.Exec(`CREATE TABLE IF NOT EXISTS calco_logs(
						id SERIAL PRIMARY KEY,
						time TIMESTAMP,
						username TEXT,
						request TEXT,
						answer REAL
					)`)
	if err != nil {
		panic(err)
	}
}

func (db postgresDatabase) Close() {
	db.wrappee.Close()
}

func (db postgresDatabase) AddLog(log Log) {
	_, err := db.wrappee.Exec(`INSERT INTO calco_logs 
				VALUES(NEXTVAL('calco_logs_id_seq'), $1, $2, $3, $4)`, 
						log.Time, log.Username, log.Request, log.Answer)

	if err != nil {
		panic(err)
	}
}

func (db postgresDatabase) GetLogsByUsername(username string) ([]Log, error) {
	rows, err := db.wrappee.Query("SELECT * FROM calco_logs WHERE username=$1", username)

	if err != nil {
		return nil, err
	}

	logs := make([]Log, 0 )
	for rows.Next() {
		log := Log{}

		err := rows.Scan(&log.Id, &log.Time, &log.Username, &log.Request, &log.Answer)

		if err != nil {
			return nil, err
		}

		logs = append(logs, log)
	}

	return logs, nil
}