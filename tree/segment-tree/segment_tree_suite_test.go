package segmenttree_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSegmentTree(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SegmentTree Suite")
}
