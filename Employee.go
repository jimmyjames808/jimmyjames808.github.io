package main

import (
	"time"
)

type Employee struct {
	Name        string     `json:"name"`
	Id          int        `json:"id"`
	Shifts      []Shift    `json:"shifts"`
	PayPerHour  float64    `json:"payPerHour"`
	CurrentPay  float64    `json:"currentPay"`
	OwnerStruct *Employees `json:"-"`
}

func (e *Employee) StartShift() {
	if len(e.Shifts) > 0 {
		if e.Shifts[len(e.Shifts)-1].EndTime == nil {
			e.EndShift()
		}
	}
	e.Shifts = append(e.Shifts, Shift{StartTime: time.Now()})
	e.OwnerStruct.saveToJson()
}

func (e *Employee) EndShift() {
	if len(e.Shifts) > 0 {
		if e.Shifts[len(e.Shifts)-1].EndTime == nil {
			time2 := time.Now()
			e.Shifts[len(e.Shifts)-1].EndTime = &time2
			e.CurrentPay += e.PayPerHour * e.Shifts[len(e.Shifts)-1].EndTime.Sub(e.Shifts[len(e.Shifts)-1].StartTime).Hours()
		}
	}
	e.OwnerStruct.saveToJson()
}
