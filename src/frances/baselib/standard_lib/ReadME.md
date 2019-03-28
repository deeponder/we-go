1. 文件压缩： tar/zip
2. 文本的读写（按分隔符）：bufio
3. bytes: 与string比较，需要单独处理字符时， 用byte数组。 类别c的char
4. container: heap(优先队列), list(双向链表), ring(循环链表)
5. context: 控制go出去的协程， 协程间传递元数据。 可替代chan+select
6. crypto: 加密算法 参见baselib/crypto 库及其demo
7. errors: fmt.Errorf("%d",1)
8. sql: 参见baselib/mysqlapi 库及其demo
9. encoding: base64参见baselib/crypto/crypto.go， json， xml, csv..
10. flag: 解析命令行参数  os & flag  [示例](http://www.01happy.com/golang-command-line-arguments/)
11. fmt: 格式化输入输出
12. go： go的一些底层库ast等， 解析编译
13. html: xss过滤。 EscapeString && UnescapeString
14. io: 一切皆文件； 网络、文件、控制台； 实现了io.Reader  && io.Writer皆可称之为文件。
    https://colobu.com/2016/08/29/go-io-Reader-and-io-Writer/
15. math: baselib/math/rand.go
16. mime: 邮件格式
17. os: shell命令；信号；
18. path: 文件路径的一些操作
19. reflect: interface{}; reflect.TypeOf(), reflect.ValueOf(); baselib/logs/tlogxxx
20. regexp: 正则；regexp.Match, regexp.Compile(pattern); baselib/_demo/regexp.go
21. runtime: 线程调度(scheduler) && GC
22. sort: int string float原生支持，不用实现Len(),Swap(), Less()
23. strconv: 字符串转换操作
24. strings: 字符串操作， split, contains, compare, Join等
25. sync: lock, Once, WaitGroup, 线程安全的Map
26. testing: 配合 go test 实现单元测试和性能测试  https://github.com/deeponder/build-web-application-with-golang/blob/master/zh/11.3.md
27. time: baselib/time, time.Now().Format("2006-01-02 15:04:05")