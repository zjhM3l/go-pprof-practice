package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"time"

	"github.com/zjhM3l/go-pprof-practice/animal"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.SetOutput(os.Stdout)

	runtime.GOMAXPROCS(1)
	runtime.SetMutexProfileFraction(1)
	runtime.SetBlockProfileRate(1)

	// 整体性能分析：运行go run main.go后，访问http://localhost:6060/debug/pprof/可以查看pprof信息
	
	// 分析cpu（整体性能分析运行过程中）：
	// 1. 运行go tool pprof http://localhost:6060/debug/pprof/profile?seconds=10可以查看cpu profile
	// 2. topN可以查看占用资源最多的函数
	// 指标：flat：当前函数本身的执行耗时; flat%：当前函数本身的执行耗时占cpu总耗时的百分比; 
	// sum%：上面每一行的flat%的累加和; cum：当前函数加上调用它的函数的执行耗时; cum%：cum占cpu总耗时的百分比
	// Flat == Cum说明当前函数没有调用其他函数，Flat > Cum说明当前函数调用了其他函数
	// Flat == 0说明当前函数中只有其他函数的调用，Cum == 0说明当前函数没有被调用
	// 结果分析：tiger.(*Tiger).Eat占用了大部分cpu时间，说明是性能瓶颈
	// 3. list可以查看指定函数的详细信息 list Eat可以查看Eat函数的详细信息
	// 可以定位到实际的代码行，找到性能瓶颈
	// 结果分析 for i := 0; i < loop; i++是性能瓶颈，可以优化
	// cpu分析命令输入之后可以用web命令可以调用关系可视化图，更直观
	
	// 分析堆内存（整体性能分析运行过程中）：
	// go tool pprof -http=:8080 http://localhost:6060/debug/pprof/heap可以查看堆内存信息
	// 出现浏览器页面
	// VIEW菜单可以查看不同的视图，比如Top，Graph，Flame Graph，Peek，Source，Disasm
	// 结果分析：mouse.(*Mouse).Steal的line 50占用了大部分内存，说明是内存瓶颈
	// SAMPLE里面有四个指标说明内存分配的情况，inuse_space：当前内存使用量; inuse_objects：当前内存对象数量; 
	// alloc_space：分配（程序累计申请）的内存量; alloc_objects：分配的对象数量
	// 结果分析：在alloc_space中可以看到dog.(*Dog).Run的43行一直在申请内存，然后也没有用到，说明是内存泄漏
	
	// Groutine分析（整体性能分析运行过程中）：
	// go tool pprof -http=:8080 http://localhost:6060/debug/pprof/goroutine可以查看goroutine信息
	// VIEW的Flame Graph可以查看goroutine的调用关系平铺
	// Flame Graph：由上到下表示调用顺序，每一块代表一个函数，越长代表占用CPU时间越长，火焰图是动态的，支持点击块进行分析
	// 结果分析：wolf.(*Wolf).Drink.func1的34行占用了大部分goroutine

	// mutex锁分析（整体性能分析运行过程中）：
	// go tool pprof -http=:8080 http://localhost:6060/debug/pprof/mutex可以查看mutex锁信息
	// 结果分析：wolf.(*Wolf).Howl.func1的58行

	// 分析block阻塞（整体性能分析运行过程中）：
	// go tool pprof -http=:8080 http://localhost:6060/debug/pprof/block可以查看block阻塞信息
	// 结果分析：cat.(*Cat).Pee的39行
	// 额外分析，执行go tool pprof http://localhost:6060/debug/pprof/block的话可以看到终端，显示dropped 4 nodes (cum <= 1.41s)
	// block有2，但是加上http网站只显示一个，说明有一个被drop了，想看的话可以直接从http的block链接进去看
	go func() {
		if err := http.ListenAndServe(":6060", nil); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()

	for {
		for _, v := range animal.AllAnimals {
			v.Live()
		}
		time.Sleep(time.Second)
	}
}
