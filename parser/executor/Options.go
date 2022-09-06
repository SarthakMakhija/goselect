package executor

import "strings"

type Options struct {
	traverseNestedDirectories    bool
	directoriesToIgnoreTraversal map[string]bool
}

func NewDefaultOptions() *Options {
	return &Options{traverseNestedDirectories: true}
}

func (options *Options) EnableNestedTraversal() *Options {
	options.traverseNestedDirectories = true
	return options
}

func (options *Options) DisableNestedTraversal() *Options {
	options.traverseNestedDirectories = false
	return options
}

func (options *Options) DirectoriesToIgnoreTraversal(names []string) *Options {
	directoriesToIgnore := make(map[string]bool)
	for _, directory := range names {
		directoriesToIgnore[strings.ToLower(directory)] = true
	}
	options.directoriesToIgnoreTraversal = directoriesToIgnore
	return options
}

func (options Options) IsDirectoryTraversalIgnored(name string) bool {
	return options.directoriesToIgnoreTraversal[strings.ToLower(name)]
}
