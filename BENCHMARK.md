# Benchmark

## Format
`time`库，数据来自`BenchmarkTime_Format1`:
```
time.Now().Format("2006年01月02日 January 3:04:05PM Mon")

>>> 20000000	        78.5 ns/op	      48 B/op	       1 allocs/op
```

`timezh`库，数据来自`BenchmarkTime_Format2`:
```
timezh.Now().Format("2006年01月02日 一月 3:04:05下午 星期一")

>>> 5000000	       218 ns/op	     256 B/op	       4 allocs/op
```

### Analysis
1. 虽然中文化接口损耗部分性能，但换来了更好版本兼容特性
1. 而入侵源码的中文化性能最佳，但这不是本项目期望的设计

## Parse
`time`库，数据来自`BenchmarkParse1`:
```
time.Parse("2006年01月02日(January) PM3:04:05 Monday(Mon)", "2009年01月03日(Jan) PM6:15:05 Saturday(Sat)")

>>> 20000000	        97.5 ns/op	      80 B/op	       1 allocs/op
```

`timezh`库，数据来自`BenchmarkParse2`:
```
timezh.Parse("2006年01月02日(January) PM3:04:05 Monday(Mon)", "2009年01月03日(一月) 下午6:15:05 星期六(周六)")

>>> 5000000	       331 ns/op	     288 B/op	       4 allocs/op
```

`time`+`timezh`库，数据来自`BenchmarkParse3`:
```
time.Parse("2006年01月02日(January) PM3:04:05 Monday(Mon)", timezh.ParseValue("2009年01月03日(Jan) PM6:15:05 Saturday(Sat)"))

>>> 5000000	       221 ns/op	     240 B/op	       3 allocs/op
```

### Analysis
1. 第二个方案的中文解析过程需对两个参数分别进行解析
1. 第三个方案只进行一次中文解析，会带来不少性能提升
