package data_helpers

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

// Time-related functions

// Current Date and Time
func TestDate() string {
	return time.Now().Format("2006-01-02")
}

func TestTime() string {
	return time.Now().Format("15:04:05")
}

func TestTimestamp() string {
	return time.Now().Format(time.RFC3339)
}

// Future Time functions
func FutureTimeMinutes(minutes int) string {
	return time.Now().Add(time.Duration(minutes) * time.Minute).Format("15:04:05")
}

func FutureTimeHours(hours int) string {
	return time.Now().Add(time.Duration(hours) * time.Hour).Format("15:04:05")
}

func FutureDateDays(days int) string {
	return time.Now().AddDate(0, 0, days).Format("2006-01-02")
}

// Past Time functions
func PastTimeMinutes(minutes int) string {
	return time.Now().Add(-time.Duration(minutes) * time.Minute).Format("15:04:05")
}

func PastTimeHours(hours int) string {
	return time.Now().Add(-time.Duration(hours) * time.Hour).Format("15:04:05")
}

func PastDateDays(days int) string {
	return time.Now().AddDate(0, 0, -days).Format("2006-01-02")
}

// Random value generators

// Generate a random UUID
func RandomUUID() string {
	return uuid.New().String()
}

// Generate a random 4-digit number
func Random4DigitNumber() int {
	n, err := rand.Int(rand.Reader, big.NewInt(9000))
	if err != nil {
		panic(err)
	}
	return int(n.Int64()) + 1000 // Ensure the number is 4 digits
}

// Test execution variables based on dynamic calculations

// Generate dynamic test values based on test round and current date
func GenerateTestVariables(todayTestRound int) map[string]int {
	testVariables := make(map[string]int)
	testVariables["todayTestRound"] = todayTestRound

	// Single digit test variable used in dynamic test data
	testVariables["testN"] = todayTestRound

	// Base two-digit number used for IDs in dynamic test data
	testNN := 20
	testVariables["testNN"] = testNN

	// Calculate testNN1 to testNN9
	for i := 1; i <= 9; i++ {
		key := fmt.Sprintf("testNN%d", i)
		testVariables[key] = testNN + i
	}

	// Dynamic 4-digit test value derived from current month, date, and test round
	currentMonth := int(time.Now().Month())
	currentDay := time.Now().Day()
	testMD := currentMonth*31 + currentDay + todayTestRound*1000
	testVariables["testMD"] = testMD
	testVariables["testMD500"] = testMD + 500

	// 6-digit number used for IDs in dynamic test data usually
	testDay, _ := strconv.Atoi(time.Now().Format("060102")) // yyMMdd
	testVariables["testDay"] = testDay

	// 7-digit number used for IDs in dynamic test data
	testDayN, _ := strconv.Atoi(fmt.Sprintf("%d%d", testDay, testVariables["testN"]))
	testVariables["testDayN"] = testDayN

	// 8-digit number used for IDs in dynamic test data
	testDayNN, _ := strconv.Atoi(fmt.Sprintf("%d%d", testDay, testVariables["testNN"]))
	testVariables["testDayNN"] = testDayNN

	return testVariables
}

// Helper function to get time with a specific time zone offset (e.g., "+01:00")
func GetTimeWithZoneOffset(offsetHours int) string {
	return time.Now().UTC().Add(time.Duration(offsetHours) * time.Hour).Format("2006-01-02T15:04:05-07:00")
}

// Sample usage of future date and time
func FutureDateTimeSample() string {
	futureDate := FutureDateDays(1)
	futureTime := FutureTimeHours(1)
	return fmt.Sprintf("%sT%s", futureDate, futureTime)
}

// GenerateTestCode generates a 4-digit test code based on a string in the format "110-010-001"
// tcId = ((categoryNumber / 10 - 1) * 500) + ((methodNumber / 10 - 1) * 50) + testcaseNumber
func GenerateTestCode(testCodeString string) error {
	// Split the input string by the '-' delimiter
	parts := strings.Split(testCodeString, "-")
	if len(parts) != 3 {
		return fmt.Errorf("invalid test code format: %s", testCodeString)
	}

	// Convert each part to an integer
	categoryNumber, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("invalid category number: %s", parts[0])
	}

	methodNumber, err := strconv.Atoi(parts[1])
	if err != nil {
		return fmt.Errorf("invalid method number: %s", parts[1])
	}

	testcaseNumber, err := strconv.Atoi(parts[2])
	if err != nil {
		return fmt.Errorf("invalid testcase number: %s", parts[2])
	}

	// Calculate the test code as before
	tcId := ((categoryNumber/10 - 1) * 500) + ((methodNumber/10 - 1) * 50) + testcaseNumber
	TestCode = fmt.Sprintf("%04d", tcId) // Ensures the TestCode is always a 4-digit number, and updates the global TestCode variable

	return nil
}
