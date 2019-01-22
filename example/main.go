package main

import (
	"fmt"
	"net/http"

	"../../captcha"
)

const (
	dx = 150
	dy = 50
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		c := captcha.New(dx, dy)
		c.CreateBackground(nil)
		c.DrawText("")
		c.Save(w)
	})

	fmt.Println("服务已启动...")
	http.ListenAndServe(":8800", nil)
}
