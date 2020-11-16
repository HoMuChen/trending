package trending

import (
        "time"
)

type Stat struct {
        Avg     float64
        Std     float64
}

type TimeSeries interface {
        Range(key string, start time.Time, end time.Time)                               (int, error)
        Stats(key string, end time.Time, duration time.Duration, number int)            (Stat, error)
        Insert(key string, tm time.Time)                                                (error)
}
