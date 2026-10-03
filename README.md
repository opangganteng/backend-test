# Backend Test - Penjumlahan Bilangan Genap

Program ini membaca integer dari standard input, membaginya menjadi beberapa
bagian, lalu menjumlahkan bilangan genap secara konkuren menggunakan goroutine.
Jumlah worker default adalah 4 dan dapat diubah dengan flag `-workers`.

Setiap worker mengirim subtotal melalui channel. `sync.WaitGroup` memastikan
semua worker selesai sebelum channel ditutup dan seluruh subtotal dijumlahkan.

## Menjalankan

```sh
cd backend-test
go run . -workers 4
```

Masukkan integer yang dipisahkan oleh spasi atau baris baru, lalu akhiri input
dengan EOF (Ctrl+D di Linux/macOS atau Ctrl+Z lalu Enter di Windows). Program
mencetak jumlah bilangan genap.

Contoh:

```sh
echo "1 2 3 4 5 6" | go run . -workers 4
```

Output:

```text
12
```

## Pengujian

```sh
go test ./...
go test -race ./...
```
