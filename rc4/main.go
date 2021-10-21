// referred from: https://en.wikipedia.org/wiki/RC4

package RC4

type RC4 struct {
	key       string
	keyLength int
	sbox      []int
}

func NewRC4(key string) *RC4 {
	_rc4 := RC4{
		key:       key,
		keyLength: len(key),
	}

	_rc4.keyschedule()
	return &_rc4
}

// array index value swap utility
func (_rc4 *RC4) swap(firstIndex, secondIndex int) {
	temp := _rc4.sbox[firstIndex]
	_rc4.sbox[firstIndex] = _rc4.sbox[secondIndex]
	_rc4.sbox[secondIndex] = temp
}

// key schedule
func (_rc4 *RC4) keyschedule() {
	/*
		for i from 0 to 255
			S[i] := i
		endfor
		j := 0
		for i from 0 to 255
			j := (j + S[i] + key[i mod keylength]) mod 256
			swap values of S[i] and S[j]
		endfor
	*/
	maxCounter := 256

	for i := 0; i < maxCounter; i++ {
		_rc4.sbox = append(_rc4.sbox, i)
	}

	j := 0

	for i := 0; i < maxCounter; i++ {
		j = (j + _rc4.sbox[i] + int(_rc4.key[i%_rc4.keyLength])) % maxCounter
		_rc4.swap(i, j)
	}
}

// ciphertext[l] = plaintext[l] ⊕ K[l]
func (_rc4 *RC4) Encrypt(plainText []byte) []byte {
	cipherText := make([]byte, len(plainText))

	for i := 0; i < len(plainText); i++ {
		cipherText[i] = plainText[i] ^ byte(_rc4.sbox[i])
	}

	return cipherText
}
