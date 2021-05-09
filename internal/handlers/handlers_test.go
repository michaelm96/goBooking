package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type PostData struct {
	key   string
	value string
}

var theTests = []struct {
	name               string
	url                string
	method             string
	params             []PostData
	expectedStatusCode int
}{
	{
		"home",
		"/",
		"GET",
		[]PostData{},
		http.StatusOK,
	},
	{
		"about",
		"/about",
		"GET",
		[]PostData{},
		http.StatusOK,
	},
	{
		"generals-quarters",
		"/generals-quarters",
		"GET",
		[]PostData{},
		http.StatusOK,
	},
	{
		"majors-suite",
		"/majors-suite",
		"GET",
		[]PostData{},
		http.StatusOK,
	},
	{
		"search-availability",
		"/search-availability",
		"GET",
		[]PostData{},
		http.StatusOK,
	},
	{
		"contact",
		"/contact",
		"GET",
		[]PostData{},
		http.StatusOK,
	},
	{
		"make-reservation",
		"/make-reservation",
		"GET",
		[]PostData{},
		http.StatusOK,
	},
	{
		"search-availability post",
		"/search-availability",
		"POST",
		[]PostData{
			{key: "start", value: "2020-01-01"},
			{key: "end", value: "2020-01-02"},
		},
		http.StatusOK,
	},
	{
		"search-availability-json",
		"/search-availability-json",
		"POST",
		[]PostData{
			{key: "start", value: "2020-01-01"},
			{key: "end", value: "2020-01-02"},
		},
		http.StatusOK,
	},
	{
		"make-reservation post",
		"/make-reservation",
		"POST",
		[]PostData{
			{key: "first_name", value: "John"},
			{key: "last_name", value: "Doe"},
			{key: "email", value: "abc@email.com"},
			{key: "phone", value: "123-456-7890"},
		},
		http.StatusOK,
	},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	server := httptest.NewTLSServer(routes)
	defer server.Close()

	for _, tt := range theTests {
		// t.Run(tt.name, func(t *testing.T) {
		// 	if got := routes.GET(); !reflect.DeepEqual(got, tt.expectedStatusCode) {
		// 		t.Errorf("routes() = %v, want %v", got, tt.expectedStatusCode)
		// 	}
		// })
		if tt.method == "GET" {
			resp, err := server.Client().Get(server.URL + tt.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}
			if resp.StatusCode != tt.expectedStatusCode {
				t.Errorf("got statusCode = %v, want %v", resp.StatusCode, tt.expectedStatusCode)
			}

			t.Logf("got status code %v, for route %v", resp.StatusCode, tt.name)
		} else {
			values := url.Values{}
			for _, obj := range tt.params {
				values.Add(obj.key, obj.value)
			}

			resp, err := server.Client().PostForm(server.URL+tt.url, values)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if resp.StatusCode != tt.expectedStatusCode {
				t.Errorf("got statusCode = %v, want %v", resp.StatusCode, tt.expectedStatusCode)
			}
		}
	}
}
