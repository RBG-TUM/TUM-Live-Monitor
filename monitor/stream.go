package monitor

import "time"

type Stream struct {
	StreamID    uint      `json:"stream_id"`    // e.g. 1
	Title       string    `json:"title"`        // e.g. "Einführung in die Informatik 1"
	URL         string    `json:"url"`          // e.g. "https://tum.de/live/stream/1.m3u8"
	Until       time.Time `json:"until"`        // when stream is over -> stop monitoring
	LectureHall string    `json:"lecture_hall"` // e.g. "MI HS 1"
	LastUpdate  time.Time `json:"last_update"`
}
