package executor

type Options struct {
	TraverseNestedDirectories bool
}

func NewDefaultOptions() *Options {
	return &Options{TraverseNestedDirectories: true}
}

func (options *Options) EnableNestedTraversal() *Options {
	options.TraverseNestedDirectories = true
	return options
}

func (options *Options) DisableNestedTraversal() *Options {
	options.TraverseNestedDirectories = false
	return options
}
