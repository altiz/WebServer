package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type data struct {
	User string `json:"user"`
}
type data1 struct {
	Status int
}

func TestMain(t *testing.T) {
	router := setupRouter()

	Convey("Users endpoints should respond correctly", t, func() {
		Convey("User should return empty list", func() {
			// it's safe to ignore error here, because we're manually entering URL
			req, _ := http.NewRequest("GET", "http://localhost:8080/api/v1/", nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			So(w.Code, ShouldEqual, http.StatusOK)
			body := strings.TrimSpace(w.Body.String())
			So(body, ShouldEqual, "ok")
		})

		Convey("Create should return ID of newly created user", func() {
			//send message
			user := &data{User: "test"}
			dat, err := json.Marshal(user)
			So(err, ShouldBeNil)
			buf := bytes.NewBuffer(dat)
			req, err := http.NewRequest("POST", "http://localhost/api/v1/test", buf)
			So(err, ShouldBeNil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			So(w.Code, ShouldEqual, http.StatusOK)
			//resp
			body := w.Body.Bytes()
			var obj *data1

			if err := json.Unmarshal(body, &obj); err != nil {
				panic(err)
			}
			So(obj.Status, ShouldEqual, 0)
		})
	})
}
