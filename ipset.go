package ipset

import (
	"encoding/xml"
	"errors"
	"os/exec"
)

type IpSet struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Revision int64  `json:"revision"`
	Header   struct {
		Family     string `json:"family"`
		HashSize   int64  `json:"hash_size"`
		MaxElem    int64  `json:"max_elem"`
		MemSize    int64  `json:"max_size"`
		References int64  `json:"references"`
	} `json:"header"`
	Members []IpSetMember `json:"members"`
}

type IpSetMember struct {
	Elem string `json:"elem"`
}

type ListSetResult struct {
	XMLName xml.Name `xml:"ipsets"`
	IpSet   []struct {
		Name     string `xml:"name,attr"`
		Type     string `xml:"type"`
		Revision int64  `xml:"revision"`
		Header   struct {
			Family     string `xml:"family"`
			Hashsize   int64  `xml:"hashsize"`
			Maxelem    int64  `xml:"maxelem"`
			Memsize    int64  `xml:"memsize"`
			References int64  `xml:"references"`
		} `xml:"header"`
		Members struct {
			XMLName xml.Name `xml:"members"`
			Member  []struct {
				Elem string `xml:"elem"`
			} `xml:"member"`
		} `xml:"members"`
	} `xml:"ipset"`
}

func (s IpSet) ContainEntry(entry string) bool {
	for _, member := range s.Members {
		if entry == member.Elem {
			return true
		}
	}
	return false
}

// CreateSet Create a new set
func CreateSet(setName, typeName string, createOption ...string) error {
	args := append([]string{"create", setName, typeName}, createOption...)
	output, err := exec.Command("ipset", args...).CombinedOutput()
	if err != nil {
		return errors.New(string(output))
	}
	return nil
}

// DestroySet Destroy a named set or all sets
func DestroySet(setName string) error {
	args := []string{"destroy", setName}
	output, err := exec.Command("ipset", args...).CombinedOutput()
	if err != nil {
		return errors.New(string(output))
	}
	return nil
}

// AddEntry Add entry to the named set
func AddEntry(setName, entry string, addOption ...string) error {
	args := append([]string{"add", setName, entry}, addOption...)
	output, err := exec.Command("ipset", args...).CombinedOutput()
	if err != nil {
		return errors.New(string(output))
	}
	return nil
}

// DelEntry Delete entry to the named set
func DelEntry(setName, entry string) error {
	args := []string{"del", setName, entry}
	output, err := exec.Command("ipset", args...).CombinedOutput()
	if err != nil {
		return errors.New(string(output))
	}
	return nil
}

// GetSet Get the entries of a named set
func GetSet(setName string) (*IpSet, error) {
	ipSets, err := ListSet()
	if err != nil {
		return nil, err
	}

	for _, ipSet := range ipSets {
		if ipSet.Name == setName {
			return &ipSet, nil
		}
	}

	return nil, nil
}

// ListSet List the entries of a named set or all sets
func ListSet() ([]IpSet, error) {
	args := []string{"list", "-o", "xml"}
	output, err := exec.Command("ipset", args...).CombinedOutput()
	if err != nil {
		return nil, errors.New(string(output))
	}

	listSetResult := new(ListSetResult)
	if err := xml.Unmarshal(output, listSetResult); err != nil {
		return nil, err
	}

	ipSets := make([]IpSet, 0)
	for _, resultIpSet := range listSetResult.IpSet {
		var ipSet IpSet
		ipSet.Header.Family = resultIpSet.Header.Family
		ipSet.Header.HashSize = resultIpSet.Header.Hashsize
		ipSet.Header.MaxElem = resultIpSet.Header.Maxelem
		ipSet.Header.MemSize = resultIpSet.Header.Memsize
		ipSet.Header.References = resultIpSet.Header.References
		ipSet.Members = make([]IpSetMember, 0)
		for _, resultMember := range resultIpSet.Members.Member {
			var member IpSetMember
			member.Elem = resultMember.Elem
			ipSet.Members = append(ipSet.Members, member)
		}
		ipSet.Name = resultIpSet.Name
		ipSet.Revision = resultIpSet.Revision
		ipSet.Type = resultIpSet.Type
		ipSets = append(ipSets, ipSet)
	}

	return ipSets, nil
}
