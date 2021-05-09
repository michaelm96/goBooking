package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("got involved when should have been valid")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	// if this is valid something is wrong
	if form.Valid() {
		t.Error("form shows valid when requires fields missing")
	}

	// insert data
	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "a")
	postedData.Add("c", "a")

	// make a request
	r, _ = http.NewRequest("POST", "/whatever", nil)

	// assign data to postForm
	r.PostForm = postedData
	form = New(r.PostForm)

	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("show doesn't have required field when it does")
	}
}

func Test_MinLength(t *testing.T) {
	r, _ := http.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.MinLength("what", 7)
	if form.Valid() {
		t.Error("form shows min length for non existent field")
	}

	postedData := url.Values{}
	postedData.Add("whateverItIs", "hahaha")
	postedData.Add("whateverItIs2", "wkwkwkwkwk")

	form = New(postedData)

	check := form.MinLength("whateverItIs", 8)
	if check {
		t.Error("show true even tho it should be error cuz 'hahaha' is less than 8 chars")
	}

	check = form.MinLength("whateverItIs2", 8)
	if !check {
		t.Error("show false but it should be true cuz 'wkwkwkwkwk' is more than 8 chars")
	}

}
func Test_Has(t *testing.T) {
	r, _ := http.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	has := form.Has("whatever")
	if has {
		t.Error("form shows have data but actually not")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")

	form = New(postedData)
	has = form.Has("a")
	if !has {
		t.Error("form shows no any data, but actually there is")
	}
}

func Test_isEmail(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)

	form.IsEmail("x")
	if form.Valid() {
		t.Error("should be error due to non existent email")
	}
	postedData = url.Values{}
	postedData.Add("email", "abc@email.com")
	form = New(postedData)

	form.IsEmail("email")
	if !form.Valid() {
		t.Error("should be true due to email is really exist")
	}

	postedData.Add("emailErr", "abc")
	form = New(postedData)

	form.IsEmail("emailErr")
	if form.Valid() {
		t.Error("should be error due to input is not email type")
	}
}

func Test_Add(t *testing.T) {
	// var err errors doesn't work
	err := make(errors)

	err.Add("key", "value")

}

func Test_Get(t *testing.T) {
	err := make(errors)

	err.Add("key", "value")
	value := err.Get("key")
	if len(value) == 0 {
		t.Error("the length should be more than zero since the key is exist")
	}

	value = err.Get("Schlüssel")
	if len(value) != 0 {
		t.Error("the length should be zero since the key isn't exist yet")
	}
}
