package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBffApi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BffApi Suite")
}
