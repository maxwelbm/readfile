# readfile

### Result

command exec: `go test -bench .`
```shell
goos: darwin
goarch: arm64
pkg: github.com/maxwelbm/benchReadFile/pkg/readfile
BenchmarkReadTxtFile-8         	   56223	     17935 ns/op
BenchmarkReadJsonLinesFile-8   	  163551	      7244 ns/op
PASS
ok  	github.com/maxwelbm/benchReadFile/pkg/readfile	2.896s
```
