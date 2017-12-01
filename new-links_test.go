package htmlPage

import (
	"testing"
	hash "github.com/rogeecn/gohash"
)

func Test_NewLinks(t *testing.T) {
	testLinks := []string{"1", "2", "3", "4", "5"};
	lastId := "c4ca4238a0b923820dcc509a6f75849b" //1

	t.Log("test db last is first link....")
	links := NewLinks(testLinks, lastId, hash.MD5, ORDER_HEAD)
	t.Logf("HEAD: %+v", links)

	t.Log("--------------------------------------")
	links = NewLinks(testLinks, lastId, hash.MD5, ORDER_TAIL)
	t.Logf("TAIL: %+v\n\n\n", links)

	t.Log("--------------------------------------")
	t.Log("test db last is middle link....")
	lastId = "c81e728d9d4c2f636f067f89cc14862c" //2
	links = NewLinks(testLinks, lastId, hash.MD5, ORDER_HEAD)
	t.Logf("HEAD: %+v", links)

	t.Log("--------------------------------------")
	links = NewLinks(testLinks, lastId, hash.MD5, ORDER_TAIL)
	t.Logf("TAIL: %+v\n\n\n", links)

	t.Log("--------------------------------------")
	t.Log("test db last is last link....")
	lastId = "e4da3b7fbbce2345d7772b0674a318d5" //5
	links = NewLinks(testLinks, lastId, hash.MD5, ORDER_HEAD)
	t.Logf("HEAD: %+v", links)
	t.Log("--------------------------------------")
	links = NewLinks(testLinks, lastId, hash.MD5, ORDER_TAIL)
	t.Logf("TAIL: %+v\n\n\n", links)

	t.Log("--------------------------------------")
	t.Log("test db last is no link....")
	lastId = "123e4da3b7fbbce2345d7772b0674a318d5" //none
	links = NewLinks(testLinks, lastId, hash.MD5, ORDER_HEAD)
	t.Logf("HEAD: %+v", links)

	t.Log("--------------------------------------")
	links = NewLinks(testLinks, lastId, hash.MD5, ORDER_TAIL)
	t.Logf("TAIL: %+v", links)
	t.Log("done")
}
