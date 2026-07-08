import axios, { AxiosError } from 'axios';
import type { MenuItem, ApiResponse } from '../types';

// Konfigurasi Axios Instance
const apiClient = axios.create({
  baseURL: import.meta.env.VITE_API_URL || 'http://localhost:8080/api',
  timeout: 10000, // 10 detik timeout
  headers: {
    'Content-Type': 'application/json',
  },
});

// Error Handler yang lebih informatif
const handleError = (error: AxiosError) => {
  if (error.response) {
    // Server respond dengan status code error (4xx, 5xx)
    console.error('API Error:', error.response.data);
    throw new Error(`Server error: ${error.response.status}`);
  } else if (error.request) {
    // Request dibuat tapi nggak ada response (network error)
    console.error('Network Error:', error.request);
    throw new Error('Tidak dapat terhubung ke server. Periksa koneksi internet Anda.');
  } else {
    // Error saat setup request
    console.error('Request Error:', error.message);
    throw new Error(error.message);
  }
};

/**
 * Fetch semua menu yang available dari backend
 * Endpoint: GET /api/menu
 */
export const getMenu = async (): Promise<MenuItem[]> => {
  try {
    const response = await apiClient.get<ApiResponse<MenuItem[]>>('/menu');
    return response.data.data;
  } catch (error) {
    handleError(error as AxiosError);
    throw error;
  }
};

/**
 * Fetch semua menu (termasuk yang sold out) - untuk admin
 * Endpoint: GET /api/admin/menu
 */
export const getAdminMenu = async (): Promise<MenuItem[]> => {
  try {
    const response = await apiClient.get<ApiResponse<MenuItem[]>>('/admin/menu');
    return response.data.data;
  } catch (error) {
    handleError(error as AxiosError);
    throw error;
  }
};

/**
 * Create menu baru - untuk admin
 * Endpoint: POST /api/admin/menu
 */
export const createMenu = async (menuData: Omit<MenuItem, 'id' | 'created_at' | 'updated_at'>): Promise<MenuItem> => {
  try {
    const response = await apiClient.post<ApiResponse<MenuItem>>('/admin/menu', menuData);
    return response.data.data;
  } catch (error) {
    handleError(error as AxiosError);
    throw error;
  }
};

/**
 * Update menu - untuk admin
 * Endpoint: PUT /api/admin/menu/:id
 */
export const updateMenu = async (id: number, menuData: Partial<MenuItem>): Promise<MenuItem> => {
  try {
    const response = await apiClient.put<ApiResponse<MenuItem>>(`/admin/menu/${id}`, menuData);
    return response.data.data;
  } catch (error) {
    handleError(error as AxiosError);
    throw error;
  }
};

/**
 * Delete menu - untuk admin
 * Endpoint: DELETE /api/admin/menu/:id
 */
export const deleteMenu = async (id: number): Promise<void> => {
  try {
    await apiClient.delete(`/admin/menu/${id}`);
  } catch (error) {
    handleError(error as AxiosError);
    throw error;
  }
};