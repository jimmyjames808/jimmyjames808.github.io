package main

import "time"

type Shift struct {
	StartTime time.Time  `json:"start_time"`
	EndTime   *time.Time `json:"end_time"`
}
