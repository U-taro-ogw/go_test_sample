package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	//. "github.com/U-taro-ogw/go_test_sample/bff_api"
)

var _ = Describe("BffApi", func() {

	Describe("Apis", func() {
		Context("Request Headerにjwt tokenが存在する場合", func() {
			Context("Redisに認証情報が存在する場合", func() {
				// 他APIへRequestできる状態
				Context("全てのAPIから200が返却された場合", func() {
					It("200を返却する", func() {

					})

					It("API情報を返却する", func() {

					})
				})

				Context("一部のAPIから200以外が返却された場合", func() {
					Context("401の場合", func() {
						It("401を返却する", func() {})
					})

					Context("403の場合", func() {
						It("403を返却する", func() {})
					})

					Context("500の場合", func() {
						It("500を返却する", func() {})
					})
				})

				Context("全てのAPIから200以外が返却された場合", func() {
					Context("401の場合", func() {
						It("401を返却する", func() {})
					})

					Context("403の場合", func() {
						It("403を返却する", func() {})
					})

					Context("500の場合", func() {
						It("500を返却する", func() {})
					})
				})
			})

			Context("Redisに認証情報が存在しない場合", func() {
				// 401を返却する
				It("401を返却する", func() {})
			})
		})

		Context("Request Headerにjwt tokenが存在しない場合", func() {
			// 401を返却する
			It("401を返却する", func() {})
		})
	})
})
