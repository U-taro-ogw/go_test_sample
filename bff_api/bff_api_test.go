package main_test

import (
	. "github.com/U-taro-ogw/go_test_sample/bff_api"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/onsi/gomega/ghttp"

	"fmt"
	"net/http"
	"net/http/httptest"
)

var app App

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	app.Router.ServeHTTP(rr, req)
	return rr
}

var _ = Describe("BffApi", func() {

	var (
		server *ghttp.Server
		statusCode int
		body []byte
		path string
		addr string
	)

	BeforeEach(func() {
		app = App{}
		app.Initialize()
		server = ghttp.NewServer()
	})

	AfterEach(func() {
		server.Close()
	})

	Describe("Apis", func() {
		Context("Request Headerにjwt tokenが存在する場合", func() {
			Context("Redisに認証情報が存在する場合", func() {
				// 他APIへRequestできる状態
				Context("全てのAPIから200が返却された場合", func() {
					BeforeEach(func() {
						statusCode = 200
						path = "/api_info"
						body = []byte(`{"api_name": "flask_api_one", "info": {"framework": "flask", "language": "python"}}`)
						//addr = "http://flask_api_one:5000" + path
						addr = "http://" + server.Addr() + path
						server.AppendHandlers(
							ghttp.CombineHandlers(
								ghttp.VerifyRequest("GET", path),
								ghttp.RespondWithPtr(&statusCode, &body),
							),
						)

					})

					It("200を返却する", func() {
						// 下記であれば、ghttpで作成したresponseを返却することができる
						//jwt := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.6_Y1pbXBLXRNFTDjs1hIKlesJ9oRjvg1u5OCAkMTqkI"
						//c, b, e := GetResponse(addr, jwt)

						// rspecのrequest specを想定したテストの書き方をしていたが
						// こちらだと addrを渡す術がないためできない。
						req, _ := http.NewRequest("GET", "/v1/apis", addr)
						response := executeRequest(req)

						Expect(response.Code).To(Equal(http.StatusOK))
					})

					It("API情報を返却する", func() {
						//req, _ := http.NewRequest("GET", "/v1/apis", nil)
						//response := executeRequest(req)

						//res := map[string]string{"key":"value"}
						//bytes, _ := json.Marshal(res)

						//fmt.Println("response.Body.String()")
						//fmt.Println(response.Body.String())

						//Expect(response.Body.String()).To(Equal(string(bytes)))
					})
				})

				Context("一部のAPIから200以外が返却された場合", func() {
					Context("401の場合", func() {
						//It("401を返却する", func() {})
					})

					Context("403の場合", func() {
						//It("403を返却する", func() {})
					})

					Context("500の場合", func() {
						//It("500を返却する", func() {})
					})
				})

				Context("全てのAPIから200以外が返却された場合", func() {
					Context("401の場合", func() {
						//It("401を返却する", func() {})
					})

					Context("403の場合", func() {
						//It("403を返却する", func() {})
					})

					Context("500の場合", func() {
						//It("500を返却する", func() {})
					})
				})
			})

			Context("Redisに認証情報が存在しない場合", func() {
				// 401を返却する
				//It("401を返却する", func() {})
			})
		})

		Context("Request Headerにjwt tokenが存在しない場合", func() {
			// 401を返却する
			//It("401を返却する", func() {})
		})
	})
})
