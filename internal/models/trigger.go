package models

import "time"

type Trigger struct {
	ID             uint                 `json:"id"`
	Type           string               `json:"type"`
	Schedule       ScheduledTriggerType `json:"schedule,omitempty"` // could be anything from ScheduledTriggerType i.e. OneTimeTrigger, FixedIntervalTrigger, DailyFixedTimeTrigger
	FixedInterval  string               `json:"fixed_interval,omitempty"`
	DailyFixedTime string               `json:"daily_fixed_time,omitempty"`
	CreatedAt      time.Time            `json:"created_at"`
	TriggeredAt    time.Time            `json:"triggered_at,omitempty"`
}
