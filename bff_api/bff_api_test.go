package main_test

import (
	. "github.com/U-taro-ogw/go_test_sample/bff_api"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	//"github.com/jarcoal/httpmock"
	//"github.com/onsi/gomega/ghttp"

	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	//"os"
)

var app App

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	app.Router.ServeHTTP(rr, req)
	return rr
}

var _ = Describe("BffApi", func() {

	//var server *ghttp.Server
	//var returnedJson string
	//var statusCode int

	BeforeEach(func() {
		app = App{}
		app.Initialize()

		//server.
		//server = ghttp.NewServer()
		//os.Setenv("http://flask_api_one:5000", server.Addr())
		//server.AppendHandlers(
		//	ghttp.CombineHandlers(
		//		ghttp.VerifyRequest("GET", "/api_info"),
		//		ghttp.RespondWithJSONEncoded(http.StatusOK, `{"api_name": "flask_api_one", "info": {"framework": "flask", "language": "python"}}`),
		//	),
		//)
		//httpmock.Activate()
	})

	//AfterEach(func() {
		//server.Close()
	//	httpmock.DeactivateAndReset()
	//})

	Describe("Apis", func() {
		Context("Request Headerにjwt tokenが存在する場合", func() {
			Context("Redisに認証情報が存在する場合", func() {
				// 他APIへRequestできる状態
				Context("全てのAPIから200が返却された場合", func() {
					//BeforeEach(func() {
						//returnedJson = `{"api_name": "flask_api_one", "info": {"framework": "flask", "language": "python"}}`
						//statusCode = 200
					//})
					//httpmock.RegisterResponder("GET", "http://flask_api_one:5000/api_info",
					//	httpmock.NewStringResponder(200, `{"api_name": "flask_api_one", "info": {"framework": "flask", "language": "python"}}`))

					It("200を返却する", func() {
						server := httptest.NewServer(http.HandlerFunc(func(rs http.ResponseWriter, rq *http.Request) {
							jsonStr := `{"api_name": "flask_api_one", "info": {"framework": "flask", "language": "python"}}`
							//rs.Write()
							rs.Write(([]byte)(jsonStr))
							//if rq.RequestURI == `http://flask_api_one:5000/api_info` {}
						}))
						defer server.Close()

						req, _ := http.NewRequest("GET", "/v1/apis", nil)
						response := executeRequest(req)

						fmt.Println("response.Code")
						fmt.Println(response.Code)

						Expect(response.Code).To(Equal(http.StatusOK))
					})

					It("API情報を返却する", func() {
						req, _ := http.NewRequest("GET", "/v1/apis", nil)
						response := executeRequest(req)

						res := map[string]string{"key":"value"}
						bytes, _ := json.Marshal(res)

						fmt.Println("response.Body.String()")
						fmt.Println(response.Body.String())

						Expect(response.Body.String()).To(Equal(string(bytes)))
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
