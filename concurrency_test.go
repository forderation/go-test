package main

import (
	"fmt"
	"reflect"
	"testing"

	support "github.com/forderation/go-test/support"
)

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}
	want := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds":    false,
	}
	got := support.CheckWebsites(support.MockWebsiteChecker, websites)
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("Wanted %v, got %v", want, got)
	}
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = fmt.Sprintf("a url")
	}
	for i := 0; i < b.N; i++ {
		support.CheckWebsites(support.SlowStubWebsiteChecker, urls)
	}
}
