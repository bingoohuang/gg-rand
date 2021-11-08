# Go生成图形验证码

Go生成图形验证码，带噪点，干扰线，正弦干扰线；支持生成普通的字符串和简单的数学算术运算公式,

source fork from [here](github.com/vcqr/captcha), font from [lifei6671/gocaptcha](https://github.com/lifei6671/gocaptcha/)

## Example

```go
package main

import (
	"github.com/bingoohuang/gg-rand/pkg/c7a"
	"image/png"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		cp := c7a.NewCaptcha(150, 40, 5)
		// cp.SetMode(1) // 设置为数学公式
		code, img := cp.OutPut()
		// 备注：code 可以根据情况存储到session，并在使用时取出验证

		log.Printf("Captcha: %s", code)
		w.Header().Set("Content-Type", "image/png; charset=utf-8")
		png.Encode(w, img)
	})

	if err := http.ListenAndServe("localhost:1200", nil); err != nil {
		log.Fatal(err)
	}
}

```
