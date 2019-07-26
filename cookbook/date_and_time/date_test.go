package date_and_time

import (
	"testing"
	"time"
)

func TestLifeDays(t *testing.T) {
	birthday := time.Date(1991, time.January, 21, 0, 0, 0, 0, time.Now().Location())
	t.Log(LifeDays(birthday))
}

func TestYearsOld(t *testing.T) {
	birthday := time.Date(1991, time.January, 21, 0, 0, 0, 0, time.Now().Location())
	t.Log(YearsOld(birthday))
}
