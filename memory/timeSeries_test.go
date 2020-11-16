package memory_test

import (
        "fmt"
        "testing"
        "time"

        "github.com/HoMuChen/trending/memory"
)

func TestCount(t *testing.T) {
        ts := memory.NewTimeSeries()

        ts.Insert("a", time.Unix(1605000000, 0))
        ts.Insert("a", time.Unix(1605000000, 0))
        ts.Insert("a", time.Unix(1606000000, 0))
        ts.Insert("a", time.Unix(1604000000, 0))
        ts.Insert("b", time.Unix(1604000000, 0))

        if count_a, _ := ts.Range("a", time.Unix(1600000000, 0), time.Unix(1700000000, 0)); count_a != 4 {
                t.Errorf("count for a should be 1 but got: %v", count_a)
        }
        if count_a, _ := ts.Range("a", time.Unix(1604500000, 0), time.Unix(1607000000, 0)); count_a != 3 {
                t.Errorf("count for a should be 0 but got: %v", count_a)
        }
        if count_b, _ := ts.Range("b", time.Unix(1603000000, 0), time.Unix(1607000000, 0)); count_b != 1 {
                t.Errorf("count for b should be 0 but got: %v", count_b)
        }
        if count_b, _ := ts.Range("b", time.Unix(1606000000, 0), time.Unix(1607000000, 0)); count_b != 0 {
                t.Errorf("count for b should be 0 but got: %v", count_b)
        }
}


func TestStats(t *testing.T) {
        ts := memory.NewTimeSeries()

        ts.Insert("a", time.Unix(1605000000, 0).Add(-6 * time.Hour * 24))
        ts.Insert("a", time.Unix(1605000000, 0).Add(-5 * time.Hour * 24))
        ts.Insert("a", time.Unix(1605000000, 0).Add(-4 * time.Hour * 24))
        ts.Insert("a", time.Unix(1605000000, 0).Add(-3 * time.Hour * 24))
        ts.Insert("a", time.Unix(1605000000, 0).Add(-2 * time.Hour * 24))
        ts.Insert("a", time.Unix(1605000000, 0).Add(-1 * time.Hour * 24))
        ts.Insert("a", time.Unix(1605000000, 0))
        ts.Insert("b", time.Unix(1605000000, 0))
        ts.Insert("b", time.Unix(1605000000, 0))
        ts.Insert("b", time.Unix(1605000000, 0))

        stat, _ := ts.Stats("a", time.Unix(1605000000, 1), time.Hour * 24, 7)

        if stat.Avg != 1 {
                t.Errorf("Avg should be 0.1 but got: %v", stat.Avg)
        }

        if stat.Std != 0 {
                t.Errorf("Avg should be 0 but got: %v", stat.Std)
        }
}
