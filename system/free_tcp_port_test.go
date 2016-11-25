package system_test

import (
	. "github.com/pivotal-cf/cf-redis-broker/system"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Find next available TCP port in a range", func() {
	var freeTcpPort FreeTcpPort
	BeforeEach(func() {
		freeTcpPort = NewFreeTcpPort()
	})
	Context("with a valid range", func() {
		It("should give a free port", func() {
			port, err := freeTcpPort.FindFreePortInRange(40000, 40005)
			Expect(err).ShouldNot(HaveOccurred())

			Expect(err).ShouldNot(HaveOccurred())
			Expect(port).To(BeNumerically(">=", 40000))
			Expect(port).To(BeNumerically("<=", 40005))
		})
		It("should give the only free port if the the minimum port is equals to maximum port", func() {
			freeTcpPort = NewFreeTcpPort()
			freeTcpPort.(*FreeRangeTcpPort).IsPortAvailable = func(num int) bool {
				return true
			}

			port, err := freeTcpPort.FindFreePortInRange(65000, 65000)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(err).ShouldNot(HaveOccurred())

			Expect(port).To(BeEquivalentTo(65000))
		})
	})
	Context("with an invalid range port", func() {
		It("should return an error if minimum port is greater than maximum port", func() {
			_, err := freeTcpPort.FindFreePortInRange(6000, 5000)
			Expect(err).Should(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("minimum port is higher than maximum port"))
		})
		It("should return an error if minimum port is lower than accepted minimum port", func() {
			_, err := freeTcpPort.FindFreePortInRange(MIN_ACCEPTED_PORT - 1, 5000)
			Expect(err).Should(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("minimum port is lower than"))
		})
		It("should return an error if maximum port is lower than accepted maximum port", func() {
			_, err := freeTcpPort.FindFreePortInRange(4000, MAX_ACCEPTED_PORT + 1)
			Expect(err).Should(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("maximum port is higher than"))
		})
	})
	Context("When no port is available in range", func() {
		BeforeEach(func() {
			freeTcpPort = NewFreeTcpPort()
			freeTcpPort.(*FreeRangeTcpPort).IsPortAvailable = func(num int) bool {
				return false
			}
		})
		It("Should return an error", func() {
			_, err := freeTcpPort.FindFreePortInRange(40000, 40005)
			Expect(err).Should(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("No port is available in the range"))
		})
	})

})

