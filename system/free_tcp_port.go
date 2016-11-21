package system

import (
        "net"
        "strconv"
        "errors"
)

func FindFreePort(num int) (int, error) {
        t := strconv.Itoa(num)
        toto := ":" + t
        l, err := net.Listen("tcp", toto)
        if err != nil {
                return -1, err
        }
        parsedPort, _ := strconv.ParseInt(l.Addr().String()[5:], 10, 32)
        return int(parsedPort), nil
}


func FindFreeInRangePort(Maxport int, Minport int)(int,error) {
        i := Minport
        port, err :=FindFreePort(i)
        for (err != nil) && (i<Maxport) {
        i=i+1
        port, err =FindFreePort(i)
        }

        if (err != nil) {
        return -1, errors.New("Sorry No Free port in this range  is availble")
        }
        return port, err
}
