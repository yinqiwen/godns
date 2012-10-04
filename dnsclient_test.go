package godns

import (
	//"net"
	"testing"
)

func TestLoopkupIPOverTcp(t *testing.T) {
	options := &LookupOptions{
		DNSServers: OpenDNSServers, Net: "tcp"}
	addrs, err := LookupIP("www.youtube.com", options)
	if nil != err {
		t.Error("Failed to load ini file for reason:" + err.Error())
		return
	}
	for _, ip := range addrs {
		t.Logf("%v\n", ip.String())
	}
	t.Fail()
}

func TestLoopkupIPOverUdp(t *testing.T) {
	options := &LookupOptions{
		DNSServers: OpenDNSServers, Net: "udp"}
	addrs, err := LookupIP("www.youtube.com", options)
	if nil != err {
		t.Error("Failed to load ini file for reason:" + err.Error())
		return
	}
	for _, ip := range addrs {
		t.Logf("%v\n", ip.String())
	}
	t.Fail()
}
