package neighborhood_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestNeighborhood(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Neighborhood Suite")
}
