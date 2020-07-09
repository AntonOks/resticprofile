package schedule

import (
	"fmt"
	"path"
	"path/filepath"

	"github.com/creativeprojects/resticprofile/calendar"
	"github.com/creativeprojects/resticprofile/config"
)

// Job scheduler
type Job struct {
	configFile string
	profile    *config.Profile
	schedules  []*calendar.Event
}

// NewJob instantiates a Job object to schedule jobs
func NewJob(configFile string, profile *config.Profile) *Job {
	return &Job{
		configFile: configFile,
		profile:    profile,
	}
}

// Create a new job
func (j *Job) Create() error {
	err := checkSystem()
	if err != nil {
		return err
	}

	err = j.checkSchedules()
	if err != nil {
		return err
	}

	err = j.createJob()
	if err != nil {
		return err
	}

	return nil
}

// Update an existing job
func (j *Job) Update() error {
	err := checkSystem()
	if err != nil {
		return err
	}
	return nil
}

// Remove a job
func (j *Job) Remove() error {
	err := checkSystem()
	if err != nil {
		return err
	}
	err = j.removeJob()
	if err != nil {
		return err
	}
	return nil
}

// Status of a job
func (j *Job) Status() error {
	err := checkSystem()
	if err != nil {
		return err
	}

	return j.displayStatus()
}

func (j *Job) checkSchedules() error {
	var err error
	if j.profile.Schedule == nil || len(j.profile.Schedule) == 0 {
		return fmt.Errorf("no schedule found for profile '%s'", j.profile.Name)
	}
	j.schedules, err = loadSchedules(j.profile.Schedule)
	return err
}

// absolutePathToBinary returns an absolute path to the resticprofile binary
func absolutePathToBinary(currentDir, binaryPath string) string {
	binary := binaryPath
	if !filepath.IsAbs(binary) {
		binary = path.Join(currentDir, binary)
	}
	return binary
}