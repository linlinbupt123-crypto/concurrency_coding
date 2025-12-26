panic vs fatal

https://logeshwrites.hashnode.dev/error-handling-in-go-understanding-panic-recover-and-logfatal
panic error

1,Invalid operations (e.g., index out of bounds)

2,Invalid states (e.g., corrupted data)

3,Critical errors where proceeding would make the program unsafe

panic是可以recover, recover可以不退出程序，继续执行。

log.Fatal: Logging Critical Errors and Exiting

fatal error是不能recover的，直接退出程序。比如map在没有锁的情况下，并发的去写。会报fatal error，程序退出。