package trending

import (
        "time"
)

type TimeSeries interface {
        Range(key string, start time.Time, end time.Time)                               (count int, err error)
        Stats(key string, end time.Time, duration time.Duration, number int)            (avg, std float64, err error)
        Insert(key string, tm time.Time)                                                (err error)
}
