package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"os"
	"time"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)


func GetSecretKey(key string) string {
	err := godotenv.Load(".env")

	if err != nil{
		log.Fatal(err)
	}

	return os.Getenv(key)
}

var secretKey = []byte(GetSecretKey("SECRET_KEY"))

func HashPassword(password string) (string, error) {
	hashpassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost )
	return string(hashpassword), err
}

func ComparePassword(hashedPassword, plainPassword string) error {
	err:= bcrypt.CompareHashAndPassword([]byte(hashedPassword),[]byte(plainPassword))
	return err
}

func CreateJwtToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userEmail": email,
		"expires": time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)

	return tokenString, err
}

func VerifyJwtToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token)(interface {}, error){
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("Invalid Token")
	}

	return nil
}


func EncryptPassword(value []byte, keyPharse string)[]byte{
	aesBlock, err := aes.NewCipher([]byte(keyPharse))

	if err != nil {
		fmt.Println(err)
	}

	gcmInstance,err := cipher.NewGCM(aesBlock)

	if err != nil {
		fmt.Println(err)
	}

	nonce := make([]byte, gcmInstance.NonceSize())
	_, _ = io.ReadFull(rand.Reader, nonce)

	cipheredText := gcmInstance.Seal(nonce, nonce, value, nil)

	return cipheredText

}

func DecryptPassword(ciphered []byte, keyPhrase string)[]byte{
	aesBlock, err := aes.NewCipher([]byte(keyPhrase))

	if err != nil {
		fmt.Println(err)
	}

	gcmInstance,err := cipher.NewGCM(aesBlock)

	if err != nil {
		fmt.Println(err)
	}

	nonceSize := gcmInstance.NonceSize()
	nonce, cipheredText := ciphered[:nonceSize], ciphered[nonceSize:]

	plainPwd, err := gcmInstance.Open(nil, nonce, cipheredText, nil)
	if err != nil {
		log.Fatalln(err)
	}
	return plainPwd

}

/* func DecryptPassword(ciphered []byte, keyPhrase string) []byte {
	aesBlock, err := aes.NewCipher([]byte(keyPhrase))
	if err != nil {
			log.Fatal(err)
	}

	gcmInstance, err := cipher.NewGCM(aesBlock)
	if err != nil {
			log.Fatal(err)
	}

	nonceSize := gcmInstance.NonceSize()
	if len(ciphered) < nonceSize {
			log.Fatal("ciphered text is too short")
	}

	nonce, cipheredText := ciphered[:nonceSize], ciphered[nonceSize:]

	plainPwd, err := gcmInstance.Open(nil, nonce, cipheredText, nil)
	if err != nil {
			log.Fatal(err)
	}
	return plainPwd
} */
