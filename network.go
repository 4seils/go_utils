package go_utils

import (
	"errors"
	"net"
)

// Get preferred outbound ip of this machine
func GetOutboundIP() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String(), nil
}

func GetMacAddress(ip string) (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, i := range ifaces {
		mac := i.HardwareAddr.String()
		addrs, err := i.Addrs()
		if err != nil {
			continue
		}
		for _, a := range addrs {
			switch v := a.(type) {
			case *net.IPNet:
				if ip == v.IP.String() {
					return mac, nil
				}
			}

		}
	}
	return "", errors.New("No matching mac with IP:" + ip)
}

/*
func main() {
	ip, err := GetOutboundIP()
	if err == nil {
		GetMacAddress(ip)
	}
	fmt.Println(getMacAddr())
}
*/
