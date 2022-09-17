package bench

import (
	"goselect/parser/context"
	"testing"
)

func BenchmarkLikeMatch(b *testing.B) {
	functions := context.NewFunctions()
	for index := 0; index < b.N; index++ {
		_, err := functions.Execute("like", context.StringValue("sample.log"), context.StringValue(".*.log"))
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkLikeNotMatch(b *testing.B) {
	functions := context.NewFunctions()
	for index := 0; index < b.N; index++ {
		_, err := functions.Execute("like", context.StringValue("README.md"), context.StringValue(".*.log"))
		if err != nil {
			panic(err)
		}
	}
}
