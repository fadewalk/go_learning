









加锁后慢了很多

```bash
wenyu@fadewalk MINGW64 ~/Documents/MyCode/GoProject/GeekTime/go_learning/code/ch48/lock (dev)
$ go test -bench=.
goos: windows
goarch: amd64
pkg: ch48/lock
cpu: AMD Ryzen 7 8845H w/ Radeon 780M Graphics
BenchmarkLockFree-16                 522           2299849 ns/op
BenchmarkLock-16                      10         104159950 ns/op
PASS
ok      ch48/lock       3.926s

```