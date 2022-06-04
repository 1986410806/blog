package service

import (
	"context"
	"fmt"
	"testing"
)

func TestGetUpToken(t *testing.T) {
	var token = GetUpToken(context.TODO())
	if token == "" {
		t.Fail()
	}
	fmt.Println("token:", token)
}
