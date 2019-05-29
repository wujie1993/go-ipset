# go-ipset

[![GoDoc](https://godoc.org/github.com/wujie1993/go-ipset?status.svg)](https://godoc.org/github.com/wujie1993/go-ipset)
[![Go Report Card](https://goreportcard.com/badge/github.com/wujie1993/go-ipset)](https://goreportcard.com/report/github.com/wujie1993/go-ipset)
[![CircleCI](https://circleci.com/gh/wujie1993/go-ipset.svg?style=svg)](https://circleci.com/gh/wujie1993/go-ipset)

## Example

```golang
package main

import (
	"fmt"

	"github.com/wujie1993/go-ipset"
)

func main() {
	var setName string = "test_set"
	var setEntry string = "192.168.1.0/24"

	// Create a ipset by set name and set type
	if err := ipset.CreateSet(setName, "hash:net"); err != nil {
		fmt.Println(err)
		return
	}

	// Get a ipset by set name
	ipSet, err := ipset.GetSet(setName)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v", ipSet)

	// List all ipset
	ipSets, err := ipset.ListSet()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v", ipSets)

	// Add a entry in ipset
	if err := ipset.AddEntry(setName, setEntry); err != nil {
		fmt.Println(err)
		return
	}

	// Delete a entry in ipset
	if err := ipset.DelEntry(setName, setEntry); err != nil {
		fmt.Println(err)
		return
	}

	// Destroy a ipset
	if err := ipset.DestroySet(setName); err != nil {
		fmt.Println(err)
		return
	}
}
```

For more usages, please reading [godoc](https://godoc.org/github.com/wujie1993/go-ipset) or [ipset_test.go](./ipset_test.go)
