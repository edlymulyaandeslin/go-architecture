package testing

import "errors"

func SayHello(name string) (string, error) {
	if name == "" {
		return "", errors.New("uhhh error no name")
	}
	return "hello " + name, nil
}
