


终端命令

执行

```bash

wenyu@fadewalk MINGW64 ~/Documents/MyCode/GoProject/GeekTime/go_learning/code/ch47 (dev)
$ go tool pprof cpu.prof
BenchmarkProcessRequestOld-16              57998             21770 ns/op
PASS
ok      go_learning/code/ch47   3.142s

wenyu@fadewalk MINGW64 ~/Documents/MyCode/GoProject/GeekTime/go_learning/code/ch47 (dev)
$ go test -bench=. -cpuprofile=cpu.prof
goos: windows
goarch: amd64
pkg: go_learning/code/ch47
cpu: AMD Ryzen 7 8845H w/ Radeon 780M Graphics
BenchmarkProcessRequest-16                176965              6186 ns/op
BenchmarkProcessRequestOld-16              52232             24852 ns/op
PASS
ok      go_learning/code/ch47   3.308s

wenyu@fadewalk MINGW64 ~/Documents/MyCode/GoProject/GeekTime/go_learning/code/ch47 (dev)
$ go tool pprof cpu.prof 
File: ch47.test.exe
Build ID: C:\Users\wenyu\AppData\Local\Temp\go-build3217277683\b001\ch47.test.exe2025-02-23 17:02:58.0296329 +0800 CST
Type: cpu
Time: Feb 23, 2025 at 5:02pm (CST)
Duration: 2.85s, Total samples = 3.70s (129.78%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 1510ms, 40.81% of 3700ms total
Dropped 80 nodes (cum <= 18.50ms)
Showing top 10 nodes out of 186
      flat  flat%   sum%        cum   cum%
     280ms  7.57%  7.57%      280ms  7.57%  runtime.stdcall2
     190ms  5.14% 12.70%      190ms  5.14%  runtime.memmove
     140ms  3.78% 16.49%      290ms  7.84%  github.com/mailru/easyjson/jlexer.(*Lexer).FetchToken
     140ms  3.78% 20.27%      460ms 12.43%  runtime.mallocgc
     140ms  3.78% 24.05%      270ms  7.30%  runtime.scanobject
     140ms  3.78% 27.84%      140ms  3.78%  runtime.stdcall3
     140ms  3.78% 31.62%      280ms  7.57%  strconv.ParseInt
     140ms  3.78% 35.41%      140ms  3.78%  strconv.ParseUint
     100ms  2.70% 38.11%      430ms 11.62%  runtime.concatstrings
     100ms  2.70% 40.81%      100ms  2.70%  runtime.stdcall1
(pprof) top -cum
Showing nodes accounting for 0.09s, 2.43% of 3.70s total
Dropped 80 nodes (cum <= 0.02s)
Showing top 10 nodes out of 186
      flat  flat%   sum%        cum   cum%
         0     0%     0%      2.40s 64.86%  testing.(*B).launch
         0     0%     0%      2.40s 64.86%  testing.(*B).runN
         0     0%     0%      1.32s 35.68%  go_learning/code/ch47.BenchmarkProcessRequestOld
     0.02s  0.54%  0.54%      1.32s 35.68%  go_learning/code/ch47.processRequestOld
     0.01s  0.27%  0.81%      1.15s 31.08%  runtime.systemstack
         0     0%  0.81%      1.08s 29.19%  go_learning/code/ch47.BenchmarkProcessRequest
     0.01s  0.27%  1.08%      1.08s 29.19%  go_learning/code/ch47.processRequest
         0     0%  1.08%      0.94s 25.41%  go_learning/code/ch47.(*Request).UnmarshalJSON (partial-inline)
     0.05s  1.35%  2.43%      0.94s 25.41%  go_learning/code/ch47.easyjson6a975c40DecodeCh471
         0     0%  2.43%      0.72s 19.46%  runtime.gcBgMarkWorker.func2
(pprof) list processRequest
Total: 3.70s
ROUTINE ======================== go_learning/code/ch47.processRequest in C:\Users\wenyu\Documents\MyCode\GoProject\GeekTime\go_learning\code\ch47\optmization.go
      10ms      1.08s (flat, cum) 29.19% of Total
         .          .     22:func processRequest(reqs []string) []string {
         .          .     23:   reps := []string{}
         .          .     24:   for _, req := range reqs {
         .          .     25:           reqObj := &Request{}
         .      710ms     26:           reqObj.UnmarshalJSON([]byte(req))
         .          .     27:           //      json.Unmarshal([]byte(req), reqObj)
         .          .     28:
         .          .     29:           var buf strings.Builder
      10ms       10ms     30:           for _, e := range reqObj.PayLoad {
         .      220ms     31:                   buf.WriteString(strconv.Itoa(e))
         .       60ms     32:                   buf.WriteString(",")
         .          .     33:           }
         .          .     34:           repObj := &Response{reqObj.TransactionID, buf.String()}
         .       40ms     35:           repJson, err := repObj.MarshalJSON()
         .          .     36:           //repJson, err := json.Marshal(&repObj)
         .          .     37:           if err != nil {
         .          .     38:                   panic(err)
         .          .     39:           }
         .       40ms     40:           reps = append(reps, string(repJson))
         .          .     41:   }
         .          .     42:   return reps
         .          .     43:}
         .          .     44:
         .          .     45:func processRequestOld(reqs []string) []string {
ROUTINE ======================== go_learning/code/ch47.processRequestOld in C:\Users\wenyu\Documents\MyCode\GoProject\GeekTime\go_learning\code\ch47\optmization.go
      20ms      1.32s (flat, cum) 35.68% of Total
         .          .     45:func processRequestOld(reqs []string) []string {
         .          .     46:   reps := []string{}
         .          .     47:   for _, req := range reqs {
         .          .     48:           reqObj := &Request{}
         .      510ms     49:           json.Unmarshal([]byte(req), reqObj)
         .          .     50:           ret := ""
      10ms       10ms     51:           for _, e := range reqObj.PayLoad {
      10ms      540ms     52:                   ret += strconv.Itoa(e) + ","
         .          .     53:           }
         .       10ms     54:           repObj := &Response{reqObj.TransactionID, ret}
         .      250ms     55:           repJson, err := json.Marshal(&repObj)
         .          .     56:           if err != nil {
         .          .     57:                   panic(err)
         .          .     58:           }
         .          .     59:           reps = append(reps, string(repJson))
         .          .     60:   }
(pprof) \q
unrecognized command: "\\q"
(pprof) ^D
unrecognized command: "\x04"
(pprof) exit

wenyu@fadewalk MINGW64 ~/Documents/MyCode/GoProject/GeekTime/go_learning/code/ch47 (dev)
$ go test -bench=. -memprofile=mem.prof
goos: windows
goarch: amd64
pkg: go_learning/code/ch47
cpu: AMD Ryzen 7 8845H w/ Radeon 780M Graphics
BenchmarkProcessRequest-16                175840              6373 ns/op
BenchmarkProcessRequestOld-16              50572             21060 ns/op
PASS
ok      go_learning/code/ch47   2.867s

wenyu@fadewalk MINGW64 ~/Documents/MyCode/GoProject/GeekTime/go_learning/code/ch47 (dev)
$ go tool pprof mem.prof
File: ch47.test.exe
Build ID: C:\Users\wenyu\Documents\MyCode\GoProject\GeekTime\go_learning\code\ch47\ch47.test.exe2025-02-23 17:03:01.3488803 +0800 CST
Type: alloc_space
Time: Feb 23, 2025 at 5:09pm (CST)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 1969.67MB, 99.59% of 1977.69MB total
Dropped 21 nodes (cum <= 9.89MB)
Showing top 10 nodes out of 29
Showing top 10 nodes out of 29
      flat  flat%   sum%        cum   cum%
  921.68MB 46.60% 46.60%  1127.29MB 57.00%  go_learning/code/ch47.processRequestOld
  476.32MB 24.08% 70.69%   481.32MB 24.34%  go_learning/code/ch47.easyjson6a975c40DecodeCh471
  197.57MB  9.99% 80.68%   197.57MB  9.99%  strings.(*Builder).WriteString (inline)
  118.04MB  5.97% 86.65%   849.89MB 42.97%  go_learning/code/ch47.processRequest
   93.02MB  4.70% 91.35%    93.02MB  4.70%  github.com/mailru/easyjson/buffer.getBuf
   88.03MB  4.45% 95.80%    88.03MB  4.45%  github.com/mailru/easyjson/buffer.(*Buffer).BuildBytes
   43.51MB  2.20% 98.00%   136.53MB  6.90%  github.com/mailru/easyjson/buffer.(*Buffer).ensureSpaceSlow
   21.51MB  1.09% 99.09%    73.52MB  3.72%  encoding/json.Marshal
      10MB  0.51% 99.59%   132.08MB  6.68%  encoding/json.Unmarshal
         0     0% 99.59%   120.08MB  6.07%  encoding/json.(*decodeState).object
(pprof)









```