package main

import (
	"errors"
)

var ErrNoAvatarURL = errors.New("chat: アバターのURLを取得できません。")

type Avatar interface {
	GetAvatarURL(c *client) (string, error)
}
