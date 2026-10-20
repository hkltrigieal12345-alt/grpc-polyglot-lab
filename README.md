# gRPC Polyglot Lab

## Deskripsi Proyek

Proyek ini merupakan implementasi komunikasi antarbahasa pemrograman menggunakan gRPC dan Protocol Buffers (Protobuf).

Pada proyek ini:
- Go digunakan sebagai server gRPC.
- Python digunakan sebagai client gRPC.
- Node.js digunakan sebagai client gRPC.
- Protobuf digunakan untuk mendefinisikan struktur data dan layanan komunikasi.

Server menyediakan layanan untuk mengambil profil pengguna berdasarkan ID.

## Teknologi yang Digunakan

- Go
- Python
- Node.js
- gRPC
- Protocol Buffers
- Visual Studio Code

## Struktur Folder

```text
grpc-polyglot-lab/
├── go-server/
│   ├── pb/
│   │   ├── user.pb.go
│   │   └── user_grpc.pb.go
│   ├── main.go
│   ├── go.mod
│   └── go.sum
├── python-client/
│   ├── client.py
│   ├── user_pb2.py
│   └── user_pb2_grpc.py
├── node-client/
│   ├── client.js
│   ├── package.json
│   └── package-lock.json
└── user.proto
```

## Persyaratan

Pastikan perangkat telah memiliki:
- Go
- Python
- Node.js dan npm
- Protocol Buffers Compiler (protoc)

## Instalasi dan Persiapan

### 1. Mengunduh dependensi Go

Buka terminal di folder proyek, kemudian jalankan:

```powershell
cd go-server
go mod tidy
```

Setelah selesai, kembali ke folder utama:

```powershell
cd ..
```

### 2. Memasang dependensi Python

Jalankan:

```powershell
cd python-client
py -m pip install grpcio grpcio-tools
```

Kembali ke folder utama:

```powershell
cd ..
```

### 3. Memasang dependensi Node.js

Jalankan:

```powershell
cd node-client
npm install
```

Kembali ke folder utama:

```powershell
cd ..
```

## Menjalankan Program

### 1. Menjalankan Server Go

Buka terminal pertama di folder utama proyek.

Jalankan:

```powershell
cd go-server
go run main.go
```

Jika berhasil, terminal menampilkan:

```text
Server gRPC berjalan pada port 50051
```

Biarkan terminal ini tetap terbuka selama pengujian client.

### 2. Menjalankan Python Client

Buka terminal kedua di VS Code.

Jalankan:

```powershell
cd "$HOME\Downloads\SEMESTER 5\PAK REZA\grpc-polyglot-lab\python-client"
py client.py
```

Output yang diharapkan:

```text
=== PROFIL PENGGUNA ===
Nama  : Budi Santoso
Email : budi@universitas.ac.id
Telepon: 081234567890
```

### 3. Menjalankan Node.js Client

Buka terminal ketiga di VS Code.

Jalankan:

```powershell
cd "$HOME\Downloads\SEMESTER 5\PAK REZA\grpc-polyglot-lab\node-client"
node client.js
```

Output yang diharapkan:

```text
=== PROFIL PENGGUNA ===
Nama   : Budi Santoso
Email  : budi@universitas.ac.id
Telepon: 081234567890
```

## Definisi Layanan

Layanan didefinisikan di dalam file `user.proto`.

Pesan `UserRequest` digunakan untuk mengirimkan ID pengguna.

Pesan `UserResponse` digunakan untuk mengembalikan informasi pengguna berupa:
- Nama
- Email
- Nomor telepon

Layanan `UserService` memiliki metode RPC `GetUserProfile` untuk mengambil profil pengguna berdasarkan ID.

## Pengujian

Pengujian dilakukan dengan menjalankan server Go dan menghubungkannya menggunakan dua client berbeda.

Hasil pengujian yang diharapkan:
1. Server Go berjalan pada port 50051.
2. Python Client berhasil menerima profil pengguna.
3. Node.js Client berhasil menerima profil pengguna.
4. Kedua client menampilkan nama, email, dan nomor telepon.
5. Permintaan dengan ID pengguna yang tidak tersedia menghasilkan status `NotFound`.

## Kesimpulan

Proyek ini menunjukkan penggunaan gRPC dan Protocol Buffers untuk komunikasi antara server Go dengan client Python dan Node.js.

Dengan menggunakan definisi layanan yang sama, beberapa bahasa pemrograman dapat berkomunikasi melalui antarmuka RPC yang konsisten.