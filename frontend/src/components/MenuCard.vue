<script setup lang="ts">
import type { MenuItem } from '../types';
import { useCartStore } from '../stores/cart';
import { computed } from 'vue';

const props = defineProps<{ item: MenuItem }>();
const cart = useCartStore();

const cartQty = computed(() => {
  const found = cart.items.find(i => i.item.id === props.item.id);
  return found ? found.quantity : 0;
});

const formatPrice = (price: number) => {
  return `Rp ${price.toLocaleString('id-ID')}`;
};
</script>

<template>
  <div class="bg-white p-4 rounded-xl shadow-sm border border-coffee-100 flex justify-between items-start gap-4">
    <div class="flex-1">
      <h3 class="font-bold text-lg text-coffee-900">{{ item.name }}</h3>
      <p class="text-sm text-coffee-600 mt-1 italic leading-relaxed">
        "{{ item.description }}"
      </p>
      <p class="font-bold text-coffee-800 mt-3">{{ formatPrice(item.price) }}</p>
    </div>
    
    <div class="flex flex-col items-center gap-2">
      <!-- Jika belum ada di cart -->
      <button 
        v-if="cartQty === 0"
        @click="cart.addItem(item)"
        class="bg-coffee-800 text-white px-4 py-2 rounded-lg font-semibold text-sm hover:bg-coffee-900 transition active:scale-95"
      >
        + Tambah
      </button>
      
      <!-- Jika sudah ada di cart (Counter) -->
      <div v-else class="flex items-center gap-3 bg-coffee-100 rounded-lg p-1">
        <button @click="cart.removeItem(item.id)" class="w-8 h-8 flex items-center justify-center bg-white rounded-md shadow-sm font-bold text-coffee-800">-</button>
        <span class="font-bold w-4 text-center">{{ cartQty }}</span>
        <button @click="cart.addItem(item)" class="w-8 h-8 flex items-center justify-center bg-coffee-800 text-white rounded-md shadow-sm font-bold">+</button>
      </div>
    </div>
  </div>
</template>