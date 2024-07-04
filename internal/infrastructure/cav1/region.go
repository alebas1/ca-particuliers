package cav1

import (
	"errors"
)

type Region struct {
	Name         string
	InternalName string
	Alias        string
}

type RegionMap map[string]Region

func getRegionalBankAlias(regionCode string) (string, error) {
	var regions RegionMap = RegionMap{
		"62": Region{Name: "Pas-de-Calais", InternalName: "Nord de France", Alias: "ca-norddefrance"},
	}
	alias := regions[regionCode].Alias
	if alias == "" {
		return "", errors.New("Region not found")
	}
	return alias, nil
}
