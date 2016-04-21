package ewa_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestEwa(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ewa Suite")
}
