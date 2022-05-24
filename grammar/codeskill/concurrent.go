package codeskill

type SendMsg string

func dispatch(in chan *SendMsg, out chan chan *SendMsg) {
	var taskQ []*SendMsg
	var workerQ []chan *SendMsg
	for {
		var activeTask *SendMsg
		var activeWorker chan *SendMsg
		if len(taskQ) > 0 && len(workerQ) > 0 {
			activeTask = taskQ[0]
			activeWorker = workerQ[0]
		}
		select {
		case t := <-in:
			taskQ = append(taskQ, t)
		case w := <-out:
			workerQ = append(workerQ, w)
		case activeWorker <- activeTask:
			taskQ = taskQ[1:]
			workerQ = workerQ[1:]
		}
	}
}

type OperatorFunc func(msg *SendMsg)

func creatWorker(worker chan *SendMsg, out chan chan *SendMsg, f OperatorFunc) {
	for {
		select {
		case t := <-worker:
			out <- worker
			f(t)
		}
	}
}

func ConcurrentFromChan(in chan *SendMsg, num int, f OperatorFunc) {
	// 分发
	out := make(chan chan *SendMsg, num)
	go dispatch(in, out)

	// 生产者
	for i := 0; i < num; i++ {
		worker := make(chan *SendMsg)
		out <- worker
		go creatWorker(worker, out, f)
	}
}
