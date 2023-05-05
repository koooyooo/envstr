
![logo](./logo.jpg)

# envstr
[![test](https://github.com/koooyooo/envstr/actions/workflows/test.yaml/badge.svg)](https://github.com/koooyooo/envstr/actions/workflows/test.yaml)
[![lint](https://github.com/koooyooo/envstr/actions/workflows/lint.yaml/badge.svg)](https://github.com/koooyooo/envstr/actions/workflows/lint.yaml)

envstr は文字列に環境変数の値を埋め込むライブラリです。

## Usage

`go get` で取得
```bash
$ go get github.com/koooyooo/envstr
```

### Example
以下の利用例の形式で利用
```go
os.Setenv("Hello", "World")
result := envstr.Apply("I like the word ${Hello} !")
fmt.Println(result) // I like the word World !
```
