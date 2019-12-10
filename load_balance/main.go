package main

//负载均衡

type LB struct {
	availableWorkerChan chan func()
}

func NewLB(workers ...int) *LB {
	return &LB{}
}

func (lb *LB) submit() {

}

type Worker struct {
	workerCount int
	workerChan  chan func()
}

func (w *Worker) Ready(f func()) bool {

}

func main() {
	lb := NewLB(1, 1, 2, 1)

}
