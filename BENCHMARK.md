## Benchmark

`time`库，数据来自`BenchmarkTime_Format1`:
```
import "time"

time.Now().Format("2006年01月02日 January 3:04:05PM Mon")

>>> 20000000	        78.5 ns/op	      48 B/op	       1 allocs/op
```

`timezh`库，数据来自`BenchmarkTime_Format2`:
```
import "github.com/WindomZ/timezh"

timezh.Now().Format("2006年01月02日 一月 3:04:05下午 星期一")

>>> 5000000	       256 ns/op	     368 B/op	       4 allocs/op
```

### Analysis
1. 虽然中文化损耗了部分性能，但换来了版本兼容特性
1. 入侵源码中文化的性能最佳，但这不是该项目期望的
