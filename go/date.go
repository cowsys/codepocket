package main

import "fmt"
import "time"

func main() {
	// print year, month, day
	now := time.Now()
	year, month, day := now.Date()
	fmt.Println(year, month, day)
	// ->2016 January 24

	// print hour, min, sec
	hour, min, sec := now.Clock()
	fmt.Println(hour, min, sec)
	// ->18 20 42

	// format DATE:(2015/06/03 09:09:09)①
	//// func Date(
	//// year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time {
	//// Localはtime/zoneinfoで定義されている
	ot := time.Date(2016, 6, 3, 9, 9, 9, 0, time.Local)
	year, month, day = ot.Date()
	//// func (t Time) Date() (year int, month Month, day int) {
	hour, min, sec = ot.Clock()
	fmt.Printf("%v/%v/%v %v:%v:%v\n",
		year, int(month), day,
		hour, min, sec,
	)
	// format DATE:(2015/06/03 09:09:09)②
	fmt.Printf("%4d/%02d/%02d %02d:%02d:%02d\n",
		year, int(month), day,
		hour, min, sec,
	)

	// format DATE:(2015年6月3日({金}) x時x分x秒)
	{
		japanesewday := [...]string{"日", "月", "火", "水", "木", "金", "土"}
		fmt.Printf("%4d年%d月%d日(%s) %d時%d分%d秒\n",
			year, int(month), day, japanesewday[ot.Weekday()],
			hour, min, sec,
		)
	}

	{
		// simply use Year, Month, Day, Weekday
		fmt.Println(now.Year())
		fmt.Println(now.Month())
		fmt.Println(now.Day())
		fmt.Println(now.Weekday())
		// 2016
		// January
		// 24
		// Sunday
	}
	{
		// simply use YearDay
		fmt.Println(now.YearDay())
	}

	{
		// TASK: challenge tasks
		// calculate time
		//// add
		//// minus
		// change time zone
		// use After , Before, Equal
	}
}
