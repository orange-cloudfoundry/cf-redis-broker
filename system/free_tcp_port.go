package system

import (
	"net"
	"strconv"
	"errors"
)

const (
	MIN_ACCEPTED_PORT int = 1024
	MAX_ACCEPTED_PORT int = 65535
)

type FreeTcpPort interface {
	FindFreePortInRange(minport int, maxport int) (int, error)
}
type FreeRangeTcpPort struct {
	FreeTcpPort
	IsPortAvailable func(num int) bool
}

func NewFreeTcpPort() FreeTcpPort {
	return &FreeRangeTcpPort{IsPortAvailable: isPortAvailable}
}
func isPortAvailable(num int) bool {
	l, err := net.Listen("tcp", ":" + strconv.Itoa(num))
	if err != nil {
		return false
	}
	l.Close()
	return true
}

func (f FreeRangeTcpPort) FindFreePortInRange(minport int, maxport int) (int, error) {
	if minport > maxport {
		return -1, errors.New("Not valid range port: minimum port is higher than maximum port")
	}
	if minport < MIN_ACCEPTED_PORT {
		return -1, errors.New("Not valid range port: minimum port is lower than " + strconv.Itoa(MIN_ACCEPTED_PORT))
	}
	if maxport > MAX_ACCEPTED_PORT {
		return -1, errors.New("Not valid range port: maximum port is higher than " + strconv.Itoa(MIN_ACCEPTED_PORT))
	}
	port := minport
	for port < maxport {
		if f.IsPortAvailable(port) {
			return port, nil
		}
		port++
	}
	return -1, errors.New("Sorry no port is available in this range")
}
