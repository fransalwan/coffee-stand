# ☕ Stand Coffee OS (MVP)

> *"Bukan sekadar aplikasi kasir. Ini adalah Sistem Operasi buat stand kopi biar owner-nya bisa fokus nyeduh dan ngobrol, sementara sistem yang ngurusin antrean dan catat duit."*

## 🧠 Filosofi & Latar Belakang
Project ini dibangun dengan prinsip **"Scratch Your Own Itch"** dan **"Via Negativa"** (menambah nilai dengan mengurangi friksi). 

Banyak sistem kasir/POS di luar sana yang ribet, penuh *dark patterns*, dan bikin pusing owner stand kopi yang margin-nya tipis. **Kopi Teman** dibangun dengan tujuan:
1. **Peace of Mind buat Owner:** Dashboard yang cuma nunjukin 3 angka ajaib (Omzet, Cup Terjual, Stok Kritis). Nggak ada grafik yang bikin pusing.
2. **Frictionless buat Customer:** Menu "Jujur" (ada *storytelling* bahan baku) dan fitur *Skip The Queue* (Pre-order via WA).
3. **Habit Building:** Sistem stempel digital yang ngunci kebiasaan pelanggan tanpa perlu kartu fisik yang gampang ilang.

## 🛠 Tech Stack
Kita pakai *stack* yang *bulletproof*, cepet, dan *low-resource* (biar nggak bakar duit server di fase awal).

**Backend (Current Focus):**
- **Language:** Golang (Concurrent, cepat, low memory footprint)
- **Framework:** Gin (Lightweight HTTP router)
- **ORM:** GORM (Untuk kecepatan development MVP)
- **Database:** PostgreSQL (Relational, solid)

**Frontend (Planned):**
- **Framework:** Vue 3 (Composition API) + TypeScript
- **State Management:** Pinia
- **Styling:** Tailwind CSS (Mobile-first, utility-first)

## 📂 Struktur Project (Backend)
```text
backend/
├── cmd/
│   └── main.go              # Entry point & routing
├── internal/
│   ├── database/
│   │   ├── database.go      # Koneksi Postgres & Auto-migrate
│   │   └── seeder.go        # Isi data dummy "Menu Jujur"
│   ├── handlers/            # Logic HTTP (Controllers)
│   └── models/              # Struct Database (GORM Models)
├── .env                     # Environment variables (Wajib ada)
├── go.mod
└── go.sum