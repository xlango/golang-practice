package main

import (
	"github.com/coreos/go-iptables/iptables"
)

func main() {
	ipt, err := iptables.New()

	if err != nil {

	}

	// Create a chain, add a rule
	chainName := "INPUT"

	err = ipt.NewChain("filter", chainName)

	if err != nil {

	}

	defer func() {
		ipt.ClearChain("filter", chainName)
		ipt.DeleteChain("filter", chainName)
	}()

	err = ipt.Append("filter", chainName, "-p", "tcp", "-j", "ACCEPT")

	if err != nil {

	}

}
