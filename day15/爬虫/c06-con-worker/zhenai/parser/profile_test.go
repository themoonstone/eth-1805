package parser

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParseUser(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")


	//contents, err := fetcher.Fetch("http://album.zhenai.com/u/1358992404")
	if nil != err {
		panic(err)
	}
	results := ParseUser(contents, "白雪王子")
	fmt.Printf("length : %v\n", len(results.Items))
	if len(results.Items ) != 1 {
		t.Errorf("Items should contain 1" +
			"element; but was %v", results.Items)
	}
	fmt.Printf("%v\n", results.Items)

}
