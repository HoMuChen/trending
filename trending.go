package trending

import (
        "time"
)

type trending struct {
        ts              TimeSeries

        interval        time.Duration
        n               int
}

type TrendingServce interface {
        Add(topic string, tm time.Time)           error
        Score(topic string, tm time.Time)         (float64, error)
}

func NewTrending(ts TimeSeries, interval time.Duration, n int) TrendingServce {
        return &trending{
                ts,
                interval,
                n,
        }
}

func (t *trending) Add (topic string, tm time.Time) error {
        if err := t.ts.Insert(topic, tm); err != nil {
                return err
        }

        return nil
}

func (t *trending) Score (topic string, tm time.Time) (float64, error) {
        count, _ := t.ts.Range(topic, tm.Add(-1 * t.interval), tm)

        if count == 0 {
                return 0.0, nil
        }

        if stat, err := t.ts.Stats(topic, tm, t.interval, t.n); err != nil {
                return (float64(count) - stat.Avg) / stat.Std, nil
        }

        return 0.0, nil
}
