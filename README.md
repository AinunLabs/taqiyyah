# Amaliah

Daftar konten diambil berdasarkan dari Buku Risalah Amaliah, konten digital diambil dari berbagai sumber.

## Instalasi

Menggunakan Go Get

```cmd
go get github.com/letstaqwa/amaliah
```

atau [Dep](https://golang.github.io/dep/)

```cmd
dep ensure --add github.com/letstaqwa/amaliah
```

## Penggunaan

### Mengambil semua konten

```go
package main

import (
	"fmt"

	"github.com/letstaqwa/amaliah"
)

func main() {
	listOfAmaliah := amaliah.All()

	for _, amaliah := range listOfAmaliah {
		fmt.Println(amaliah.Titles[0].Value)
		fmt.Println(amaliah.Contents[0].Value)
		fmt.Println(amaliah.Contents[1].Value)
	}
}
```

### Mengambil hanya 1 konten

```go
package main

import (
	"fmt"

	"github.com/letstaqwa/amaliah"
)

func main() {
	amaliah40 := amaliah.Get(40)
	fmt.Println(amaliah40.Titles[0].Value)
	fmt.Println(amaliah40.Contents[0].Value)
	fmt.Println(amaliah40.Contents[1].Value)
}
```

## Kontribusi

Konten mungkin saja terdapat perbedaan dan tidak selengkap seperti yang ada di dalam Buku Risalah Amaliah, bantu kami untuk memperbaiki dan melengkapinya dengan membuat [issue](https://github.com/letstaqwa/amaliah/issues) pada repository ini. Atau bagi Anda seorang Developer, silakan Fork dan buat Pull Request ke repository ini.
