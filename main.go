package main

import (
	"fmt"
	"time"

	"github.com/hashicorp/memberlist"
)

func main() {
	/* Create the initial memberlist from a safe configuration.
	   Please reference the godoc for other default config types.
	   http://godoc.org/github.com/hashicorp/memberlist#Config
	*/
	list, err := memberlist.Create(memberlist.DefaultLANConfig())
	if err != nil {
		panic("Failed to create memberlist: " + err.Error())
	}

	// Join an existing cluster by specifying at least one known member.
	_, err = list.Join([]string{"127.0.0.1"})
	if err != nil {
		panic("Failed to join cluster: " + err.Error())
	}

	for {
		time.Sleep(5 * time.Second)

		// Ask for members of the cluster
		for _, member := range list.Members() {
			fmt.Printf("Member: %s %s\n", member.Name, member.Addr)
		}
	}

	// Continue doing whatever you need, memberlist will maintain membership
	// information in the background. Delegates can be used for receiving
	// events when members join or leave.
}
