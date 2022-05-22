package test

import (
	"context"
	"errors"
	"personal-page-back/domain"

	"testing"
)

type mockEmailGateWayNoError struct{}

func (m mockEmailGateWayNoError) SendEmail (email domain.Email) error{
	return nil
}

type mockEmailGateWayAllError struct{
	err error
}

func (m mockEmailGateWayAllError) SendEmail (email domain.Email) error{
	return m.err
}


func TestSendEmail(t *testing.T) {
	email := domain.Email{From: "joe@example.com", Body: "Hello world"}
	domain.InitEmailRepository(mockEmailGateWayNoError{}) 
	err := domain.SendMail(context.Background(), email)
	if err != nil{
		t.Error("Error is not nil")
	}
}

func TestSendEmailError(t *testing.T) {
	email := domain.Email{From: "joe@example.com", Body: "Hello world"}

	testError:= errors.New("Error")
	domain.InitEmailRepository(mockEmailGateWayAllError{err: testError}) 

	err := domain.SendMail(context.Background(), email)
	if err != testError{
		t.Error("Errors are not the same")
	}
}