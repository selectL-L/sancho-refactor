package utils

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// TimeUtils contains utility functions for time operations
type TimeUtils struct {
	DefaultTimezone string
	TimezoneMap     map[string]string // Maps user IDs to timezone strings
}

// NewTimeUtils creates a new TimeUtils instance
func NewTimeUtils(defaultTimezone string) *TimeUtils {
	return &TimeUtils{
		DefaultTimezone: defaultTimezone,
		TimezoneMap:     make(map[string]string),
	}
}

// LoadTimezones loads user timezones from file
func (t *TimeUtils) LoadTimezones(filename string) error {
	bytes, err := ReadFile(filename)
	if err != nil {
		return err
	}

	lines := strings.Split(string(bytes), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		parts := strings.SplitN(line, " ", 2)
		if len(parts) != 2 {
			continue
		}

		userID := parts[0]
		timezone := parts[1]

		t.TimezoneMap[userID] = timezone
	}

	return nil
}

// SaveTimezones saves user timezones to file
func (t *TimeUtils) SaveTimezones(filename string) error {
	var builder strings.Builder

	for userID, timezone := range t.TimezoneMap {
		builder.WriteString(userID)
		builder.WriteString(" ")
		builder.WriteString(timezone)
		builder.WriteString("\n")
	}

	return WriteFile(filename, []byte(builder.String()))
}

// SetUserTimezone sets a user's timezone
func (t *TimeUtils) SetUserTimezone(userID, timezone string) error {
	// Validate timezone
	_, err := time.LoadLocation(timezone)
	if err != nil {
		return fmt.Errorf("invalid timezone: %w", err)
	}

	t.TimezoneMap[userID] = timezone
	return nil
}

// GetUserTimezone gets a user's timezone
func (t *TimeUtils) GetUserTimezone(userID string) string {
	if tz, ok := t.TimezoneMap[userID]; ok {
		return tz
	}
	return t.DefaultTimezone
}

// ParseDateTime parses a date/time string in various formats
func (t *TimeUtils) ParseDateTime(input string, userID string) (time.Time, error) {
	// If it's a unix timestamp
	if timestamp, err := strconv.ParseInt(input, 10, 64); err == nil {
		return time.Unix(timestamp, 0), nil
	}

	// Get user's timezone
	tzString := t.GetUserTimezone(userID)
	location, err := time.LoadLocation(tzString)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid timezone: %w", err)
	}

	// Try various date formats
	formats := []string{
		time.RFC3339,
		"2006-01-02T15:04:05",
		"2006-01-02 15:04:05",
		"2006-01-02 15:04",
		"2006-01-02",
		"01/02/2006 15:04:05",
		"01/02/2006 15:04",
		"01/02/2006",
		"01.02.2006 15:04:05",
		"01.02.2006 15:04",
		"01.02.2006",
		"15:04:05",
		"15:04",
	}

	for _, format := range formats {
		if t, err := time.ParseInLocation(format, input, location); err == nil {
			return t, nil
		}
	}

	return time.Time{}, errors.New("could not parse date/time")
}

// ParseRelativeTime parses a relative time string like "2h30m"
func (t *TimeUtils) ParseRelativeTime(input string) (time.Duration, error) {
	// Direct duration parsing
	if duration, err := time.ParseDuration(input); err == nil {
		return duration, nil
	}

	// More complex time string parsing
	totalDuration := time.Duration(0)

	// Parse time units like "2h", "30m", "1d", etc.
	re := regexp.MustCompile(`(\d+)([smhdwy])`)
	matches := re.FindAllStringSubmatch(input, -1)

	if len(matches) == 0 {
		return 0, errors.New("invalid time format")
	}

	for _, match := range matches {
		if len(match) != 3 {
			continue
		}

		value, err := strconv.Atoi(match[1])
		if err != nil {
			continue
		}

		unit := match[2]

		switch unit {
		case "s":
			totalDuration += time.Duration(value) * time.Second
		case "m":
			totalDuration += time.Duration(value) * time.Minute
		case "h":
			totalDuration += time.Duration(value) * time.Hour
		case "d":
			totalDuration += time.Duration(value) * 24 * time.Hour
		case "w":
			totalDuration += time.Duration(value) * 7 * 24 * time.Hour
		case "y":
			totalDuration += time.Duration(value) * 365 * 24 * time.Hour
		}
	}

	if totalDuration == 0 {
		return 0, errors.New("could not parse time duration")
	}

	return totalDuration, nil
}

// ParseFuzzyTime parses human-readable time descriptions
func (t *TimeUtils) ParseFuzzyTime(input string, userID string) (time.Time, error) {
	input = strings.ToLower(strings.TrimSpace(input))

	// Get user's timezone
	tzString := t.GetUserTimezone(userID)
	location, err := time.LoadLocation(tzString)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid timezone: %w", err)
	}

	now := time.Now().In(location)

	// Handle special cases
	switch input {
	case "now", "right now", "immediately":
		return now, nil
	case "tomorrow":
		return now.AddDate(0, 0, 1), nil
	case "next week":
		return now.AddDate(0, 0, 7), nil
	case "next month":
		return now.AddDate(0, 1, 0), nil
	case "next year":
		return now.AddDate(1, 0, 0), nil
	}

	// Try to parse relative expressions like "in 2 hours"
	if strings.HasPrefix(input, "in ") {
		durationStr := strings.TrimPrefix(input, "in ")
		duration, err := t.ParseComplexDuration(durationStr)
		if err == nil {
			return now.Add(duration), nil
		}
	}

	// Try to parse time expressions like "at 3pm"
	if strings.HasPrefix(input, "at ") {
		timeStr := strings.TrimPrefix(input, "at ")
		return t.ParseTimeOfDay(timeStr, now, location)
	}

	// Try to parse date expressions like "on 01.02.2023"
	if strings.HasPrefix(input, "on ") {
		dateStr := strings.TrimPrefix(input, "on ")
		return t.ParseDate(dateStr, now, location)
	}

	return time.Time{}, errors.New("could not parse time description")
}

// ParseComplexDuration parses complex duration strings like "2 hours 30 minutes"
func (t *TimeUtils) ParseComplexDuration(input string) (time.Duration, error) {
	input = strings.ToLower(strings.TrimSpace(input))
	totalDuration := time.Duration(0)

	// Define patterns for time units
	patterns := []struct {
		re   *regexp.Regexp
		unit time.Duration
	}{
		{regexp.MustCompile(`(\d+)\s*(s|sec|second|seconds)`), time.Second},
		{regexp.MustCompile(`(\d+)\s*(m|min|minute|minutes)`), time.Minute},
		{regexp.MustCompile(`(\d+)\s*(h|hr|hour|hours)`), time.Hour},
		{regexp.MustCompile(`(\d+)\s*(d|day|days)`), 24 * time.Hour},
		{regexp.MustCompile(`(\d+)\s*(w|wk|week|weeks)`), 7 * 24 * time.Hour},
		{regexp.MustCompile(`(\d+)\s*(mo|month|months)`), 30 * 24 * time.Hour},
		{regexp.MustCompile(`(\d+)\s*(y|yr|year|years)`), 365 * 24 * time.Hour},
	}

	// Extract durations
	for _, pattern := range patterns {
		matches := pattern.re.FindAllStringSubmatch(input, -1)
		for _, match := range matches {
			if len(match) < 2 {
				continue
			}

			value, err := strconv.Atoi(match[1])
			if err != nil {
				continue
			}

			totalDuration += time.Duration(value) * pattern.unit
		}
	}

	if totalDuration == 0 {
		return 0, errors.New("could not parse duration")
	}

	return totalDuration, nil
}

// ParseTimeOfDay parses time expressions like "3pm" or "15:30"
func (t *TimeUtils) ParseTimeOfDay(timeStr string, baseTime time.Time, location *time.Location) (time.Time, error) {
	var hour, minute, second int

	// Handle 12-hour format like "3pm", "3:30pm"
	re12Hour := regexp.MustCompile(`(\d{1,2})(?::(\d{2}))?(?::(\d{2}))?\s*(am|pm)`)
	if matches := re12Hour.FindStringSubmatch(timeStr); matches != nil {
		hour, _ = strconv.Atoi(matches[1])
		if matches[2] != "" {
			minute, _ = strconv.Atoi(matches[2])
		}
		if matches[3] != "" {
			second, _ = strconv.Atoi(matches[3])
		}

		if matches[4] == "pm" && hour < 12 {
			hour += 12
		} else if matches[4] == "am" && hour == 12 {
			hour = 0
		}
	} else {
		// Handle 24-hour format like "15:30", "15"
		re24Hour := regexp.MustCompile(`(\d{1,2})(?::(\d{2}))?(?::(\d{2}))?`)
		if matches := re24Hour.FindStringSubmatch(timeStr); matches != nil {
			hour, _ = strconv.Atoi(matches[1])
			if matches[2] != "" {
				minute, _ = strconv.Atoi(matches[2])
			}
			if matches[3] != "" {
				second, _ = strconv.Atoi(matches[3])
			}
		} else {
			return time.Time{}, errors.New("could not parse time")
		}
	}

	// Create a new time with the specified hour, minute, and second
	result := time.Date(
		baseTime.Year(),
		baseTime.Month(),
		baseTime.Day(),
		hour,
		minute,
		second,
		0,
		location,
	)

	// If the resulting time is in the past, add a day
	if result.Before(baseTime) {
		result = result.AddDate(0, 0, 1)
	}

	return result, nil
}

// ParseDate parses date expressions like "01.02.2023"
func (t *TimeUtils) ParseDate(dateStr string, baseTime time.Time, location *time.Location) (time.Time, error) {
	// Try various date formats
	formats := []string{
		"2006-01-02",
		"01/02/2006",
		"01.02.2006",
		"02/01/2006",
		"02.01.2006",
	}

	for _, format := range formats {
		if date, err := time.ParseInLocation(format, dateStr, location); err == nil {
			// Keep the same time of day as the base time
			return time.Date(
				date.Year(),
				date.Month(),
				date.Day(),
				baseTime.Hour(),
				baseTime.Minute(),
				baseTime.Second(),
				0,
				location,
			), nil
		}
	}

	// Handle special formats like "today", "tomorrow", etc.
	switch dateStr {
	case "today":
		return baseTime, nil
	case "tomorrow":
		return baseTime.AddDate(0, 0, 1), nil
	case "yesterday":
		return baseTime.AddDate(0, 0, -1), nil
	}

	// Handle day of week
	daysOfWeek := map[string]int{
		"sunday":    0,
		"monday":    1,
		"tuesday":   2,
		"wednesday": 3,
		"thursday":  4,
		"friday":    5,
		"saturday":  6,
	}

	if dayNum, ok := daysOfWeek[dateStr]; ok {
		daysToAdd := (dayNum - int(baseTime.Weekday()) + 7) % 7
		if daysToAdd == 0 {
			daysToAdd = 7
		}
		return baseTime.AddDate(0, 0, daysToAdd), nil
	}

	// Handle "next <day>"
	for day, dayNum := range daysOfWeek {
		if dateStr == "next "+day {
			daysToAdd := (dayNum - int(baseTime.Weekday()) + 7) % 7
			if daysToAdd == 0 {
				daysToAdd = 7
			}
			return baseTime.AddDate(0, 0, daysToAdd), nil
		}
	}

	return time.Time{}, errors.New("could not parse date")
}

// FormatDuration formats a time.Duration in a human-readable format
func (t *TimeUtils) FormatDuration(d time.Duration) string {
	d = d.Round(time.Second)

	if d < time.Minute {
		return fmt.Sprintf("%d seconds", d/time.Second)
	}

	if d < time.Hour {
		m := d / time.Minute
		d -= m * time.Minute
		s := d / time.Second
		if s == 0 {
			return fmt.Sprintf("%d minutes", m)
		}
		return fmt.Sprintf("%d minutes %d seconds", m, s)
	}

	if d < 24*time.Hour {
		h := d / time.Hour
		d -= h * time.Hour
		m := d / time.Minute
		if m == 0 {
			return fmt.Sprintf("%d hours", h)
		}
		return fmt.Sprintf("%d hours %d minutes", h, m)
	}

	days := d / (24 * time.Hour)
	d -= days * 24 * time.Hour
	h := d / time.Hour
	if h == 0 {
		return fmt.Sprintf("%d days", days)
	}
	return fmt.Sprintf("%d days %d hours", days, h)
}

// FormatTimestamp formats a timestamp for display in Discord
func (t *TimeUtils) FormatTimestamp(timestamp int64, format string) string {
	switch format {
	case "R": // Relative
		return fmt.Sprintf("<t:%d:R>", timestamp)
	case "F": // Full date and time
		return fmt.Sprintf("<t:%d:F>", timestamp)
	case "T": // Time only
		return fmt.Sprintf("<t:%d:T>", timestamp)
	case "D": // Date only
		return fmt.Sprintf("<t:%d:D>", timestamp)
	case "d": // Short date
		return fmt.Sprintf("<t:%d:d>", timestamp)
	case "t": // Short time
		return fmt.Sprintf("<t:%d:t>", timestamp)
	default: // Default timestamp
		return fmt.Sprintf("<t:%d>", timestamp)
	}
}

// Now returns the current time in a user's timezone
func (t *TimeUtils) Now(userID string) (time.Time, error) {
	tzString := t.GetUserTimezone(userID)
	location, err := time.LoadLocation(tzString)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid timezone: %w", err)
	}

	return time.Now().In(location), nil
}

// GetTimezoneLocation returns the time.Location for a user
func (t *TimeUtils) GetTimezoneLocation(userID string) (*time.Location, error) {
	tzString := t.GetUserTimezone(userID)
	return time.LoadLocation(tzString)
}

// ValidateTimezone checks if a timezone string is valid
func (t *TimeUtils) ValidateTimezone(timezone string) error {
	_, err := time.LoadLocation(timezone)
	if err != nil {
		return fmt.Errorf("invalid timezone: %w", err)
	}
	return nil
}

// ConvertTime converts a time to another timezone
func (t *TimeUtils) ConvertTime(tm time.Time, targetTimezone string) (time.Time, error) {
	loc, err := time.LoadLocation(targetTimezone)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid timezone: %w", err)
	}

	return tm.In(loc), nil
}
