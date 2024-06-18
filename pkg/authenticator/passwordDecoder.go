package authenticator

import "fmt"

func ComputePasswordCombination(passcode []string, keyLayout []string) []string {
	combination := []string{}
	for _, key := range passcode {
		keyIndex := 0
		for i, layoutKey := range keyLayout {
			if layoutKey == key {
				keyIndex = i
				break
			}
		}
		combination = append(combination, fmt.Sprint(keyIndex))
	}
	return combination
}
