package memory

import (
        "math"
        "time"
)

type TimeSeries struct {
        data    map[string][]time.Time
}

func NewTimeSeries() *TimeSeries {
        return &TimeSeries{
                make(map[string][]time.Time),
        }
}

func (t *TimeSeries) Range(tag string, start time.Time, end time.Time) (int, error) {
        count := 0
        points, ok := t.data[tag]

        if !ok {
                return count, nil
        }

        for _, point := range points {
                if point.After(start) && end.After(point) {
                        count++
                }
        }

        return count, nil
}

func (t *TimeSeries) Stats(tag string, end time.Time, duration time.Duration, number int) (float64, float64, error) {
        points, ok := t.data[tag]
        if !ok {
                return 0, 0, nil
        }

        start := end.Add(duration * time.Duration(number) * -1)

        count := 0
        buckets := make([]int, number)
        for _, point := range points {
                if point.After(start) && end.After(point) {
                        index := int(point.Sub(start).Seconds() / duration.Seconds())

                        buckets[index]++
                        count++
                }
        }

        avg := float64(count) / float64(number)
        std := t.calStd(buckets, avg)

        return avg, std, nil
}

func (t *TimeSeries) Insert(tag string, tm time.Time) error {
        t.data[tag] = append(t.data[tag], tm)

        return nil
}

func (t *TimeSeries) calStd(nums []int, avg float64) float64 {
        var sum float64 = 0
        for _, num := range nums {
                sum += (float64(num) - avg) * (float64(num) - avg)
        }
        sum = sum / float64(len(nums))

        return math.Sqrt(sum)
}
