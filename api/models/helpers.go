package models

import (
	"strconv"
	"strings"
	"time"
)

var mounths = map[string]time.Month{
	"Января":   time.January,
	"Февраля":  time.February,
	"Марта":    time.March,
	"Апреля":   time.April,
	"Мая":      time.May,
	"Июня":     time.June,
	"Июля":     time.July,
	"Августа":  time.August,
	"Сентября": time.September,
	"Октября":  time.October,
	"Ноября":   time.November,
	"Декабря":  time.December,
}

func ParseBankerTime(raw string) *time.Time {
	split := strings.Split(raw, " ")
	currentTime := time.Now()
	loc, _ := time.LoadLocation("Europe/Moscow")
	currentTime.In(loc)
	if split[0] != "" {
		day := 0
		month := time.January
		year := currentTime.Year()
		rawHoursMinutes := ""
		switch split[0] {
		case "сегодня":
			day = currentTime.Day()
			month = currentTime.Month()
			rawHoursMinutes = split[2]
		case "вчера":
			tomorrow := currentTime.Add(-24 * time.Hour)
			day = tomorrow.Day()
			month = tomorrow.Month()
			rawHoursMinutes = split[2]
		default:
			day, _ = strconv.Atoi(split[0])
			month = mounths[split[1]]
			if len(split) == 4 {
				rawHoursMinutes = split[3]
			} else {
				year, _ = strconv.Atoi(split[3])
				rawHoursMinutes = split[4]
			}
		}
		timeSplit := strings.Split(rawHoursMinutes, ":")
		hours, _ := strconv.Atoi(timeSplit[0])
		minutes, _ := strconv.Atoi(timeSplit[1])
		result := time.Date(year, month, day, hours, minutes, 0, 0, loc)
		utcLoc, _ := time.LoadLocation("UTC")
		result.In(utcLoc)
		return &result
	}
	return nil
}
