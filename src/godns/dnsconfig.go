package godns

import ()

var GoogleDNSServers = []string{"8.8.8.8", "8.8.4.4"}
var OpenDNSServers = []string{"208.67.222.222", "208.67.220.220"}

type dnsConfig struct {
	servers  []string // servers to use
	search   []string // suffixes to append to local name
	ndots    int      // number of dots in name to trigger absolute lookup
	timeout  int      // seconds before giving up on packet
	attempts int      // lost packets before giving up on server
	rotate   bool     // round robin among servers
}

func dnsConfigWithServers(servers []string) (*dnsConfig, error) {
	conf := new(dnsConfig)
	conf.servers = servers
	//conf.servers = make([]string, 3)[0:0] // small, but the standard limit
	conf.search = make([]string, 0)
	conf.ndots = 1
	conf.timeout = 5
	conf.attempts = 2
	conf.rotate = false
	return conf, nil
}
