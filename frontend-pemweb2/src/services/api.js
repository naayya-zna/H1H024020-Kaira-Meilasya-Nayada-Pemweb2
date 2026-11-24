import axios from "axios";

// Menentukan backend yang digunakan
const backendAktif = import.meta.env.VITE_BACKEND || "fiber";

// Memilih URL backend berdasarkan pengaturan .env
const baseURL =
    backendAktif === "laravel"
        ? import.meta.env.VITE_API_LARAVEL
        : import.meta.env.VITE_API_FIBER;

// Menentukan header request
const headers = {
    "Content-Type": "application/json",
    Accept: "application/json",
};

// API Key hanya digunakan oleh backend Fiber
if (backendAktif === "fiber") {
    headers["X-Api-Key"] = import.meta.env.VITE_API_KEY;
}

// Membuat konfigurasi Axios
const apiFiber = axios.create({
    baseURL,
    headers,
});

// Menambahkan token JWT ke setiap request jika tersedia
apiFiber.interceptors.request.use(
    (config) => {
        const token = localStorage.getItem("token");

        if (token) {
            config.headers.Authorization = `Bearer ${token}`;
        }

        return config;
    },
    (error) => {
        return Promise.reject(error);
    }
);

// Menangani respons error dari backend
apiFiber.interceptors.response.use(
    (response) => response,
    (error) => {
        if (error.response?.status === 401) {
            localStorage.removeItem("token");
        }

        return Promise.reject(error);
    }
);

export default apiFiber;