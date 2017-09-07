package main

import (
	"testing"
	"path/filepath"
	"io/ioutil"
	"os"
	gomniauthtest "github.com/stretchr/gomniauth/test"
)

func TestAuthAvatar(t *testing.T) {
	var authAvatar AuthAvatar
	testUser := &gomniauthtest.TestUser{}
	testUser.On("AvatarURL").Return("", ErrNoAvatarURL)
	testChatUser := &chatUser{User: testUser}
	url, err := authAvatar.GetAvatarURL(testChatUser)

	if err != ErrNoAvatarURL {
		t.Error("値が存在しない場合、AuthAvatar.GetAvatarURLは" + "ErrNoAvatarURLを返すべきです")
	}

	testUrl := "http://url-to-avatar"
	testUser = &gomniauthtest.TestUser{}
	testChatUser.User = testUser
	testUser.On("AvatarURL").Return(testUrl, nil)
	url, err = authAvatar.GetAvatarURL(testUser)
	if err != nil {
		t.Error("値が存在しない場合、AuthAvatar.GetAvatarURLは" + "ErrNoAvatarURLを返すべきはありません")
	} else {
		if url != testUrl {
			t.Error("値が存在しない場合、AuthAvatar.GetAvatarURLは" + "ErrNoAvatarURLを返すべきです")
		}
	}
}

func TestGravatarAvatar(t *testing.T)  {
	var gravatarAvatar GravatarAvatar
	user := &chatUser{uniqueID:"abc"}
	url, err := gravatarAvatar.GetAvatarURL(user)
	if err != nil {
		t.Error("GravatarAvatar.GetAvatarURLはエラーを返すべきではありません")
	}
	if url != "//www.gravatar.com/avatar/abc" {
		t.Errorf("GravatarAvatar.GetAvatarURLが%sという誤った値を返しました", url)
	}
}

func TestFileSystemAvatar(t *testing.T)  {
	filename := filepath.Join("avatars", "abc.jpg")
	ioutil.WriteFile(filename, []byte{}, 0777)
	defer func() {os.Remove(filename)}()

	var fileSystemAvatar FileSystemAvatar
	user := &chatUser{uniqueID:"abc"}
	url, err := fileSystemAvatar.GetAvatarURL(user)
	if err != nil {
		t.Error("FileSystemAvatar.GetAvatarURLはエラーを返すべきではありません。")
	}
	if url != "/avatars/abc.jpg" {
		t.Errorf("FileSystemAvatar.GetAvatarURLが%sという誤った値を返しました", url)
	}
}