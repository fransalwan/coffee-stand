import { defineStore } from 'pinia';
import { ref, computed } from 'vue';
import type { MenuItem, CartItem } from '../types';

export const useCartStore = defineStore('cart', () => {
  const items = ref<CartItem[]>([]);

  const totalCount = computed(() => 
    items.value.reduce((sum, item) => sum + item.quantity, 0)
  );

  const totalPrice = computed(() => 
    items.value.reduce((sum, item) => sum + (item.item.price * item.quantity), 0)
  );

  const addItem = (menuItem: MenuItem) => {
    const existingItem = items.value.find(i => i.item.id === menuItem.id);
    if (existingItem) {
      existingItem.quantity++;
    } else {
      items.value.push({ item: menuItem, quantity: 1 });
    }
  };

  const removeItem = (menuId: number) => {
    const index = items.value.findIndex(i => i.item.id === menuId);
    if (index !== -1) {
      if (items.value[index].quantity > 1) {
        items.value[index].quantity--;
      } else {
        items.value.splice(index, 1);
      }
    }
  };

  const clearCart = () => {
    items.value = [];
  };

  // THE KILLER FEATURE: Generate WA Link
  const generateWhatsAppLink = (ownerNumber: string) => {
    if (items.value.length === 0) return '';

    let message = `Halo kak, saya mau pesan (Skip Antre) dari Web Kopi Teman:\n\n`;
    
    items.value.forEach((cartItem, index) => {
      const subtotal = cartItem.item.price * cartItem.quantity;
      message += `${index + 1}. ${cartItem.item.name} (${cartItem.quantity}x) - Rp ${subtotal.toLocaleString('id-ID')}\n`;
      if (cartItem.item.description) {
         message += `   _Catatan: ${cartItem.item.description}_\n`; // Opsional, kalau ada request khusus
      }
    });

    message += `\n*Total: Rp ${totalPrice.value.toLocaleString('id-ID')}*`;
    message += `\n\nSaya bayar pakai (Cash / QRIS di tempat). Tolong disiapkan ya kak!`;

    // Ganti 0 di awal nomor dengan 62, atau pastikan format udah 628xxx
    const cleanNumber = ownerNumber.replace(/\D/g, '').replace(/^0/, '62');
    
    return `https://wa.me/${cleanNumber}?text=${encodeURIComponent(message)}`;
  };

  return { items, totalCount, totalPrice, addItem, removeItem, clearCart, generateWhatsAppLink };
});