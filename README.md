# timezh - Chinese style time

> Golang库 - 时间格式中文化

[![Build Status](https://travis-ci.org/WindomZ/timezh.svg?branch=master)](https://travis-ci.org/WindomZ/timezh)
[![Coverage Status](https://coveralls.io/repos/github/WindomZ/timezh/badge.svg?branch=master)](https://coveralls.io/github/WindomZ/timezh?branch=master)
[![GoDoc](https://godoc.org/github.com/WindomZ/timezh?status.svg)](https://godoc.org/github.com/WindomZ/timezh)
[![Go Report Card](https://goreportcard.com/badge/github.com/WindomZ/timezh)](https://goreportcard.com/report/github.com/WindomZ/timezh)

扩展原生`time`快速实现时间格式中文化

## Install
```bash
go get github.com/WindomZ/timezh
```

### Roadmap
- [x] Format
  - [x] Format(string) string
  - [x] FormatMix(string) string
  - [x] FormatLayout(string) string
  - [x] FormatChinese(string) string
- [x] Parse
  - [x] Parse(string, string) (Time, error)
  - [x] ParseLayout(string) string
  - [x] ParseValue(string) string

## Usage
|格式用途|原生`time`|引入`timezh`|
|:-----:|:-----:|:-----:|
|月份|January, Jan|一月|
|星期|Monday, Mon|星期一, 周一|
|上下午|PM, pm|下午|

`time`库**不支持**中文化：
```go
import "time"

time.Now().Format("2006年01月02日(January) PM3:04:05 Monday(Mon)")

>>> 2009年01月03日(January) PM6:15:05 Saturday(Sat)
```

`timezh`库**支持**中文化：
```go
import "github.com/WindomZ/timezh"

timezh.Now().Format("2006年01月02日(一月) 下午3:04:05 星期一(周一)")

>>> 2009年01月03日(一月) 下午6:15:05 星期六(周六)
```

`timezh`库**支持**中文解析：
```go
import "github.com/WindomZ/timezh"

timezh.Parse("2006年01月02日(一月) 下午3:04:05 星期一(周一)", "2009年01月03日(一月) 下午6:15:05 星期六(周六)")

time.Now().Format("2006年01月02日(January) PM3:04:05 Monday(Mon)")
timezh.Now().Format("2006年01月02日(一月) 下午3:04:05 星期一(周一)")

>>> 2009年01月03日(January) PM6:15:05 Saturday(Sat)
>>> 2009年01月03日(一月) 下午6:15:05 星期六(周六)
```

### Advanced
混杂中英文：
```go
import "github.com/WindomZ/timezh"

timezh.Now().FormatMix("2006年01月02日(January, 一月) 下午3:04:05PM 星期一(Mon, 周一)")

>>> 2009年01月03日(January, 一月) 下午6:15:05PM 星期六(Sat, 周六)
```

文本中文化：
```go
import "github.com/WindomZ/timezh"

timezh.FormatChinese("2009年01月03日(January) 6:15:05PM 星期六(Sat)")

>>> 2009年01月03日(一月) 6:15:05下午 星期六(周六)
```

## Benchmark
详见[BENCHMARK](BENCHMARK.md)

## Contributing
欢迎你Fork，提交PR，在[issues page](https://github.com/WindomZ/timezh/issues)汇报Bugs、提出新想法等，
我很乐意能一起参与。

如果你喜欢这个项目，可以点下 :star: 予以支持，谢谢！

## 许可
[MIT](https://github.com/WindomZ/timezh/blob/master/LICENSE)
