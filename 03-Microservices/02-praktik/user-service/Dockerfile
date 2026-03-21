# 1. Gunakan sistem operasi Linux mini yang sudah terinstal Golang
FROM golang:alpine

# 2. Buat folder bernama /app di dalam kontainer Linux tersebut
WORKDIR /app

# 3. Pindahkan catatan resep (go.mod) ke dalam kontainer dan download Gin
COPY go.mod go.sum ./
RUN go mod download

# 4. Pindahkan file main.go Anda ke dalam kontainer
COPY *.go ./

# 5. Compile (Kompilasi) kode Golang Anda menjadi program mesin bernama "user-app"
RUN go build -o user-app

# 6. Buka lubang ventilasi di port 8081 agar bisa diakses dari luar kotak
EXPOSE 8081

# 7. Perintah pamungkas saat kontainer dinyalakan: Jalankan programnya!
CMD ["./user-app"]