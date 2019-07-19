package main

import "time"

type TheManSayHello struct {


}

func (man *TheManSayHello)Hello(ctx context.Context, req *String, res *String) error {
	req.Value = "hello: " + res.GetValue() + time.Now().String()
	return nil
}
