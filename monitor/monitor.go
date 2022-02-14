package monitor

import (
	"fmt"
	log "github.com/sirupsen/logrus"
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
	i := viper.GetString("baseurl")
	t := viper.GetString("token")
	return &Monitor{instance: i, token: t}
}

// Run starts the monitoring process
func (m *Monitor) Run() {
	log.Info("starting monitor")
	for {
		go func() {
			err := m.fetchStreams()
			if err != nil {
				log.Error(fmt.Errorf("fetch stream: %v", err))
			}
		}()
		time.Sleep(time.Second * 10)
		// todo
	}
}

func (m *Monitor) fetchStreams() error {
	m.streams = []*Stream{{
		StreamID:    0,
		Title:       "Eidi",
		URL:         "https://live.stream.lrz.de/livetum/smil:70-dda14e25_all.smil/playlist.m3u8",
		Until:       time.Now(),
		LectureHall: "HS 2",
		LastUpdate:  time.Now(),
	}}
	return nil
}
