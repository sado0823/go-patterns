package main

import "fmt"

type (
	// Option 选项设计模式
	Option func(config *Config)

	Config struct {
		Key    string
		UID    int64
		Age    int64
		Avatar string
	}
)

func WithUID(uid int64) Option {
	return func(config *Config) {
		config.UID = uid
	}
}

func WithAge(age int64) Option {
	return func(config *Config) {
		config.Age = age
	}
}

func WithAvatar(avatar string) Option {
	return func(config *Config) {
		config.Avatar = avatar
	}
}

func NewConfig(key string, fns ...Option) *Config {
	// default config
	cfg := &Config{
		Key:    key,
		UID:    888,
		Age:    100,
		Avatar: "no avatar",
	}

	for _, fn := range fns {
		fn(cfg)
	}

	return cfg
}

func main() {
	cfg1 := NewConfig("test1")
	fmt.Println("cfg1: ", cfg1)

	cfg2 := NewConfig("test2", WithUID(123), WithAge(18), WithAvatar("avatar2"))
	fmt.Println("cfg2: ", cfg2)
}
