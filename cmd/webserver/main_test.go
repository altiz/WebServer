package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type data struct {
	User string `json:"user"`
}

func TestMain(t *testing.T) {
	router := setupRouter()

	Convey("Users endpoints should respond correctly", t, func() {
		Convey("User should return empty list", func() {
			// it's safe to ignore error here, because we're manually entering URL
			req, _ := http.NewRequest("GET", "http://localhost:8080/v1/", nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			So(w.Code, ShouldEqual, http.StatusOK)
			body := strings.TrimSpace(w.Body.String())
			So(body, ShouldEqual, "ok")
		})

		Convey("Create should return ID of newly created user", func() {
			user := &data{User: "Test user"}
			dat, err := json.Marshal(user)
			So(err, ShouldBeNil)
			buf := bytes.NewBuffer(dat)
			req, err := http.NewRequest("POST", "http://localhost/v1/test", buf)
			So(err, ShouldBeNil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			So(w.Code, ShouldEqual, http.StatusOK)

			body := w.Body.Bytes()
			var users []*data
			json.Unmarshal(body, &user)
			fmt.Println(user)
			So(len(users), ShouldEqual, 1)
		})
	})
}
