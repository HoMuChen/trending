package trending_test

import (
        "testing"
        "time"

        "github.com/HoMuChen/trending/memory"
        "github.com/HoMuChen/trending"
)


func TestTrending(t *testing.T) {
        ts := memory.NewTimeSeries()

        tr := trending.NewTrending(
                ts,
                trending.WithGranularity(time.Hour * 24),
                trending.WithSamples(14),
        )

        tr.Add("a", time.Now().Add(-10 * time.Hour * 24))
        tr.Add("a", time.Now().Add(-9 * time.Hour * 24))
        tr.Add("a", time.Now().Add(-8 * time.Hour * 24))
        tr.Add("a", time.Now().Add(-7 * time.Hour * 24))
        tr.Add("a", time.Now().Add(-6 * time.Hour * 24))
        tr.Add("a", time.Now().Add(-5 * time.Hour * 24))
        tr.Add("a", time.Now().Add(-4 * time.Hour * 24))
        tr.Add("a", time.Now().Add(-3 * time.Hour * 24))
        tr.Add("a", time.Now().Add(-2 * time.Hour * 24))
        tr.Add("a", time.Now().Add(-1 * time.Hour * 24))
        tr.Add("a", time.Now())

        score, _ := tr.Score("a", time.Now())

        if score != 0 {
                t.Errorf("score should be 0 but got: %v", score)
        }
}
