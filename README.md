# Vigenère Cipher

This is Vigenère Cipher algorithm written in go

## Description

The Vigenère cipher is a classic cipher originally developed by Italian cryptographer Giovan Battista Bellaso and published in 1553. It is named after a later French cryptographer Blaise de Vigenère, who had developed a stronger autokey cipher (a cipher that incorporates the message of the text into the key).

The cipher survived three centuries of attempts to break it, earning it the nickname "le chiffre indéchiffrable" or "the indecipherable cipher."

From Wikipedia:

The Vigenère cipher is a method of encrypting alphabetic text by using a series of different Caesar ciphers based on the letters of a keyword. It is a simple form of polyalphabetic substitution.

. . .

In a Caesar cipher, each letter of the alphabet is shifted along some number of places; for example, in a Caesar cipher of shift 3, A would become D, B would become E, Y would become B and so on. The Vigenère cipher consists of several Caesar ciphers in sequence with different shift values.

Assume the key is repeated for the length of the text, character by character. Note that some implementations repeat the key over characters only if they are part of the alphabet -- this is not the case here.

The shift is derived by applying a Caesar shift to a character with the corresponding index of the key in the alphabet.

Visual representation:

```go
"my secret code i want to secure"; // message
"passwordpasswordpasswordpasswor"; // key
```

## Example

```go
	encodedString := encode("Vigenere-cipher algorihtm", encryptionKey)
	decodedString := decode(encodedString, encryptionKey)

	fmt.Println("Encoded String", encodedString) //$sCu{%|s|p+;EnzwsFkd<_[:A
	fmt.Println("Decoded String", decodedString) // Vigenere-cipher algorihtm
```

## Local test

1. Clone the repository

```bash
    git clone https://github.com/Jonath-z/go-vigenere-cipher.git
```

2. Navigate into the cloned folder
3. Run

```shell
  go run main.go
```

## Author

- [Jonathan z.](https://www.linkedin.com/in/jonathan-z-0a40ab209/)
