package workpool1

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type ProccessorFunc func(resource interface{}) error

type ResultProcessorFunc func(result Result) error

type Job struct {
	Id       int
	resource interface{}
}

type Result struct {
	Job
	Err error
}

type Pool struct {
	Size      int
	Jobs      chan Job
	Results   chan Result
	Done      chan bool
	completed bool
}

func NewPool(size int) *Pool {
	log.Print("Creating a new Pool")
	r := &Pool{Size: size}
	r.Jobs = make(chan Job, size)
	r.Results = make(chan Result, size)
	return r
}

func (p *Pool) Start(resources []interface{}, procFunc ProccessorFunc, resultProcFunc ResultProcessorFunc) {
	startTime := time.Now()
	go p.Allocate(resources)
	p.Done = make(chan bool)
	go p.Collect(resultProcFunc)
	go p.WorkPool(procFunc)
	<-p.Done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	log.Printf("total time taken: [%f] seconds", diff.Seconds())
}

func (p *Pool) Allocate(resources []interface{}) {
	defer p.closeJobChan()
	log.Printf("Allocating [%d] resources", len(resources))
	for i, r := range resources {
		job := Job{Id: i, resource: r}
		fmt.Println("sending ", i)
		p.Jobs <- job
	}
	fmt.Println("allocte done.")
}

func (p *Pool) Exec(wg *sync.WaitGroup, procFunc ProccessorFunc) {
	defer wg.Done()
	for j := range p.Jobs {
		output := Result{j, procFunc(j.resource)}
		p.Results <- output
	}
	fmt.Println("goRoutine work done.")
}

func (p *Pool) WorkPool(processor ProccessorFunc) {
	defer p.closeResultChan()
	var wg sync.WaitGroup
	for i := 0; i < p.Size; i++ {
		fmt.Println(i)
		wg.Add(1)
		go p.Exec(&wg, processor)
		log.Printf("Spawned work goRoutine [%d]", i)
	}
	log.Print("Worker Pool done spawning work goRoutines")
	wg.Wait()
	log.Println("jobs are finished")
}

func (p *Pool) closeJobChan() {
	fmt.Println("closing job channel")
	close(p.Jobs)
}

func (p *Pool) closeResultChan() {
	fmt.Println("closing result channel")
	close(p.Results)
}

func (p *Pool) Collect(proc ResultProcessorFunc) {
	for result := range p.Results {
		outcome := proc(result)
		log.Printf("Job with id: [%d] completed, outcome: %s", result.Job.Id, outcome)
	}
	log.Print("goRoutine collect done, setting channel done as completed")
	p.Done <- true
	p.completed = true
}

func (p *Pool) IsCompleted() bool {
	return p.completed
}
