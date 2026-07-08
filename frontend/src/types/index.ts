// Tipe data untuk Menu Item dari Backend
export interface MenuItem {
  id: number;
  name: string;
  description: string;
  price: number;
  is_available: boolean;
  created_at: string;
  updated_at: string;
}

// Tipe data untuk Item di Cart (Keranjang)
export interface CartItem {
  item: MenuItem;
  quantity: number;
}

// Tipe data untuk Response dari Backend
export interface ApiResponse<T> {
  data: T;
  error?: string;
}