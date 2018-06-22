# Amaliah

Daftar konten diambil berdasarkan dari Buku Risalah Amaliah, konten digital diambil dari berbagai sumber.

## Instalasi

Menggunakan Go Get

```cmd
go get github.com/taqiyyah/amaliah
```

atau [Dep](https://golang.github.io/dep/)

```cmd
dep ensure --add github.com/taqiyyah/amaliah
```

## Penggunaan

### Mengambil semua konten

```go
package main

import (
	"fmt"

	"github.com/taqiyyah/amaliah"
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

	"github.com/taqiyyah/amaliah"
)

func main() {
	amaliah40 := amaliah.Get(40)
	fmt.Println(amaliah40.Source)
}
```

## Kontribusi

Konten mungkin saja terdapat perbedaan dan tidak selengkap seperti yang ada di dalam Buku Risalah Amaliah, bantu kami untuk memperbaiki dan melengkapinya dengan membuat [issue](https://github.com/taqiyyah/amaliah/issues) pada repository ini. Atau bagi Anda seorang Developer, silakan Fork dan buat Pull Request ke repository ini.