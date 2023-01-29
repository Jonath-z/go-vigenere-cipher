package main

import (
	"fmt"
	"strings"
)

var alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ~`!@#$%^&*(-_+={}[]|\\/:;\"'<>,.? "

var alphabetCharacterset = []rune(alphabet)

func contains(set []int, _index int) bool {
	for index, _ := range set {
		if index == _index {
			return true
		}
	}
	return false
}

func getLetterPositionOfKeyInTheAlphabet(key string, alphabet string) []int {
	// set of key's letter position in the alphabet
	set := []int{}

	//Array of key's characters
	keyCharacterSet := []rune(key)

	// Loop through keyCharacterSet and add the postion to set
	for i := 0; i < len(keyCharacterSet); i++ {
		for j := 0; j < len(alphabetCharacterset); j++ {
			if keyCharacterSet[i] == alphabetCharacterset[j] {
				set = append(set, j)
			}
		}
	}

	return set
}

func encode(str string, encryptionKey []int) string {
	encodedString := []string{}

	// str character set
	strCharacterSet := []rune(str)
	fmt.Println("string set", strCharacterSet)

	for i := 0; i < len(strCharacterSet); i++ {

		// check if strCharacterSet current index exist in key
		if contains(encryptionKey, i) {

			// jump = index of the current character in the aplhabet adding the encryption key letter index
			jumpTo := strings.IndexRune(alphabet, strCharacterSet[i]) + encryptionKey[i]

			// check if jump is greater then the aphabet length
			if jumpTo > len(alphabetCharacterset) {
				encodedString = append(
					encodedString,
					string(alphabetCharacterset[jumpTo%len(alphabetCharacterset)]))
			} else {
				encodedString = append(
					encodedString,
					string(alphabetCharacterset[jumpTo]))
			}

		} else {
			jumpTo := strings.IndexRune(
				alphabet,
				strCharacterSet[i]) + encryptionKey[i%len(encryptionKey)]

			if jumpTo > len(alphabetCharacterset)-1 {
				encodedString = append(
					encodedString,
					string(alphabetCharacterset[jumpTo%len(alphabetCharacterset)]))
			} else {
				encodedString = append(
					encodedString,
					string(alphabetCharacterset[jumpTo]))
			}
		}
	}

	return strings.Join(encodedString, "")

}

func decode(encodedStr string, encryptionKey []int) string {
	decodedString := []string{}

	//encodedStr character set
	encodedStrCharacterSet := []rune(encodedStr)

	for i := 0; i < len(encodedStrCharacterSet); i++ {

		if contains(encryptionKey, i) {
			jumpTo := strings.IndexRune(alphabet, encodedStrCharacterSet[i]) - encryptionKey[i]

			if jumpTo < 0 {
				jumpTo += len(alphabetCharacterset)
			}

			decodedString = append(decodedString, string(alphabetCharacterset[jumpTo]))
		} else {
			jumpTo := strings.IndexRune(alphabet, encodedStrCharacterSet[i]) - encryptionKey[i%len(encryptionKey)]

			if jumpTo < 0 {
				jumpTo += len(alphabetCharacterset)
			}

			decodedString = append(decodedString, string(alphabetCharacterset[jumpTo]))
		}
	}

	return strings.Join(decodedString, "")
}

func vinegereCipher(key string) {
	encryptionKey := getLetterPositionOfKeyInTheAlphabet(key, alphabet)
	encodedString := encode("Vigenere-cipher algorihtm", encryptionKey)
	decodedString := decode(encodedString, encryptionKey)
	fmt.Println("Encoded String", encodedString)
	fmt.Println("Decoded String", decodedString)
}

func main() {
	vinegereCipher("kkwq!!!oin$&xjixsue/3*#(@o")
}
