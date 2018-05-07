package main

import (
	"fmt"
	"time"

	"github.com/WindomZ/timezh"
)

func main() {
	t, _ := time.Parse("2006-01-02 15:04:05", "2009-01-03 18:15:05")

	fmt.Println("Time:", t.Format("2006-01-02 15:04:05"))

	format(t)
	parse(t)
}

func format(t time.Time) {
	fmt.Println("===========Format===========")
	fmt.Println("\"time\"效果：")
	fmt.Println("  - time.Format(英文):  ", t.Format("2006-01-02(January) PM3:04:05 Monday(Mon)"))
	fmt.Println("  - time.Format(中文):  ", t.Format("2006年01月02日(January) PM3:04:05 Monday(Mon)"))

	fmt.Println("\"timezh\"同样可以：")
	fmt.Println("  - timezh.Format(英文):", timezh.T(t).Format("2006-01-02(January) PM3:04:05 Monday(Mon)"))
	fmt.Println("  - timezh.Format(中文):", timezh.T(t).Format("2006年01月02日(January) PM3:04:05 Monday(Mon)"))

	fmt.Println("\"timezh\"还可以：")
	fmt.Println("  - timezh.Format(特性):", timezh.T(t).Format("2006年01月02日(一月) 下午3:04:05 星期一(周一)"))
	fmt.Println("  - timezh.Format(特性):", timezh.T(t).Format("2006年01月02日(January, 一月) 下午3:04:05PM 星期一(Mon, 周一)"))
	fmt.Println("  - timezh.Format(特性):", timezh.T(t).FormatMix("2006年01月02日(January, 一月) 下午3:04:05PM 星期一(Mon, 周一)"))
}

func parse(t time.Time) {
	tmp, _ := timezh.Parse("2006年01月02日(一月) 下午3:04:05 星期一(周一)",
		timezh.T(t).Format("2006年01月02日(一月) 下午3:04:05 星期一(周一)"))

	fmt.Println("===========Parse===========")
	fmt.Println("Unix time对比：", t.UnixNano(), "==", tmp.UnixNano())
}
