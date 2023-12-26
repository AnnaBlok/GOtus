package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	var out Bi
	for _, stage := range stages {
		select {
		case <-done:
			break
		default:
			out = make(Bi)
			go executeStage(stage, in, out, done)
			in = out
		}
	}
	return out
}

func executeStage(stage Stage, in In, out Bi, done In) {
	defer close(out)
	stageOut := stage(in)
	for {
		select {
		case <-done:
			return
		case v, ok := <-stageOut:
			if !ok {
				return
			}
			out <- v
		}
	}
}
