<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { getMenu } from '../services/api';
import type { MenuItem } from '../types';
import MenuCard from '../components/MenuCard.vue';
import CartBottomSheet from '../components/CartBottomSheet.vue';

const menus = ref<MenuItem[]>([]);
const isLoading = ref(true);
const error = ref('');

onMounted(async () => {
  try {
    menus.value = await getMenu();
  } catch (err) {
    error.value = 'Gagal memuat menu. Pastikan koneksi stabil.';
    console.error(err);
  } finally {
    isLoading.value = false;
  }
});
</script>

<template>
  <div class="min-h-screen pb-24">
    <!-- Header -->
    <header class="bg-coffee-800 text-white p-6 sticky top-0 z-40 shadow-md">
      <div class="max-w-md mx-auto">
        <h1 class="text-2xl font-bold tracking-tight">Kopi Teman</h1>
        <p class="text-coffee-200 text-sm mt-1">Seduhan jujur, bikin hari tenang.</p>
      </div>
    </header>

    <!-- Content -->
    <main class="max-w-md mx-auto p-4">
      <div v-if="isLoading" class="text-center py-10 text-coffee-600">
        Menyiapkan menu...
      </div>
      
      <div v-else-if="error" class="text-center py-10 text-red-500">
        {{ error }}
      </div>

      <div v-else class="space-y-4">
        <div class="mb-6">
          <h2 class="text-lg font-bold text-coffee-900 mb-1">Menu Hari Ini</h2>
          <p class="text-sm text-coffee-600">Pilih pesananmu, langsung kami siapkan via WhatsApp.</p>
        </div>

        <MenuCard v-for="menu in menus" :key="menu.id" :item="menu" />
      </div>
    </main>

    <!-- Floating Cart -->
    <CartBottomSheet />
  </div>
</template>