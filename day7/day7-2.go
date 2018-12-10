package main

type worker struct {
	workid      int
	workingTime int
	workingTask rune
}

func (w *worker) work(s *step) {
	//println(w.workid, string(s.name))
	s.finish()
	w.workingTime = s.endTime
}

func part2() {
	//读取文件
	filePath := "./input.txt"
	lines := make(chan string)
	go readFile(filePath, lines)

	//按行读取信息
	//创建map[byte]*step
	stepMap := lines2StepMap(lines)

	//取出一个可执行任务
	//找到该任务开始执行时间在worker完成时间之后，且最早完成的worker，将任务指派给他
	//更新任务完成时间
	workers := make([]*worker, 0)
	for i := 0; i < 2; i++ {
		var wt worker
		wt.workid = i
		wt.workingTime = 0
		workers = append(workers, &wt)
	}
	for len(stepMap) > 0 {
		for k, v := range stepMap {
			if v.preStepCount == 0 {
				freeWorker := workers[0]
				for _, w := range workers {
					if (w.workingTime < freeWorker.workingTime) && (v.beginTime >= w.workingTime) {
						freeWorker = w
					}
				}
				freeWorker.work(v)
				delete(stepMap, k)
			}
		}
	}

	for _, w := range workers {
		println(w.workingTime)
	}
}
