# ☕ Kopi Teman Backend API

> *Backend service untuk "Stand Coffee OS" MVP. Dibangun dengan prinsip lean, pragmatic, dan fokus pada stabilitas operasional UMKM kopi skala stand.*

## 📋 Table of Contents
- [Tujuan & Filosofi](#-tujuan--filosofi)
- [Tech Stack](#-tech-stack)
- [Prerequisites](#-prerequisites)
- [Installation & Setup](#-installation--setup)
- [Database Schema](#-database-schema)
- [API Documentation](#-api-documentation)
- [Environment Variables](#-environment-variables)
- [Testing Guide](#-testing-guide)
- [Deployment](#-deployment)
- [Troubleshooting](#-troubleshooting)
- [Roadmap](#-roadmap)

---

## 🎯 Tujuan & Filosofi

Backend ini dirancang khusus untuk menyelesaikan masalah nyata di lapangan:

### Prinsip Desain
- **Anti-Ribet:** Struktur API yang flat, responsif, dan mudah di-consume oleh frontend/mobile.
- **Data Integrity:** Rancangan skema mengutamakan akurasi laporan (misal: `price_snapshot` agar histori omzet tidak berubah saat harga menu diupdate).
- **Frictionless Identity:** Identitas pelanggan berbasis `phone_number` (WA), tanpa sistem login/register yang memberatkan UX stand kopi.
- **Auto-Ready:** Saat pertama kali dijalankan, database otomatis ter-migrate & ter-seed dengan data awal. Langsung siap dipakai tanpa setup manual.
- **Auditability:** Setiap perubahan stempel dan transaksi tercatat dalam log terpisah untuk transparansi penuh.

### Problem Statement
Stand kopi kecil biasanya menghadapi 3 masalah utama:
1. **Antrean panjang** di jam sibuk → Solusi: Pre-order via WA
2. **Customer retention rendah** → Solusi: Loyalty stamp digital
3. **Cashflow tidak terprediksi** → Solusi: Dashboard omzet real-time

Backend ini adalah fondasi untuk menyelesaikan ketiga masalah tersebut.

---

## 🛠 Tech Stack

| Component | Technology | Version | Purpose |
|-----------|------------|---------|---------|
| Language | Go | 1.21+ | Concurrent, cepat, low memory footprint |
| HTTP Router | Gin | Latest | Lightweight & fast routing |
| ORM | GORM | v2 | Rapid development & database abstraction |
| Database | PostgreSQL | 14+ | Relational, ACID-compliant, solid |
| Env Management | godotenv | Latest | Load environment variables from `.env` |
| Testing | Postman | Latest | API testing & documentation |

### Kenapa Stack Ini?
- **Go + Gin:** Performa tinggi, binary size kecil (< 20MB), cocok untuk deployment di server murah.
- **GORM:** Mempercepat development MVP. Bisa migrasi ke SQLC nanti kalau butuh performa maksimal.
- **PostgreSQL:** ACID-compliant, reliable untuk data transaksi & keuangan.

---

## 📦 Prerequisites

Sebelum memulai, pastikan Anda memiliki:

### Required
- [Go 1.21+](https://go.dev/dl/) - Download dan install
- [PostgreSQL 14+](https://www.postgresql.org/download/) - Database server
- Git - Untuk version control

### Recommended Tools
- [Postman](https://www.postman.com/downloads/) atau [Insomnia](https://insomnia.rest/download) - API testing
- [DBeaver](https://dbeaver.io/download/) atau [pgAdmin](https://www.pgadmin.org/download/) - Database GUI
- [VS Code](https://code.visualstudio.com/) + Go extension - Code editor

### Verify Installation
```bash
go version          # Harus menampilkan go1.21 atau lebih baru
psql --version      # Harus menampilkan psql (PostgreSQL) 14 atau lebih baru