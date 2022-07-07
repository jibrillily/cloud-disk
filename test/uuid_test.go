package test

import (
	"fmt"
	"github.com/satori/go.uuid"
	"testing"
)

func TestGeneraterUUID(t *testing.T) {
	v4 := uuid.NewV4().String()
	fmt.Println(v4)

}
