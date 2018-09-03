package automata_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAutomata(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Automata Suite")
}
