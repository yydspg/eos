
package network

import (
	"errors"
	"net"
	"net/http"
	"strings"

	_ "github.com/eos/tools/errs"
)

// Define http headers.
const (
	XForwardedFor = "X-Forwarded-For"
	XRealIP       = "X-Real-IP"
	XClientIP     = "x-client-ip"
)

func GetLocalIP() (string, error) {
	// Fetch all network interfaces
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	// Iterate over each interface
	var publicIP string
	for _, iface := range interfaces {
		// Check if the interface is up and not a loopback
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
			continue
		}

		// Get all addresses associated with the interface
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}

		// Check each address for a valid IPv4 address that is not a loopback
		for _, addr := range addrs {
			// Try to parse the address as an IPNet (CIDR notation)
			ipNet, ok := addr.(*net.IPNet)
			if !ok || ipNet.IP.IsLoopback() {
				continue
			}

			ip4 := ipNet.IP.To4()
			if ip4 != nil && !ip4.IsLoopback() {
				// Ensure the IP is not a multicast address
				if !ip4.IsMulticast() {
					if !ipNet.IP.IsPrivate() && publicIP == "" {
						// Priority return to internal network IP
						publicIP = ipNet.IP.String()
					} else {
						return ip4.String(), nil
					}

				}
			}
		}
	}

	if publicIP != "" {
		return publicIP, nil
	}
	// If no suitable IP is found, return an error
	return "", errors.New("no suitable local IP address found")
}

func GetRpcRegisterIP(configIP string) (string, error) {
	registerIP := configIP
	if registerIP == "" {
		ip, err := GetLocalIP()
		if err != nil {
			return "", err
		}
		registerIP = ip
	}
	return registerIP, nil
}

func GetListenIP(configIP string) string {
	if configIP == "" {
		return "0.0.0.0"
	}
	return configIP
}

// RemoteIP returns the remote ip of the request.
func RemoteIP(req *http.Request) string {
	if ip := req.Header.Get(XClientIP); ip != "" {
		return ip
	} else if ip := req.Header.Get(XRealIP); ip != "" {
		return ip
	} else if ip := req.Header.Get(XForwardedFor); ip != "" {
		parts := strings.Split(ip, ",")
		return strings.TrimSpace(parts[0])
	}

	ip, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		ip = req.RemoteAddr
	}

	if ip == "::1" {
		return "127.0.0.1"
	}
	return ip
}
