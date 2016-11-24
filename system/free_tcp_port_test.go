package system_test

import (
        "net"
        "regexp"
        "strconv"

        "github.com/pivotal-cf/cf-redis-broker/system"

        . "github.com/onsi/ginkgo"
        . "github.com/onsi/gomega"
)

var _ = Describe("Next available TCP port", func() {

        It("finds a the  free TCP port in the range ", func() {
                port, _ := system.FindFreeInRangePort(40005,40000)
                portStr := strconv.Itoa(port)

                matched, err := regexp.MatchString("^[0-9]+$", portStr)
                Ω(matched).To(Equal(true))

                l, err := net.Listen("tcp", ":"+portStr)
                Ω(err).ToNot(HaveOccurred())
                l.Close()
        })
        It("test the case when no port available in the range ", func() {
                _, err := system.FindFreeInRangePort(-1,-1)
                Ω(err).To(HaveOccurred())
                Ω(err).err.String().Should(ContainSubstring("No Free port"))                
        })

})

