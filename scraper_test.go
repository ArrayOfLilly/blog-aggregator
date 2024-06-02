package main

import (
	"testing"
	"reflect"
)

func testFetchFeed(t *testing.T) {
	input := "https://blog.boot.dev/index.xml"
	expected, err := fetchFeed(input)
	if err != nil {
		t.Errorf("Error: %v", err.Error())
	}
	if reflect.TypeOf(expected) != reflect.TypeOf(RSSFeed{}) {
		t.Errorf("Error: not a valid RSS feed")
	}
}