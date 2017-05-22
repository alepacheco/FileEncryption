package FileEncryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

var plaintext []byte

var block cipher.Block
var key []byte

// Ext is the encrypted appended extension
var Ext = ".enc"

func main() {
	return
}

// InitializeBlock Sets up the encription with a key
func InitializeBlock(myKey []byte) {
	key = myKey
	block, _ = aes.NewCipher(key)

}
func initIV() (stream cipher.Stream, iv []byte) {
	iv = make([]byte, aes.BlockSize)
	_, err := rand.Read(iv)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	stream = cipher.NewCTR(block, iv[:])
	return stream, iv
}
func initWithIV(myIv []byte) cipher.Stream {
	return cipher.NewCTR(block, myIv[:])
}

// Decrypter decryps a file given its filepath
func Decrypter(path string) (err error) {
	if block == nil {
		return errors.New("Need to Initialize Block first. Call: InitializeBlock(myKey []byte)")
	}

	inFile, err := os.Open(path)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	deobfPath := filenameDeobfuscator(path)
	outFile, err := os.OpenFile(deobfPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		return
	}

	iv := make([]byte, aes.BlockSize)
	io.ReadFull(inFile, iv[:])
	stream := initWithIV(iv)
	inFile.Seek(aes.BlockSize, 0) // Read after the IV

	reader := &cipher.StreamReader{S: stream, R: inFile}
	if _, err = io.Copy(outFile, reader); err != nil {
		fmt.Println(err)
	}
	inFile.Close()

	os.Remove(path)
	return
}

// Encrypter encrypts a file given its filepatth
func Encrypter(path string) (err error) {
	if block == nil {
		return errors.New("Need to Initialize Block first. Call: InitializeBlock(myKey []byte)")
	}

	inFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	obfuscatePath := filenameObfuscator(path)
	outFile, err := os.OpenFile(obfuscatePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	fmt.Println(outFile.Name())

	if err != nil {
		fmt.Println(err)
		return
	}

	stream, iv := initIV()
	outFile.Write(iv)
	writer := &cipher.StreamWriter{S: stream, W: outFile}

	if _, err = io.Copy(writer, inFile); err != nil {
		fmt.Println(err.Error())
	}
	inFile.Close()
	outFile.Close()
	os.Remove(path)
	return nil
}

func filenameObfuscator(path string) string {
	filenameArr := strings.Split(path, string(os.PathSeparator))
	filename := filenameArr[len(filenameArr)-1]
	path2 := strings.Join(filenameArr[:len(filenameArr)-1], string(os.PathSeparator))

	return path2 + string(os.PathSeparator) + filename + Ext

}
func filenameDeobfuscator(path string) string {
	//get the path for the output
	opPath := strings.Trim(path, Ext)
	// Divide filepath
	filenameArr := strings.Split(opPath, string(os.PathSeparator))
	//Get  filename
	filename := filenameArr[len(filenameArr)-1]
	// get parent dir
	path2 := strings.Join(filenameArr[:len(filenameArr)-1], string(os.PathSeparator))
	return path2 + string(os.PathSeparator) + filename
}
