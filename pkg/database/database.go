package database

import (
	"database/sql"
	"job-visualizer/pkg/structs"
	"log"
	"os"
)

func CreateDatabase() *sql.DB {
	os.Remove("job_data.sqlite")
	db, err := sql.Open("sqlite", "job_data.sqlite")
	checkError(err)
	// defer db.Close()
	return db
}

func SetupDatabase(db *sql.DB) {
	createMainTable(db)
	createSecondaryTables(db)
}

func WriteToDatabase(db *sql.DB, allJobData []structs.JobData) {
	insertQueryJobData := `INSERT INTO job_data (location, job_title, company_name, description, date_posted, salary,
		work_from_home, qualifications, links, country) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`
	insertQueryQualifications := `INSERT OR IGNORE INTO qualifications (id, qualifications) VALUES (?, ?);`
	insertQueryLinks := `INSERT OR IGNORE INTO links (id, links) VALUES (?, ?);`
	for _, job := range allJobData {
		result, err := db.Exec(insertQueryJobData, job.Location, job.JobTitle, job.CompanyName, job.Description, job.DatePosted,
			job.Salary, job.WorkFromHome, job.Qualifications, job.Links, job.Country)
		checkError(err)
		id, err := result.LastInsertId()
		checkError(err)
		_, err = db.Exec(insertQueryQualifications, id, job.Qualifications)
		checkError(err)
		_, err = db.Exec(insertQueryLinks, id, job.Links)
		checkError(err)
	}
}

func createMainTable(db *sql.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS job_data(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		location TEXT NOT NULL,
		job_title TEXT NOT NULL,
		company_name TEXT NOT NULL,
		description TEXT,
		date_posted TEXT NOT NULL,
		salary INT,
		work_from_home TEXT,
		qualifications TEXT,
		links TEXT,
		country TEXT
	);`)
	checkError(err)
}

func createSecondaryTables(db *sql.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS qualifications(
		id INTEGER PRIMARY KEY,
		qualifications TEXT NOT NULL,
		FOREIGN KEY (id) REFERENCES job_data(id)
	);`)
	checkError(err)
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS links(
		id INTEGER PRIMARY KEY,
		links TEXT NOT NULL,
		FOREIGN KEY (id) REFERENCES job_data(id)
	);`)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
