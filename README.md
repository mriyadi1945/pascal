# Pascal Module

Module di buat dengan menggunakan bahasa Go dan hanya untuk ujicoba

## Instalasi

Untuk menggunakan modul ini dalam project Go Anda, jalankan perintah berikut:

```bash
go get github.com/mriyadi1945/pascal
```

## Prasyarat

- Go 1.16 atau lebih tinggi
- Git terinstall di sistem Anda

## Cara Penggunaan

### 1. Import Module

Tambahkan import di file Go Anda:

```go
import "github.com/mriyadi1945/pascal"
```

### 2. Inisialisasi Module

```go
package main

import (
    "fmt"
    "github.com/mriyadi1945/pascal"
)

func main() {
    // Contoh penggunaan
    // Sesuaikan dengan fungsi yang tersedia di module
}
```

## Tahapan Penggunaan

### Tahap 1: Setup Project

1. Buat project Go baru atau gunakan project yang sudah ada
```bash
mkdir myproject
cd myproject
go mod init myproject
```

2. Install module pascal
```bash
go get github.com/mriyadi1945/pascal
```

### Tahap 2: Import dan Konfigurasi

1. Import module di file main.go atau file lain yang memerlukan
2. Inisialisasi konfigurasi yang diperlukan (jika ada)

### Tahap 3: Implementasi

Gunakan fungsi-fungsi yang tersedia dari module sesuai kebutuhan aplikasi Anda.

## Contoh Penggunaan

```go
package main

import (
    "fmt"
    "log"
    "github.com/mriyadi1945/pascal"
)

func main() {
    // Contoh implementasi
    nilai := 15
	  pascal.PascalIterative(nilai)
}
```

## Dokumentasi

Untuk dokumentasi lengkap, kunjungi:
- [GoDoc](https://pkg.go.dev/github.com/mriyadi1945/pascal)
- [Repository GitHub](https://github.com/mriyadi1945/pascal)

## Kontak

Repository: [https://github.com/mriyadi1945/pascal](https://github.com/mriyadi1945/pascal)

---

**Catatan**: README ini adalah template umum. Sesuaikan dengan fungsi dan fitur aktual dari module pascal.
