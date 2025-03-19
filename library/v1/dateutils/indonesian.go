package dateutils

import (
	"fmt"
	"strconv"
	"time"
)

func SimpleFormatInBahasa(t time.Time) string {
	now := t.UTC().Add(time.Hour * 7).String()
	year, _ := strconv.ParseUint(now[0:4], 10, 0)
	month, _ := strconv.ParseUint(now[5:7], 10, 0)
	day, _ := strconv.ParseUint(now[8:10], 10, 0)

	return fmt.Sprintf("%d %s %d", day, getMonthInBahasa(uint8(month)), year)
}

func getDayInBahasa(t time.Time) string {
	switch t.Weekday() {
	case 0:
		return "Minggu"
	case 1:
		return "Senin"
	case 2:
		return "Selasa"
	case 3:
		return "Rabu"
	case 4:
		return "Kamis"
	case 5:
		return "Jum'at"
	case 6:
		return "Sabtu"
	}
	return ""
}

func getMonthInBahasa(month uint8) string {
	switch month {
	case 1:
		return "Januari"
	case 2:
		return "Februari"
	case 3:
		return "Maret"
	case 4:
		return "April"
	case 5:
		return "Mei"
	case 6:
		return "Juni"
	case 7:
		return "Juli"
	case 8:
		return "Agustus"
	case 9:
		return "September"
	case 10:
		return "Oktober"
	case 11:
		return "November"
	case 12:
		return "Desember"
	}
	return ""
}
