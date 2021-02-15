package main

import (
	"reflect"
	"testing"
	"time"
)

func slowStubVerificatorWebsite(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkVerifyWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "one url"
	}

	for i := 0; i < b.N; i++ {
		VerifyWebsites(slowStubVerificatorWebsite, urls)
	}
}

func mockVerificatorWebsite(url string) bool {
	if url == "waat://www.furhurterwe.geds" {
		return false
	}

	return true
}

func TestVerifyWebsites(t *testing.T) {
	websites := []string{
		"https://www.google.com.br",
		"https://www.blog.gypsydave5.com",
		"waat://www.furhurterwe.geds",
	}

	expect := map[string]bool{
		"https://www.google.com.br":       true,
		"https://www.blog.gypsydave5.com": true,
		"waat://www.furhurterwe.geds":     false,
	}

	result := VerifyWebsites(mockVerificatorWebsite, websites)

	if !reflect.DeepEqual(expect, result) {
		t.Fatalf("expect %v, received %v", expect, result)
	}
}
