package monitor

import (
	"fmt"
	"sync"
	"time"

	"github.com/spf13/viper"
)

type Monitor struct {
	instance string // url of TUM-Live instance to monitor
	token    string // token to access TUM-Live instance

	lock    sync.Mutex
	streams []*Stream
}

// AddStream adds a stream to monitor
func (m *Monitor) AddStream(s *Stream) {
	fmt.Println("[Info] adding stream")
	m.lock.Lock()
	defer m.lock.Unlock()

	m.streams = append(m.streams, s)
}

// New creates a new Monitor instance
func New() *Monitor {
	i := viper.GetString("instance")
	t := viper.GetString("token")
	return &Monitor{instance: i, token: t}
}

// Run starts the monitoring process
func (m *Monitor) Run() {
	fmt.Println("[Info] starting monitor")
	for {
		time.Sleep(time.Second * 10)
		// todo
	}
}
