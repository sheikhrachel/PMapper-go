package execUtils

import (
	"fmt"
	"runtime"
	"sync"
)

const recoverGoroutineLog = "RecoverGoroutine, skip: %+v. %+v recovered: %+v, file: %+v, line: %+v"

// RecoverGoroutine recovers from a panic in a goroutine and logs the stack trace
func RecoverGoroutine(wg *sync.WaitGroup) {
	defer wg.Done()
	if r := recover(); r != nil {
		for i := 0; i < 10; i++ {
			pc, file, line, ok := runtime.Caller(i)
			if !ok {
				break
			}
			fmt.Println(fmt.Sprintf(recoverGoroutineLog, i, runtime.FuncForPC(pc).Name(), r, file, line))
		}
	}
}
