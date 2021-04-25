package go_camp

import "fmt"

func AuthenticateRequestOrigin(r *Request) error  {
	err := authenticate(r.User)
	if err != nil {
		return err
	}
	return nil
}

func AuthenticateRequest(r *Request) error {
	return authenticate(r.User)
}

func AuthenticateRequestOptimized(r *Request) error {
	err := authenticate(r.User)
	if err != nil {
		return fmt.Errorf("authenticate failed: %v", err)
	}
	return nil
}

func authenticate(user string) error {
	return nil
}

type Request struct {
	User string
}
