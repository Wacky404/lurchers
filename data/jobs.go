package data

import (
	"encoding/json"
	"time"
)

type Job struct {
	Posting *JobPosting
	Details *JobDetails
}

func NewJob() *Job {
	return &Job{Posting: &JobPosting{}, Details: &JobDetails{}}
}

type JobPosting struct {
	Website      string
	Url          string
	Location     string
	Company      string
	Position     string
	JobType      string
	WorkShift    string
	WorkSetting  string
	LastModified time.Time
}

type JobDetails struct {
	Skills             json.Marshaler
	Licenses           json.Marshaler
	Certs              json.Marshaler
	Education          json.Marshaler
	Benefits           json.Marshaler
	FullJobDescription string
}
