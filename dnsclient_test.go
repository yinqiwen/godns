package godns

import (
	//"net"
	"testing"
)

func TestLoopkupIP(t *testing.T) {
	options := &LookupOptions{
		DNSServers: GoogleDNSServers}
	addrs, err := LookupIP("youtube.com", options)
	if nil != err {
		t.Error("Failed to load ini file for reason:" + err.Error())
		return
	}
	for _, ip := range addrs {
		t.Logf("%v\n", ip.String())
	}
	t.Fail()
}
