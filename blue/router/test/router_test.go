package router_test

import (
	"fmt"
	"play-go/blue/router"
	"reflect"
	"testing"
)

func newTestRouter() *router.Router {
	r := router.NewRouter()
	r.AddRoute("GET", "/", nil)
	r.AddRoute("GET", "/hello/:name", nil)
	r.AddRoute("GET", "/hello/b/c", nil)
	r.AddRoute("GET", "/hi/:name", nil)
	r.AddRoute("GET", "/assets/*filepath", nil)
	return r
}

func TestParsePattern(t *testing.T) {
	ok := reflect.DeepEqual(router.ParsePattern("/p/:name"), []string{"p", ":name"})
	ok = ok && reflect.DeepEqual(router.ParsePattern("/p/*"), []string{"p", "*"})
	ok = ok && reflect.DeepEqual(router.ParsePattern("/p/*name/*"), []string{"p", "*name"})

	if !ok {
		t.Fatal("test ParsePattern failed")
	}
}

func TestGetRoute(t *testing.T) {
	r := newTestRouter()
	n, params := r.GetRoute("GET", "/hello/yang")

	if n == nil {
		t.Fatal("nil shouldn't be returned")
	}

	if n.Parttern != "/hello/:name" {
		t.Fatal("should match /hello/:name")
	}

	if params["name"] != "yang" {
		t.Fatal("name should be equal to 'geektutu'")
	}

	fmt.Printf("matched path: %s, params['name']: %s\n", n.Parttern, params["name"])
}
