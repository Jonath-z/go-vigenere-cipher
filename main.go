package main

import (
	"fmt"
	"strings"
)

var alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ~`!@#$%^&*()-_+={}[]|\\/:;\"'<>,.? "

var alphabetCharacterset = []rune(alphabet)

func contains(set []int, _index int) bool {
	for index, _ := range set {
		if index == _index {
			return true
		}
	}
	return false
}

func getLetterPositionOfKeyInTheAlphabet(key string) []int {
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
	encryptionKey := getLetterPositionOfKeyInTheAlphabet(key)
	encodedString := encode("Vigenere-cipher algorihtm", encryptionKey)
	decodedString := decode(encodedString, encryptionKey)
	encodedString2 := encode("Now that you've gotten your code to a stable place (nicely done, by the way), add a test. Testing your code during development can expose bugs that find their way in as you make changes. In this topic, you add a test for the Hello function. Note: This topic is part of a multi-part tutorial that begins with Create a Go module. Go's built-in support for unit testing makes it easier to test as you go. Specifically, using naming conventions, Go's testing package, and the go test command, you can quickly write and execute tests. In the greetings directory, create a file called greetings_test.go. Ending a file's name with _test.go tells the go test command that this file contains test functions. In greetings_test.go, paste the following code and save the file.", encryptionKey)
	decodedString2 := decode(encodedString2, encryptionKey)
	fmt.Println("Encoded String", encodedString)
	fmt.Println("Decoded String", decodedString)
	fmt.Println("Encoded String 1", encodedString2)
	fmt.Println("Decoded String 2", decodedString2)
}

func main() {
	vinegereCipher("kkwq!!!oin$&xjixsue/3*#(@o")
}
