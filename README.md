# ☕ Stand Coffee OS (MVP)

> *"Bukan sekadar aplikasi kasir. Ini adalah Sistem Operasi buat stand kopi biar owner-nya bisa fokus nyeduh dan ngobrol, sementara sistem yang ngurusin antrean dan catat duit."*

**🔗 Live Demo & Progress:** [github.com/fransalwan/coffee-stand](https://github.com/fransalwan/coffee-stand)

---

## 🧠 Filosofi & Latar Belakang

Project ini dibangun dengan prinsip **"Scratch Your Own Itch"** dan **"Via Negativa"** (menambah nilai dengan mengurangi friksi). 

Banyak sistem kasir/POS di luar sana yang ribet, penuh *dark patterns*, dan bikin pusing owner stand kopi yang margin-nya tipis. **Stand Coffee OS** dibangun dengan tujuan:

1. **Peace of Mind buat Owner:** Dashboard yang cuma nunjukin 3 angka ajaib (Omzet, Cup Terjual, Stok Kritis). Nggak ada grafik yang bikin pusing.
2. **Frictionless buat Customer:** Menu "Jujur" (ada *storytelling* bahan baku) dan fitur *Skip The Queue* (Pre-order via WA).
3. **Habit Building:** Sistem stempel digital yang ngunci kebiasaan pelanggan tanpa perlu kartu fisik yang gampang ilang.

### 🎯 Mission
**Democratize access to good tools.** Siapa aja yang punya mimpi buka stand kopi harusnya bisa mulai tanpa ribet dan tanpa keluar duit jutaan buat sistem.

---

## 🚀 Progress & Status

### ✅ Completed (Phase 1-2)
- [x] Backend API dengan Go + Gin + Postgres
- [x] Database schema dengan price_snapshot untuk akurasi laporan
- [x] Auto-migrate & seeder dengan data dummy "Menu Jujur"
- [x] CRUD Menu API (Public & Admin routes)
- [x] Customer management via WhatsApp
- [x] Order system dengan auto +1 stempel
- [x] Loyalty stamp system (redeem 10 stamps = 1 kopi gratis)
- [x] Frontend Vue 3 + TypeScript setup
- [x] Tailwind CSS dengan custom coffee color palette
- [x] API integration & fetching real-time data
- [x] WhatsApp checkout integration

### 🚧 In Progress (Phase 3)
- [ ] Cart system & order management UI
- [ ] Owner Dashboard (input order, manage stamps)
- [ ] Dashboard stats (omzet harian, top menu)
- [ ] Deploy to production (Railway + Vercel)

### 📋 Planned (Phase 4+)
- [ ] JWT authentication untuk admin
- [ ] Background jobs (daily report, low-stock alerts)
- [ ] Multi-branch support
- [ ] WA Gateway integration (auto-reply)

---

## 🛠 Tech Stack

Kita pakai *stack* yang *bulletproof*, cepet, dan *low-resource* (biar nggak bakar duit server di fase awal).

### Backend
- **Language:** Golang 1.21+ (Concurrent, cepat, low memory footprint)
- **Framework:** Gin (Lightweight HTTP router)
- **ORM:** GORM (Untuk kecepatan development MVP)
- **Database:** PostgreSQL 14+ (Relational, ACID-compliant)

### Frontend
- **Framework:** Vue 3 (Composition API) + TypeScript
- **State Management:** Pinia
- **Styling:** Tailwind CSS v3 (Mobile-first, utility-first)
- **HTTP Client:** Axios
- **Build Tool:** Vite

### Deployment (Planned)
- **Backend:** Railway / Fly.io
- **Frontend:** Vercel / Cloudflare Pages
- **Database:** Neon / Supabase

---

##  Struktur Project

```text
coffee-stand/
├── backend/
│   ├── cmd/
│   │   └── main.go              # Entry point & routing
│   ├── internal/
│   │   ├── database/
│   │   │   ├── database.go      # Koneksi Postgres & Auto-migrate
│   │   │   └── seeder.go        # Isi data dummy "Menu Jujur"
│   │   ├── handlers/            # Logic HTTP (Controllers)
│   │   │   ├── menu_handler.go
│   │   │   ├── customer_handler.go
│   │   │   ├── order_handler.go
│   │   │   └── stamp_handler.go
│   │   └── models/              # Struct Database (GORM Models)
│   │       ├── menu.go
│   │       ├── customer.go
│   │       ├── order.go
│   │       └── stamp_log.go
│   ├── .env.example
│   ├── go.mod
│   └── README.md
│
├── frontend/
│   ├── src/
│   │   ├── assets/
│   │   ├── components/
│   │   │   ├── MenuCard.vue
│   │   │   └── CartBottomSheet.vue
│   │   ├── services/
│   │   │   └── api.ts
│   │   ├── stores/
│   │   │   └── cart.ts
│   │   ├── types/
│   │   │   └── index.ts
│   │   ├── views/
│   │   │   └── CustomerView.vue
│   │   ├── App.vue
│   │   └── main.ts
│   ├── .env
│   ├── tailwind.config.js
│   └── package.json
│
└── README.md