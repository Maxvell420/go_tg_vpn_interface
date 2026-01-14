package entities

type Job struct {
	UserId    int
	CreatedAt int
	Type      JobType
	Data      *map[string]string
}

type JobType int

const (
	TrafficUsage JobType = 1
)
