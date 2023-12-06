package pipeline

// Pipeline represents a sequence of functions to process an input.
type Pipeline struct {
	functions []func(interface{}) (interface{}, error)
}

// NewPipeline creates a new Pipeline.
func NewPipeline() *Pipeline {
	return &Pipeline{}
}

// AddFunction adds a function to the pipeline.
func (p *Pipeline) AddFunction(f func(interface{}) (interface{}, error)) {
	p.functions = append(p.functions, f)
}

// Execute processes the input through the pipeline and returns the final result.
func (p *Pipeline) Execute(input interface{}) (interface{}, error) {
	result := input
	for _, f := range p.functions {
		var err error
		result, err = f(result)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}
