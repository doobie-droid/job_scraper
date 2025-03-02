package utilities

import (
	"fmt"
	"regexp"
	"time"
)

// receives dateString in the format January 2028 and returns whether it exists in this month or the previous
func IsLessThanTwoMonths(dateString string) (bool, error) {

	fmt.Println(removeParentheses(dateString))
	layout := "2 January 2006"

	parsedTime, err := time.Parse(layout, fmt.Sprintf("28 %s", dateString))
	if err != nil {
		return false, fmt.Errorf("invalid date format: %v", err)
	}

	oneMonthAgo := time.Now().AddDate(0, -1, 0)

	return parsedTime.Year() == oneMonthAgo.Year() &&
		(parsedTime.Month() == oneMonthAgo.Month() || parsedTime.Month() == time.Now().Month()), nil
}

func removeParentheses(input string) string {
	re := regexp.MustCompile(`[\(\)]`)
	return re.ReplaceAllString(input, "")
}
