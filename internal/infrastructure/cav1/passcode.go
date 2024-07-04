package cav1

import (
	"errors"
	"fmt"
)

func computePasscodeCombination(passcode []string, keyLayout []string) ([]string, error) {
	combination := []string{}
	for _, key := range passcode {
		if keyIndex, keyPresent := getKeyPositionInKeyLayout(key, keyLayout); keyPresent {
			combination = append(combination, fmt.Sprint(keyIndex))
		} else {
			return combination, errors.New("failed to compute password combination, key layout seems invalid")
		}
	}
	return combination, nil
}

func getKeyPositionInKeyLayout(key string, keyLayout []string) (int, bool) {
	for layoutPos, layoutKey := range keyLayout {
		if key == layoutKey {
			return layoutPos, true
		}
	}
	return 0, false
}
