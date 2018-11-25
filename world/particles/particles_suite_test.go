package particles_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestParticles(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Particles Suite")
}
