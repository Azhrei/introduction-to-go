package simplestrings

import "testing"

const weekdays = "Monday Tuesday Wednesday Thursday Friday"

// TestFindTuesday tests that Tuesday is a weekday
func TestFindTuesday(t *testing.T) {
    result := Contains(weekdays, "Tuesday")
    t.Log("Preconditions established")
    if ! result {
        t.Fatal("Contains() didn't find 'Tuesday'")
    }
}

// TestNotFindSunday tests that Sunday is not a weekday
func TestNotFindSunday(t *testing.T) {
    result := Contains(weekdays, "Sunday")
    if result {
        t.Fatal("Contains() did find 'Sunday'")
    }
}

// TestFindEmptyString tests that an empty search string returns 0
func TestFindEmptyString(t *testing.T) {
    result := Contains(weekdays, "")
    if ! result {
        t.Fatal("Contains() didn't find an empty string")
    }
}

// TestNotFindMondayInEmptyString tests that the string Monday is
// not found in the empty string
func TestNotFindMondayInEmptyString(t *testing.T) {
    result := Contains("", "Monday")
    if result {
        t.Fatal("Contains() found 'Monday' in an empty string")
    }
}
