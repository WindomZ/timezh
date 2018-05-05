package timezh

var longDayNames = [...]string{
	"Sunday",
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
}

var shortDayNames = [...]string{
	"Sun",
	"Mon",
	"Tue",
	"Wed",
	"Thu",
	"Fri",
	"Sat",
}

var shortMonthNames = [...]string{
	"Jan",
	"Feb",
	"Mar",
	"Apr",
	"May",
	"Jun",
	"Jul",
	"Aug",
	"Sep",
	"Oct",
	"Nov",
	"Dec",
}

var longMonthNames = [...]string{
	"January",
	"February",
	"March",
	"April",
	"May",
	"June",
	"July",
	"August",
	"September",
	"October",
	"November",
	"December",
}

const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

var longDayMap = [...]string{
	"星期日",
	"星期一",
	"星期二",
	"星期三",
	"星期四",
	"星期五",
	"星期六",
}

var shortDayMap = [...]string{
	"周日",
	"周一",
	"周二",
	"周三",
	"周四",
	"周五",
	"周六",
}

const (
	January = iota
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)

var shortMonthMap = [...]string{
	"一月",
	"二月",
	"三月",
	"四月",
	"五月",
	"六月",
	"七月",
	"八月",
	"九月",
	"十月",
	"十一月",
	"十二月",
}

var longMonthMap = shortMonthMap

const (
	AM = iota
	PM
)

var shortPMMap = [...]string{
	"上午",
	"下午",
}

var longPMMap = shortPMMap
