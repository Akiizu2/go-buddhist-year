package buddhistyear

import (
	"strconv"
	"time"
)

func getMonthInThai(monthNum string) string {
	MonthName := []string{"-", "ม.ค.", "ก.พ.", "มี.ค.", "เม.ย.", "พ.ค.", "มิ.ย.", "ก.ค.", "ส.ค.", "ก.ย.", "ต.ค.", "พ.ย.", "ธ.ค."}
	month, err := strconv.Atoi(monthNum)
	if err != nil {
		panic(err)
	}
	return MonthName[month]
}

// ConvertTimeToBuddhist => Convert Unix Time to Buddhist Year Format
func ConvertTimeToBuddhist(originalTime int) string {
	timeUnix := originalTime + (3600 * 7) // <-- Convert to GMT+7 Time
	if timeUnix == 0 {
		return "-"
	}
	day := time.Unix(int64(timeUnix), 0).Format("2")
	month := time.Unix(int64(timeUnix), 0).Format("01")
	year := time.Unix(int64(timeUnix), 0).Format("2006")
	yearInt, _ := strconv.Atoi(year)
	bdYear := yearInt + 543
	return day + " " + getMonthInThai(month) + " " + strconv.Itoa(bdYear)
}
