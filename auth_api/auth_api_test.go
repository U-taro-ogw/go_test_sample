package main_test

import (
	//. "github.com/U-taro-ogw/go_test_sample/auth_api"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("AuthApi", func() {

	Describe("Signup", func() {
		Context("POSTパラメータが存在する場合", func() {

		})

		Context("POSTパラメータが存在しない場合", func() {
			It("400エラーを返す", func() {
				req, _ := http.NewRequest("POST", "/v1/signup", nil)
				w := httptest.NewRecorder()
				r := gin.Default()
				r.ServeHTTP(w, req)
				Expect(w.Code).To(Equal(404))
			})
		})
	})
})
