package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAuthApi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AuthApi Suite")
}
