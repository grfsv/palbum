package main

import (
	"fmt"
	"io"
	"os"
	"palbum/internal/infrastructure/persistence/auth"
	"palbum/internal/infrastructure/persistence/friend"
	"palbum/internal/infrastructure/persistence/user"

	"ariga.io/atlas-provider-gorm/gormschema"
)

func main() {
	stmts, err := gormschema.New("mysql").Load(
		&user.User{},
		&auth.Auth{},
		&friend.FriendCode{},
		&friend.FriendRequest{},
		&friend.Friendship{},
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load gorm schema: %v\n", err)
		os.Exit(1)
	}

	_, _ = io.WriteString(os.Stdout, stmts)
}
