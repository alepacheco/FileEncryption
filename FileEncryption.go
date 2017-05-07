package FileEncryption

import (
	"crypto/aes"
	"crypto/cipher"
	"io"
	"os"
)

var block cipher.Block
var iv [aes.BlockSize]byte
var stream cipher.Stream
var targetFileName string
var key []byte

// Ext is the encrypted appended extension
var Ext = ".enc"

// InitializeBlock Sets up the encription with a key
func InitializeBlock(myKey []byte, tfn string) {
	key = myKey
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	/*for index, b := range myIv {

		iv[index] = b

	}*/
	stream = cipher.NewCTR(block, iv[:])
	targetFileName = tfn
}

// StreamDecrypter decryps a file given its filepath
func StreamDecrypter(path string) (err error) {
	inFile, err := os.Open(path)
	if err != nil {
		//Couldn't open file, maybe a folder
		return
	}

	outFile, err := os.OpenFile(filenameDeobfuscator(path), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		return
	}
	defer outFile.Close()
	reader := &cipher.StreamReader{S: stream, R: inFile}
	if _, err = io.Copy(outFile, reader); err != nil {
		panic(err)
	}
	inFile.Close()

	//os.Remove(path)
	return
}

// StreamEncrypter encrypts a file given its filepatth
func StreamEncrypter(path string) (err error) {
	inFile, err := os.Open(path)
	if err != nil {
		return
	}
	outFile, err := os.OpenFile(filenameObfuscator(path), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		return
	}

	writer := &cipher.StreamWriter{S: stream, W: outFile}

	if _, err = io.Copy(writer, inFile); err != nil {
		panic(err)
	}
	inFile.Close()
	outFile.Close()
	//os.Remove(path)
	return nil
}

func filenameObfuscator(path string) string {
	/*filenameArr := strings.Split(path, string(os.PathSeparator))
	filename := filenameArr[len(filenameArr)-1]
	path2 := strings.Join(filenameArr[:len(filenameArr)-1], string(os.PathSeparator))

	return path2 + string(os.PathSeparator) + base64.Encode(filename) + Ext*/
	return path

}
func filenameDeobfuscator(path string) string {
	/*//get the path for the output
	opPath := strings.Trim(path, Ext)
	// Divide filepath
	filenameArr := strings.Split(opPath, string(os.PathSeparator))
	//Get base64 encoded filename
	filename := filenameArr[len(filenameArr)-1]
	// get parent dir
	path2 := strings.Join(filenameArr[:len(filenameArr)-1], string(os.PathSeparator))
	return path2 + string(os.PathSeparator) + base64.Decode(filename)*/
	return path
}
