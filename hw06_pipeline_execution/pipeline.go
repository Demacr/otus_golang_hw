package hw06_pipeline_execution //nolint:golint,stylecheck

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func wrapStage(stage Stage, in In, done In) (out Out) {
	shimCh := make(Bi)
	go func() {
		defer close(shimCh)
		for {
			select {
			case <-done:
				return
			case v, ok := <-in:
				if ok {
					shimCh <- v
				} else {
					return
				}
			}
		}
	}()
	return stage(shimCh)
}

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	outCh := in
	for _, stage := range stages {
		outCh = wrapStage(stage, outCh, done)
	}
	return outCh
}
