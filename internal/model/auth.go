package model

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"todo-list-server-with-go/internal/crypto"
)

func Signup(email, password string) (*Account, error) {
	account := Account{}
	// dbのAccountsテーブルに、このemailが存在するかを確認する
	// 存在する場合は、エラーを返す
	// 存在しない場合は、新しいAccountを作成し、返す
	// 作成したAccountは、DBに保存する
	db.Where("email = ?", email).First(&account)
	if account.ID != uuid.Nil {
		err := errors.New("this email is already registered")
		fmt.Println(err)
		return nil, err
	}

	encryptPw, err := crypto.PasswordEncrypt(password)
	if err != nil {
		fmt.Println("PasswordEncrypt Error: ", err)
		return nil, err
	}

	account = Account{ID: uuid.New(), Email: email, Password: encryptPw}
	db.Create(&account)
	return &account, nil
}

func Login(email, password string) (*Account, error) {
	account := Account{}
	db.Where("email = ?", email).First(&account)
	if account.ID == uuid.Nil {
		err := errors.New("this email is not registered")
		fmt.Println(err)
		return nil, err
	}

	err := crypto.CompareHashAndPassword(account.Password, password)
	if err != nil {
		fmt.Println("CompareHashAndPassword Error: ", err)
		return nil, err
	}

	return &account, nil
}
