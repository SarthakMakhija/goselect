package executor

type Options struct {
	TraverseNestedDirectories bool
}

func NewDefaultOptions() *Options {
	return &Options{TraverseNestedDirectories: true}
}

func OptionsWith(nestedTraversal bool) *Options {
	return &Options{TraverseNestedDirectories: nestedTraversal}
}
