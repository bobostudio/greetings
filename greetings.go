package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	return fmt.Sprintf(randomGreeting(), name), nil
}

func Names(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _,name := range names {
		message,err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}


func randomGreeting() string {
	greetings := []string{
		"Hello, %v 世界！",
		"你好，%v 世界！",
		"欢迎学习 Go,%v ！",
		"Go 语言真有趣！%v",
		"Go 语言是一个强大的编程语言！,%v",
	}
	return greetings[rand.Intn(len(greetings))]
}
