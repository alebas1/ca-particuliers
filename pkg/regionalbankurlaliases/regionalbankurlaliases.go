package regionalbankurlaliases

import (
	"encoding/json"
	"errors"
	"io"
	"os"
)

type Region struct {
	Name         string `json:"name"`
	InternalName string `json:"internalName"`
	Alias        string `json:"alias"`
}

type RegionMap map[string]Region

func GetRegionalBankAlias(regionCode string) (string, error) {
	jsonFile, err := os.Open("./aliases_fr.json")
	if err != nil {
		return "", err
	}
	defer jsonFile.Close()

	var regions RegionMap
	byteValueJsonFile, err := io.ReadAll(jsonFile)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(byteValueJsonFile, &regions)
	if err != nil {
		return "", err
	}
	alias := regions[regionCode].Alias
	if alias == "" {
		return "", errors.New("Region not found")
	}
	return alias, nil
}
