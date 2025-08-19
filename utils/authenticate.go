package utils

import (
	"context"
	"errors"
	"fmt"
	"os"
	"syscall"
	"time"

	"github.com/manifoldco/promptui"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/term"

	"github.com/abanoub-samy-farhan/safe-pass/client"


)

func Auth() bool {

	auth := client.InitiateClient(1)
	defer auth.Close()

	hashedPassword, e := auth.Get(context.Background(), "AUTH").Result()
	if e == redis.Nil {
		fmt.Print("User is not yet configured and has no password, run `sudo safe-pass init setup`")
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

		err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(pass))
		if err != nil {
			fmt.Print("Wrong password!")
			return false
		}


		auth.SetEx(context.Background(), "auth-tmp", EncryptData(string(pass)), 5*time.Minute)
	}
	fmt.Println()

	return true
}

func Setup() {
	if os.Getuid() != 0 {
		panic("Permssion denied")
	}
	validate := func(pass string) error {
		if len(pass) < 8 {
			return errors.New("password is too short")
		}
		return nil
	}

	fmt.Print("Checking Redis database status...")
	
	passwordPrompt := promptui.Prompt{
		Label: "Enter your password (minimum of 8 characters)",
		Mask: '*',
		Validate: validate,
	}
	
	password, err := passwordPrompt.Run()
	if err != nil {
		fmt.Printf("Error reading password: %v\n", err)
		return
	}
	auth := client.InitiateClient(1)
	if auth == nil {
		fmt.Print("Redis is not running or activated, make sure it's working properly.\n")
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	ctx := context.Background()
	auth.Set(ctx, "AUTH", string(hashedPassword), 0)
	defer auth.Close()
}
