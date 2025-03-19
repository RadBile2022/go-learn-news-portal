package dateutils

import "time"

func GenerateMonthlyNamesGrouped(startDate, endDate time.Time, length int) (int, [][]string) {
	if length == 0 {
		return 0, nil
	}

	// Normalize dates to first day of month to ensure proper comparison
	start := time.Date(startDate.Year(), startDate.Month(), 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(endDate.Year(), endDate.Month(), 1, 0, 0, 0, 0, time.UTC)

	// Handle invalid date range
	if end.Before(start) {
		return 0, nil
	}

	gap := 0

	// Generate all months first
	var allMonths []string
	current := start
	for !current.After(end) {
		allMonths = append(allMonths, current.Format("Jan 2006"))
		current = current.AddDate(0, 1, 0)
		gap += 1
	}

	// Group months into 6-month periods
	var result [][]string
	for i := 0; i < len(allMonths); i += length {
		end := i + length
		if end > len(allMonths) {
			end = len(allMonths)
		}
		result = append(result, allMonths[i:end])
	}

	return gap, result
}
