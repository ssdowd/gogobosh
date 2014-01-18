package main

import (
	"github.com/cloudfoundry-community/gogobosh"
	"fmt"
)

func main() {
	director := gogobosh.NewDirector("https://192.168.50.4:25555", "admin", "admin")
	info := director.GetInfo()
	fmt.Println("Director")
	fmt.Printf("  Name       %s\n", info.Name)
	fmt.Printf("  URL        %s\n", info.URL)
	fmt.Printf("  Version    %s\n", info.Version)
	fmt.Printf("  User       %s\n", info.User)
	fmt.Printf("  UUID       %s\n", info.UUID)
	fmt.Printf("  CPI        %s\n", info.CPI)
	if info.DNSEnabled {
		fmt.Printf("  dns        %#v (%s)\n", info.DNSEnabled, info.DNSDomainName)
	} else {
		fmt.Printf("  dns        %#v\n", info.DNSEnabled)
	}
	if info.CompiledPackageCacheEnabled {
		fmt.Printf("  compiled_package_cache %#v (provider: %s)\n", info.CompiledPackageCacheEnabled, info.CompiledPackageCacheProvider)
	} else {
		fmt.Printf("  compiled_package_cache %#v\n", info.CompiledPackageCacheEnabled)
	}
	
	fmt.Printf("  snapshots  %#v\n", info.SnapshotsEnabled)
}
