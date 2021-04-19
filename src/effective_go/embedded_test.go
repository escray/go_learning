package effective_go

import (
	"fmt"
	"log"
	"os"
	"testing"
)

type Job struct {
	Command string
	*log.Logger
}

func TestEmbedded(t *testing.T) {
	job := Job{}
	job.Println("starting now...")
	command := ""
	jobb := &Job{command, log.New(os.Stderr, "Job: ", log.Ldate)}
	_ = jobb
}

func NewJob(command string, logger *log.Logger) *Job {
	return &Job{command, logger}
}

func (job *Job) Printf(format string, args ...interface{})  {
	job.Logger.Printf("%q: %s", job.Command, fmt.Sprintf(format, args...))
}
