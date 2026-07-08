<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { getMenu } from './services/api';
import type { MenuItem } from './types';

const menus = ref<MenuItem[]>([]);
const isLoading = ref(true);
const error = ref('');

onMounted(async () => {
  try {
    console.log('Fetching menu from API...');
    menus.value = await getMenu();
    console.log('Menu loaded:', menus.value);
  } catch (err) {
    error.value = err instanceof Error ? err.message : 'Unknown error';
    console.error('Failed to fetch menu:', err);
  } finally {
    isLoading.value = false;
  }
});
</script>

<template>
  <div class="min-h-screen bg-coffee-50 p-8">
    <h1 class="text-3xl font-bold text-coffee-900 mb-4">API Test</h1>
    
    <div v-if="isLoading" class="text-coffee-600">
      Loading menu...
    </div>
    
    <div v-else-if="error" class="text-red-600 bg-red-50 p-4 rounded-lg">
      <strong>Error:</strong> {{ error }}
    </div>
    
    <div v-else>
      <p class="text-green-600 bg-green-50 p-4 rounded-lg mb-4">
        ✅ API Connected! Loaded {{ menus.length }} menus.
      </p>
      
      <div class="space-y-2">
        <div 
          v-for="menu in menus" 
          :key="menu.id"
          class="bg-white p-4 rounded-lg shadow-sm border border-coffee-100"
        >
          <h3 class="font-bold text-coffee-900">{{ menu.name }}</h3>
          <p class="text-sm text-coffee-600">{{ menu.description }}</p>
          <p class="text-coffee-800 font-semibold mt-2">Rp {{ menu.price.toLocaleString('id-ID') }}</p>
        </div>
      </div>
    </div>
  </div>
</template>