package main

import (
	"time"

	hourly "github.com/dillendev/training-hourly"
)

func getUsers() []hourly.User {
	return []hourly.User{
		{
			Id:   13934,
			Name: "Tomàs Adalgard",
			Rate: 1500,
		},
		{
			Id:   13935,
			Name: "Aravinda Dima",
			Rate: 2000,
		},
		{
			Id:   13936,
			Name: "Naoise Sulejman",
			Rate: 2240,
		},
		{
			Id:   14024,
			Name: "Funda Hulderic",
			Rate: 1985,
		},
		{
			Id:   14025,
			Name: "Bohuslav Babar",
			Rate: 1550,
		},
		{
			Id:   14026,
			Name: "Kisecawchuck Kwame",
			Rate: 2100,
		},
	}
}

var workDays = map[int][]time.Weekday{
	13934: {time.Monday, time.Tuesday, time.Thursday, time.Friday},
	13935: {time.Tuesday, time.Wednesday, time.Friday, time.Saturday},
	13936: {time.Monday, time.Tuesday, time.Wednesday, time.Friday, time.Saturday},
	14024: {time.Tuesday, time.Wednesday, time.Thursday, time.Saturday},
	14025: {time.Wednesday, time.Friday},
	14026: {time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday},
}

var projects = map[int]hourly.Project{
	13934: {
		Id:   103,
		Name: "ING Main Building",
	},
	14024: {
		Id:   104,
		Name: "Pathe Amersfoort",
	},
}

func isWorkingAt(date time.Time, days []time.Weekday) bool {
	for _, day := range days {
		if date.Weekday() == day {
			return true
		}
	}

	return false
}

// getTimeEntries generates time entries for a given user, the output is deterministic
func getTimeEntries(userId int) (entries []hourly.TimeEntry) {
	var (
		id      int
		project *hourly.Project
	)

	for startId, p := range projects {
		if userId >= startId {
			project = &p
			break
		}
	}

	days := workDays[userId]
	date := time.Date(2023, 2, 1, 0, 0, 0, 0, time.UTC)

	if userId%2 == 0 {
		id += 1

		entries = append(entries, hourly.TimeEntry{
			Id:        id,
			At:        date.Format(time.RFC3339),
			Billable:  false,
			StartedAt: date,
			StoppedAt: nil,
			Project:   *project,
		})
	}

	date = date.AddDate(0, 0, 1)

	for date.Month() == 2 {
		if isWorkingAt(date, days) {
			startedAt := date.Add(time.Hour * 8)

			if userId%2 == 0 {
				startedAt = startedAt.Add(time.Hour)
			}

			stoppedAt := startedAt.Add(time.Hour * 7)

			if userId%2 == 0 {
				stoppedAt = stoppedAt.Add(time.Hour)
			}

			if userId%3 == 0 {
				stoppedAt = stoppedAt.Add(time.Hour)
			}

			if int(date.Weekday())%2 == 0 && userId < 14000 {
				stoppedAt = stoppedAt.Add(time.Hour + (time.Minute * 30))
			}

			id += 1

			entries = append(entries, hourly.TimeEntry{
				Id:        id,
				At:        date.Format(time.RFC3339),
				Billable:  true,
				StartedAt: startedAt,
				StoppedAt: &stoppedAt,
				Project:   *project,
			})
		}

		date = date.AddDate(0, 0, 1)
	}

	return
}
