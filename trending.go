package trending

import (
        "time"
)

const (
        defaultGranularity      = time.Hour * 24
        defaultSamples          = 14
)

type trending struct {
        ts              TimeSeries

        granularity     time.Duration
        samples         int
}

type Option func(*trending)

func WithGranularity(granularity time.Duration) Option {
        return func(t *trending) {
                t.granularity = granularity
        }
}

func WithSamples(samples int) Option {
        return func(t *trending) {
                t.samples = samples
        }
}

type TrendingServce interface {
        Add(topic string, tm time.Time)           error
        Score(topic string, tm time.Time)         (float64, error)
}

func NewTrending(ts TimeSeries, options ...Option) TrendingServce {
        tr := &trending{
                ts:             ts,
                granularity:    defaultGranularity,
                samples:        defaultSamples,
        }

        for _, option := range options {
                option(tr)
        }

        return tr
}

func (t *trending) Add (topic string, tm time.Time) error {
        if err := t.ts.Insert(topic, tm); err != nil {
                return err
        }

        return nil
}

func (t *trending) Score (topic string, tm time.Time) (float64, error) {
        count, _ := t.ts.Range(topic, tm.Add(-1 * t.granularity), tm)

        if count == 0 {
                return 0.0, nil
        }

        if avg, std, err := t.ts.Stats(topic, tm, t.granularity, t.samples); err != nil {
                return (float64(count) - avg) / std, nil
        }

        return 0.0, nil
}
