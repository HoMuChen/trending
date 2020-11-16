package memory

import (
        "math"
        "time"

        "github.com/HoMuChen/trending"
)

type ts struct {
        data    map[string][]time.Time
}

func NewTimeSeries() trending.TimeSeries {
        return &ts{
                make(map[string][]time.Time),
        }
}

func (t *ts) Range(tag string, start time.Time, end time.Time) (int, error) {
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

func (t *ts) Stats(tag string, end time.Time, duration time.Duration, number int) (trending.Stat, error) {
        points, ok := t.data[tag]
        if !ok {
                return trending.Stat{0, 0}, nil
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

        ret := trending.Stat{avg, std}

        return ret, nil
}

func (t *ts) Insert(tag string, tm time.Time) error {
        t.data[tag] = append(t.data[tag], tm)

        return nil
}

func (t *ts) calStd(nums []int, avg float64) float64 {
        var sum float64 = 0
        for _, num := range nums {
                sum += (float64(num) - avg) * (float64(num) - avg)
        }
        sum = sum / float64(len(nums))

        return math.Sqrt(sum)
}
