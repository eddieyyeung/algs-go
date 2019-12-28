package sorting_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSort(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Sorting Suite")
}
