package main

import (
	"net/http"
	"os"
	"reflect"
	"testing"
)

func TestSave(t *testing.T) {
	p := Page{"test", []byte("simple test case")}
	if _, err := os.Stat("./data/" + p.Title + ".txt"); os.IsNotExist(err) {
		t.Error(
			"Expected file: ", p.Title+".txt",
			"but none created in this dir",
		)
	}
}

func TestLoadPage(t *testing.T) {
	p, err := loadPage("test")
	if err != nil {
		t.Failed()
	}
	r := &Page{"test", []byte("Hi! Show some edits")}

	for i, v := range r.Body {
		if v != p.Body[i] {
			t.Error(
				"\n Expected: ", v,
				"\n Got: ", p.Body[i],
			)
		}
	}

	if p.Title != r.Title {
		t.Errorf(
			"\n Expected page title: %s. Got title: %s \n", r.Title, p.Title)
	}
}

// func TestGetTitle(t *testing.T) {
// 	req, err := http.NewRequest("GET", "/view/test", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	rr := httptest.NewRecorder()
// 	handler := http.HandlerFunc(makeHandler(viewHandler))
// 	handler.ServeHTTP(rr, req)
// 	if status := rr.Code; status != http.StatusOK {
// 		t.Errorf("handler returned wrong status code: got %v want %v",
// 			status, http.StatusOK)
// 	}

// expected := `{"alive": true}`
// if rr.Body.String() != expected {
// 	t.Errorf("handler returned unexpected body: got %v want %v",
// 		rr.Body.String(), expected)
// }
// }

var handlers = []func(http.ResponseWriter, *http.Request, string){
	viewHandler,
	editHandler,
	saveHandler,
}

func TestMakeHandler(t *testing.T) {
	for _, f := range handlers {
		h := makeHandler(f)
		r := reflect.TypeOf(h).String()
		if r != "http.HandlerFunc" {
			t.Error(
				"\n Expected http.HandlerFunc",
				"\n Got: ", r,
			)
		}
	}
	// fmt.Println(reflect.TypeOf(viewHandler).String())
}
