package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	authDb "github.com/U-taro-ogw/go_test_sample/auth_api/db/mysql"
	redisConnect "github.com/U-taro-ogw/go_test_sample/auth_api/db/redis"
	"github.com/U-taro-ogw/go_test_sample/auth_api/handlers"
	"github.com/U-taro-ogw/go_test_sample/auth_api/models"
	"github.com/U-taro-ogw/go_test_sample/auth_api/modules"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	"reflect"
)


//type JwtTokenMock struct {}
//func (j JwtTokenMock) getEmail() string {
//	return "hoge"
//}

var _ = Describe("AuthApi", func() {

	dbCon := authDb.MysqlConnect()
	postParameter := new(models.User)
	userHandler := handlers.UserHandler{Db: dbCon}

	fmt.Println("ginkgo start")
	fmt.Println(reflect.TypeOf("ginkgo start"))
	r := GetMainEngine(userHandler)
	w := httptest.NewRecorder()
	BeforeEach(func() {
		w = httptest.NewRecorder()
		postParameter = new(models.User)
	})

	AfterEach(func() {
		// userテーブル全削除してしまうためテスト実行時にテスト用DBに切り替えること
		dbCon.Exec("DELETE FROM users")
		//defer dbCon.Close()
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

	// /v1/signinへのrequest spec
	Describe("Signin", func() {
		Context("POSTパラメータが存在する場合", func() {
			testUserEmail := "foo@example.com"
			testUserPassword := "password"
			jwtTokenStr := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.ez9OosQLyqzPvW6MoL6UzxmSrQ5BfSo8nvXbNSJdevU"
			BeforeEach(func() {
				testUser := models.User{}
				testUser.Email = testUserEmail
				testUser.Password = testUserPassword
				dbCon.Create(&testUser)
			})

			Context("パラメータ通りのuserが存在する場合", func() {
				BeforeEach(func() {
					postParameter.Email = testUserEmail
					postParameter.Password = testUserPassword
				})

				It("200を返却する", func() {
					sampleJson, _ := json.Marshal(postParameter)
					body := bytes.NewBuffer(sampleJson)

					req, _ := http.NewRequest("POST", "v1/signin", body)
					r.ServeHTTP(w, req)
					Expect(w.Code).To(Equal(200))
				})

				It("jwt tokenを返却する", func() {
					sampleJson, _ := json.Marshal(postParameter)
					body := bytes.NewBuffer(sampleJson)

					req, _ := http.NewRequest("POST", "v1/signin", body)
					r.ServeHTTP(w, req)

					m := make(map[string]string)
					json.Unmarshal(w.Body.Bytes(), &m)

					//mock := JwtTokenMock{}
					//hoge := modules.GetJwtToken(mock)
					//fmt.Println(hoge)

					// TODO 非常によろしくない。
					// jwt.New().SignedString([]byte("hoge"))
					// の部分をmock化して都合の良い文字列を返すようにしたい
					// => この考え方でコード書かないほうが良い可能性もアル
					Expect(m["jwt_token"]).To(Equal(jwtTokenStr))
				})

				It("jwt tokenを保存する", func() {
					// TODO redis保存するmoduleをmock化したい
					sampleJson, _ := json.Marshal(postParameter)
					body := bytes.NewBuffer(sampleJson)

					req, _ := http.NewRequest("POST", "v1/signin", body)
					r.ServeHTTP(w, req)

					r := redisConnect.RedisConnect()
					val := modules.RedisGet(r, jwtTokenStr)

					Expect(val).To(Equal("111"))
				})
			})

			Context("パラメータ通りのuserが存在しない場合", func() {
				Context("emailが違う場合", func() {
					BeforeEach(func() {
						postParameter.Email = testUserEmail + "a"
						postParameter.Password = testUserPassword
					})
					It("401エラーを返却する", func() {
						sampleJson, _ := json.Marshal(postParameter)
						body := bytes.NewBuffer(sampleJson)

						req, _ := http.NewRequest("POST", "v1/signin", body)
						r.ServeHTTP(w, req)
						Expect(w.Code).To(Equal(401))
					})
				})

				Context("passwordが違う場合", func() {
					BeforeEach(func() {
						postParameter.Email = testUserEmail
						postParameter.Password = testUserPassword + "a"
					})
					It("401エラーを返却する", func() {
						sampleJson, _ := json.Marshal(postParameter)
						body := bytes.NewBuffer(sampleJson)

						req, _ := http.NewRequest("POST", "v1/signin", body)
						r.ServeHTTP(w, req)
						Expect(w.Code).To(Equal(401))
					})
				})

				Context("email password 両方が違う場合", func() {
					BeforeEach(func() {
						postParameter.Email = testUserEmail + "a"
						postParameter.Password = testUserPassword + "a"
					})
					It("401エラーを返却する", func() {
						sampleJson, _ := json.Marshal(postParameter)
						body := bytes.NewBuffer(sampleJson)

						req, _ := http.NewRequest("POST", "v1/signin", body)
						r.ServeHTTP(w, req)
						Expect(w.Code).To(Equal(401))
					})
				})
			})
		})

		Context("POSTパラメータが存在しない場合", func() {
			It("400エラーを返却する", func() {
				req, _ := http.NewRequest("POST", "v1/signin", nil)
				r.ServeHTTP(w, req)

				Expect(w.Code).To(Equal(400))
			})
		})
	})
})
