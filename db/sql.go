package db

/*
sql statements to be run against the specified Database
in connect.go
*/

type JobPosting struct {
	addRow string
}

func NewJobPosting() *JobPosting {
	return &JobPosting{addRow: addRowJobPosting}
}

type JobDetails struct {
	addRow string
}

func NewJobDetails() *JobDetails {
	return &JobDetails{addRow: addRowJobDetails}
}

type DbStatements struct {
	JobPosting *JobPosting
	JobDetails *JobDetails
}

func NewDbStatements() *DbStatements {
	return &DbStatements{JobPosting: NewJobPosting(), JobDetails: NewJobDetails()}
}

// table job_posting; add row
// return a job_id; to be used as FK
var addRowJobPosting string = `
	INSERT INTO job_posting (job_type, website, url, location, company, position, work_shift, work_setting, date_added, last_updated)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	RETURNING job_id
	`

// table job_details; add row
var addRowJobDetails string = `
	INSERT INTO job_details (job_id, skills, licences, certifications, education, benefits, full_job_description)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
