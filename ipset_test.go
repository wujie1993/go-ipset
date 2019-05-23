package ipset_test

import (
	"encoding/json"
	"testing"

	"github.com/wujie1993/go-ipset"
)

const (
	testSetName  = "test_set"
	testSetEntry = "10.21.0.0/16"
)

func TestCreateSet(t *testing.T) {
	if err := ipset.CreateSet(testSetName, "hash:net"); err != nil {
		t.Fatal(err)
	}
	t.Logf("ok")
}

func TestGetSet(t *testing.T) {
	ipSet, err := ipset.GetSet(testSetName)
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.MarshalIndent(ipSet, "", "\t")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf(string(data))

	t.Logf("ok")
}

func TestListSet(t *testing.T) {
	ipSets, err := ipset.ListSet()
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.MarshalIndent(ipSets, "", "\t")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf(string(data))

	t.Logf("ok")
}

func TestAddEntry(t *testing.T) {
	if err := ipset.AddEntry(testSetName, testSetEntry); err != nil {
		t.Fatal(err)
	}
	t.Logf("ok")
}

func TestDelEntry(t *testing.T) {
	if err := ipset.DelEntry(testSetName, testSetEntry); err != nil {
		t.Fatal(err)
	}
	t.Logf("ok")
}

func TestDestroySet(t *testing.T) {
	if err := ipset.DestroySet(testSetName); err != nil {
		t.Fatal(err)
	}
	t.Logf("ok")
}
