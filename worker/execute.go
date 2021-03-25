package worker

type Executor interface {
	Exec()
}

type ExecuteWithParams struct {
	FuncWithParams
	Params
}

func (e ExecuteWithParams) Exec() {
	e.FuncWithParams(e.Params)
}

type Execute struct {
	Func
}

func (e Execute) Exec() {
	e.Func()
}

type FuncWithParams func(interface{})
type Params interface{}

type Func func() error

type ReadyResponse interface {
	WorkerReady(chan Executor)
}
