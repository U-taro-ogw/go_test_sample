package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	authDb "github.com/U-taro-ogw/go_test_sample/auth_api/db/mysql"
	"github.com/U-taro-ogw/go_test_sample/auth_api/handlers"
	"github.com/U-taro-ogw/go_test_sample/auth_api/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("AuthApi", func() {

	dbCon := authDb.MysqlConnect()
	postParameter := new(models.User)
	userHandler := handlers.UserHandler{Db: dbCon}

	r := GetMainEngine(userHandler)
	w := httptest.NewRecorder()
	BeforeEach(func() {
		w = httptest.NewRecorder()
		postParameter = new(models.User)
	})

	AfterEach(func() {
		defer dbCon.Close()
	})

	// /v1/signupへのrequest spec
	Describe("Signup", func() {
		Context("POSTパラメータが存在する場合", func() {

			Context("email password が blank でない場合", func() {
				BeforeEach(func() {
					postParameter.Email = "foo@example.com"
					postParameter.Password = "password"
				})

				It("会員登録する", func() {
					var beforeCount = 0
					dbCon.Model(&models.User{}).Count(&beforeCount)

					sampleJson, _ := json.Marshal(postParameter)
					body := bytes.NewBuffer(sampleJson)
					req, _ := http.NewRequest("POST", "v1/signup", body)

					r.ServeHTTP(w, req)

					var afterCount = 0
					dbCon.Model(&models.User{}).Count(&afterCount)

					Expect(afterCount).To(Equal(beforeCount + 1))
				})

				It("200を返却する", func() {
					sampleJson, _ := json.Marshal(postParameter)
					body := bytes.NewBuffer(sampleJson)

					req, _ := http.NewRequest("POST", "v1/signup", body)
					r.ServeHTTP(w, req)
					Expect(w.Code).To(Equal(201))
				})
			})

			Context("email が blank の場合", func() {
				It("404エラーを返却する", func() {
					postParameter.Password = "password"
					sampleJson, _ := json.Marshal(postParameter)
					body := bytes.NewBuffer(sampleJson)

					req, _ := http.NewRequest("POST", "v1/signup", body)
					r.ServeHTTP(w, req)
					Expect(w.Code).To(Equal(400))
				})
			})

			Context("password が blank の場合", func() {
				It("404エラー返却する", func() {
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
			It("400エラーを返却する", func() {
				req, _ := http.NewRequest("POST", "v1/signup", nil)
				r.ServeHTTP(w, req)

				Expect(w.Code).To(Equal(400))
			})
		})
	})
})
