package date_and_time

import "time"

//计算有生之日
func LifeDays(birthday time.Time) int {
	now := time.Now()
	now = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	lifeDays := now.Sub(birthday)
	return int(lifeDays.Hours() / 24)
}

//计算年岁
func YearsOld(birthday time.Time) int {
	now := time.Now()
	now = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	return now.Year() - birthday.Year()
}
