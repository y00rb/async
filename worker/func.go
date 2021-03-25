package worker

type Func func(interface{})

type Params interface{}

type ReadyResponse interface {
	WorkerReady(chan Params)
}
