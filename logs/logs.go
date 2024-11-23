package logs

import (
	"math/rand"
	"time"
)

var actions = []string{"logged in", "logged out", "created record", "deleted record", "updated account"}

type Log struct {
	Action    string
	Timestamp time.Time
}

func GenerateLogs(count int) []Log {
	logs := make([]Log, count)

	for i := 0; i < count; i++ {
		logs[i] = Log{
			Action:    actions[rand.Intn(len(actions))],
			Timestamp: time.Now(),
		}
	}
	return logs
}
