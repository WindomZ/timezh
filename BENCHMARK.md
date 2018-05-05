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

>>> 5000000	       218 ns/op	     256 B/op	       4 allocs/op
```

### Analysis
1. 虽然中文化接口损耗部分性能，但换来了更好版本兼容特性
1. 而入侵源码的中文化性能最佳，但这不是本项目期望的设计
