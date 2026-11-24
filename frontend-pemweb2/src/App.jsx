
import { BrowserRouter, Routes, Route, Navigate } from "react-router-dom";
import Login from "./pages/Login";
import Mahasiswa from "./pages/Mahasiswa";
import TambahMahasiswa from "./pages/TambahMahasiswa";

// Melindungi halaman yang membutuhkan login
function ProtectedRoute({ children }) {
    const token = localStorage.getItem("token");

    return token ? children : <Navigate to="/" replace />;
}

function App() {
    return (
        <BrowserRouter>
            <Routes>
                <Route path="/" element={<Login />} />

                <Route
                    path="/mahasiswa"
                    element={
                        <ProtectedRoute>
                            <Mahasiswa />
                        </ProtectedRoute>
                    }
                />

                <Route
                    path="/mahasiswa/tambah"
                    element={
                        <ProtectedRoute>
                            <TambahMahasiswa />
                        </ProtectedRoute>
                    }
                />

                <Route path="*" element={<Navigate to="/" replace />} />
            </Routes>
        </BrowserRouter>
    );
}

export default App;
