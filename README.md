# go-stemming

This library is for stemming word

## Instalasi

```bash
go get github.com/pycodeDev/go-stemming
```

## How To Use
Get Your Merchant ID and Secret Key at [Tokopay Dashboard](https://dash.go-tokopay.id/pengaturan/secret-key).

### Prefixes and Suffixes support

```golang
 []string{"peng", "pen", "ber", "per", "pe", "di", "ke", "se", "te", "me"}
	[]string{"lah", "kah", "tah", "pun", "ku", "mu", "nya", "kan", "an"}
```
> Note:
> New Version Of Stemming

### one word stem

```golang
import go_stemming "github.com/pycodeDev/go-stemming"

func main() {
    word := go_stemming.OneWordStem("berlari")
}
```

### multiple word stem
preparation for data
```golang
import go_stemming "github.com/pycodeDev/go-stemming"

func main() {
    data := []string{"berlari", "menyanyi", "beramal"}
    words := go_stemming.MultipleWordStem(data)
}
```

### License

[MIT]

### Author

[Muhammad Anang Ramadhan](mailto:muhammadanangr@gmail.com)
