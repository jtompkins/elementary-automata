package evaluators_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestEvaluators(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Evaluators Suite")
}
