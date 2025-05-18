package data

import (
	"encoding/json"
	"time"
)

type Job struct {
	posting JobPosting
	details JobDetails
}

type JobPosting struct {
	website      string
	url          string
	location     string
	company      string
	position     string
	jobType      string
	workShift    string
	workSetting  string
	lastModified time.Time
}

type JobDetails struct {
	skills             json.Marshaler
	licenses           json.Marshaler
	certs              json.Marshaler
	education          json.Marshaler
	benefits           json.Marshaler
	fullJobDescription string
}
