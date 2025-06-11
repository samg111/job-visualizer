package database

import (
	"database/sql"
	"job-visualizer/pkg/structs"
	"log"
)

func OpenOrCreateDatabase() *sql.DB {
	db, err := sql.Open("sqlite", "job_data.sqlite")
	checkError(err)
	// defer db.Close()
	return db
}

func SetupDatabase(db *sql.DB) {
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
		links TEXT
	);`)
	checkError(err)
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS qualifications(
		qualifications TEXT NOT NULL,
		FOREIGN KEY (qualifications) REFERENCES job_data(qualifications)
	);`)
	checkError(err)
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS links(
		links TEXT NOT NULL,
		FOREIGN KEY (links) REFERENCES job_data(links)
	);`)
	checkError(err)
}

func WriteToDatabase(db *sql.DB, allJobData []structs.JobData) {
	insertQueryJobData := `INSERT INTO job_data (location, job_title, company_name, description, date_posted, salary,
		work_from_home, qualifications, links) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);`
	insertQueryQualifications := `INSERT OR IGNORE INTO qualifications (qualifications) VALUES (?);`
	insertQueryLinks := `INSERT OR IGNORE INTO links (links) VALUES (?);`
	for _, job := range allJobData {
		_, err := db.Exec(insertQueryJobData, job.Location, job.JobTitle, job.CompanyName, job.Description, job.DatePosted,
			job.Salary, job.WorkFromHome, job.Qualifications, job.Links)
		checkError(err)
		_, err = db.Exec(insertQueryQualifications, job.Qualifications)
		checkError(err)
		_, err = db.Exec(insertQueryLinks, job.Links)
		checkError(err)
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
