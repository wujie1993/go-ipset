package ipset

import (
	"encoding/xml"
	"errors"
	"os/exec"
)

// IPSet defines the struct of ipset
type IPSet struct {
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
	Members []Member `json:"members"`
}

// Member is the member of ipset
type Member struct {
	Elem string `json:"elem"`
}

// ListSetResult is the return struct of ipset list
type ListSetResult struct {
	XMLName xml.Name `xml:"ipsets"`
	IPSet   []struct {
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

// ContainEntry detect if entry already exist in ipset
func (s IPSet) ContainEntry(entry string) bool {
	for _, member := range s.Members {
		if entry == member.Elem {
			return true
		}
	}
	return false
}

// CreateSet create a new ipset
// setName: the name of ipset
// typeName: the type of ipset
// createOpotion: the options of ipset
func CreateSet(setName, typeName string, createOption ...string) error {
	args := append([]string{"create", setName, typeName}, createOption...)
	output, err := exec.Command("ipset", args...).CombinedOutput()
	if err != nil {
		return errors.New(string(output))
	}
	return nil
}

// DestroySet destroy a ipset .It will delete all ipset when setName is empty.
// setName: the name of ipset
func DestroySet(setName string) error {
	args := []string{"destroy", setName}
	output, err := exec.Command("ipset", args...).CombinedOutput()
	if err != nil {
		return errors.New(string(output))
	}
	return nil
}

// AddEntry add entry to the named set
func AddEntry(setName, entry string, addOption ...string) error {
	args := append([]string{"add", setName, entry}, addOption...)
	output, err := exec.Command("ipset", args...).CombinedOutput()
	if err != nil {
		return errors.New(string(output))
	}
	return nil
}

// DelEntry delete entry to the named set
func DelEntry(setName, entry string) error {
	args := []string{"del", setName, entry}
	output, err := exec.Command("ipset", args...).CombinedOutput()
	if err != nil {
		return errors.New(string(output))
	}
	return nil
}

// GetSet get the entries of a named set
func GetSet(setName string) (*IPSet, error) {
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

// ListSet list the entries of a named set or all sets
func ListSet() ([]IPSet, error) {
	args := []string{"list", "-o", "xml"}
	output, err := exec.Command("ipset", args...).CombinedOutput()
	if err != nil {
		return nil, errors.New(string(output))
	}

	listSetResult := new(ListSetResult)
	if err := xml.Unmarshal(output, listSetResult); err != nil {
		return nil, err
	}

	ipSets := make([]IPSet, 0)
	for _, resultIPSet := range listSetResult.IPSet {
		var ipSet IPSet
		ipSet.Header.Family = resultIPSet.Header.Family
		ipSet.Header.HashSize = resultIPSet.Header.Hashsize
		ipSet.Header.MaxElem = resultIPSet.Header.Maxelem
		ipSet.Header.MemSize = resultIPSet.Header.Memsize
		ipSet.Header.References = resultIPSet.Header.References
		ipSet.Members = make([]Member, 0)
		for _, resultMember := range resultIPSet.Members.Member {
			var member Member
			member.Elem = resultMember.Elem
			ipSet.Members = append(ipSet.Members, member)
		}
		ipSet.Name = resultIPSet.Name
		ipSet.Revision = resultIPSet.Revision
		ipSet.Type = resultIPSet.Type
		ipSets = append(ipSets, ipSet)
	}

	return ipSets, nil
}
