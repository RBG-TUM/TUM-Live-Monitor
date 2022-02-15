package monitor

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
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
	var res *[]liveStreamDto
	hr, err := http.Get(m.instance + "/api/stream/live?token=" + m.token)
	if err != nil {
		return err
	}
	err = json.NewDecoder(hr.Body).Decode(&res)
	if err != nil {
		return err
	}
	if res == nil {
		return nil
	}
	m.streams = make([]*Stream, len(*res))
	for _, s := range *res {
		m.streams = append(m.streams, &Stream{
			StreamID:    s.ID,
			Title:       s.CourseName,
			Until:       s.End,
			LectureHall: s.LectureHall,
			LastUpdate:  time.Now(),
			Cam:         s.CAM,
			Pres:        s.PRES,
			Comb:        s.COMB,
		})
	}
	return nil
}

type liveStreamDto struct {
	ID          uint
	CourseName  string
	LectureHall string
	COMB        string
	PRES        string
	CAM         string
	End         time.Time
}
