package main

import (
	"bytes"
	"encoding/json"
	//"github.com/gin-gonic/gin"
	//. "github.com/U-taro-ogw/go_test_sample/auth_api"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("AuthApi", func() {

	type PostParameter struct {
		Email string `json:"email"`
		Password string `json:"password"`
	}

	postParameter := new(PostParameter)

	r := GetMainEngine()
	w := httptest.NewRecorder()
	BeforeEach(func() {
		w = httptest.NewRecorder()
		postParameter = new(PostParameter)
	})

	// /v1/signupへのrequest spec
	Describe("Signup", func() {
		Context("POSTパラメータが存在する場合", func() {

			Context("email password が blank でない場合", func() {
				It("会員登録成功", func() {
					postParameter.Email = "foo@example.com"
					postParameter.Password = "password"
					sampleJson, _ := json.Marshal(postParameter)
					body := bytes.NewBuffer(sampleJson)

					req, _ := http.NewRequest("POST", "v1/signup", body)
					r.ServeHTTP(w, req)
					Expect(w.Code).To(Equal(201))
				})
			})

			Context("email が blank の場合", func() {
				It("404エラーを返す", func() {
					postParameter.Password = "password"
					sampleJson, _ := json.Marshal(postParameter)
					body := bytes.NewBuffer(sampleJson)

					req, _ := http.NewRequest("POST", "v1/signup", body)
					r.ServeHTTP(w, req)
					Expect(w.Code).To(Equal(400))
				})
			})

			Context("password が blank の場合", func() {
				It("404エラーを返す", func() {
					postParameter.Email = "foo@example.com"
					sampleJson, _ := json.Marshal(postParameter)
					body := bytes.NewBuffer(sampleJson)

					req, _ := http.NewRequest("POST", "v1/signup", body)
					r.ServeHTTP(w, req)
					Expect(w.Code).To(Equal(400))
				})
			})
		})

		Context("POSTパラメータが存在しない場合", func() {
			It("400エラーを返す", func() {
				req, _ := http.NewRequest("POST", "v1/signup", nil)
				r.ServeHTTP(w, req)

				Expect(w.Code).To(Equal(400))
			})
		})
	})
})
