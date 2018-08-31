package generation_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGeneration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Generation Suite")
}
