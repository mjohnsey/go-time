package main

import (
	"fmt"
	"log"
	"time"

	bitbar "github.com/johnmccabe/go-bitbar"
)

// LocationTime represents a named timezone to follow
type LocationTime struct {
	Name     string
	Location *time.Location
}

// NewLocationTime is the constructor for a new LocationTime
func (lt LocationTime) NewLocationTime(name string, tzName string) *LocationTime {
	loc, err := time.LoadLocation(tzName)
	if err != nil {
		log.Fatalln(err)
	}
	return &LocationTime{Name: name, Location: loc}
}

// PrettyPrint gives a nice string to print for a location
func (lt *LocationTime) PrettyPrint(t *time.Time) string {
	timeStr24 := t.In(lt.Location).Format(lt.TimeFormat24())
	timeStr12 := t.In(lt.Location).Format("3:04 PM")
	pp := fmt.Sprintf("%s - %s - %s", lt.Name, timeStr24, timeStr12)
	return pp
}

// Now returns the current time in a given location
func (lt *LocationTime) Now() *time.Time {
	now := time.Now().In(lt.Location)
	return &now
}

// PrettyPrintNow returns the current time in given location in a formatted string
func (lt *LocationTime) PrettyPrintNow() string {
	return lt.PrettyPrint(lt.Now())
}

// TimeFormat24 is the canonical time format in 24H
func (lt LocationTime) TimeFormat24() string {
	return "15:04"
}

// TimeFormat12 is the canonical time format in 12H
func (LocationTime) TimeFormat12() string {
	return "3:04 PM"
}

func main() {

	locs := make([]*LocationTime, 0)

	locs = append(locs, LocationTime{}.NewLocationTime("NYC", "America/New_York"))
	locs = append(locs, LocationTime{}.NewLocationTime("Memphis", "America/Chicago"))
	locs = append(locs, LocationTime{}.NewLocationTime("Seattle", "America/Los_Angeles"))

	now := time.Now()

	app := bitbar.New()
	app.StatusLine(fmt.Sprintf("%s UTC", now.UTC().Format(LocationTime{}.TimeFormat24())))

	submenu := app.NewSubMenu()
	for _, t := range locs {
		submenu.Line(t.PrettyPrint(&now))
	}

	app.Render()
}
