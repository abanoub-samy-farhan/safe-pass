package utils

import (
	"context"
	"fmt"
	"syscall"
	"time"

	"github.com/redis/go-redis/v9"
	"golang.org/x/term"
)

func Auth() bool {

	auth := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       1,
	})
	defer auth.Close()

	_, e := auth.Get(context.Background(), "logs").Result()
	if e == redis.Nil {
		Setup()
		return false
	}

	_, err := auth.Get(context.Background(), "auth-tmp").Result()

	if err == redis.Nil {
		fmt.Printf("Enter your Master passkey: ")
		pass, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			fmt.Println("Error reading password:", err)
			return false
		}

		encryptedPass, err := auth.Get(context.Background(), "AUTH").Result()
		if err != nil {
			fmt.Println("Error reading password:", err)
			return false
		}

		if DecryptData(encryptedPass) != string(pass) {
			fmt.Println("Wrong passkey")
			return false
		}

		auth.SetEx(context.Background(), "auth-tmp", EncryptData(string(pass)), 5*time.Minute)
	}
	fmt.Println()

	return true
}

func Setup() {
	fmt.Println("Welcome to safe-pass")
	fmt.Printf("Please enter your master password to set up the database \n\n(Note that it would be unchangable)\nEnter passkey: ")
	pass, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Println("Error reading password:", err)
		return
	}

	encryptedPass := EncryptData(string(pass))
	auth := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       1,
	})
	defer auth.Close()

	auth.Set(context.Background(), "AUTH", encryptedPass, 0)
	auth.Set(context.Background(), "logs", true, 0)

	fmt.Println("\nDatabase set up successfully")
}
