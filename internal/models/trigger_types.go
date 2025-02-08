package models

type ScheduledTriggerType string
type ApiTriggerType string

const (
	OneTimeTrigger        ScheduledTriggerType = "one_time_trigger"
	FixedIntervalTrigger  ScheduledTriggerType = "fixed_interval_trigger"
	DailyFixedTimeTrigger ScheduledTriggerType = "daily_fixed_time_trigger"
	ApiTrigger            ApiTriggerType       = "api_trigger"
)
