
import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import apiFiber from "../services/api";

function Mahasiswa() {
    const navigate = useNavigate();

    const [mahasiswa, setMahasiswa] = useState([]);
    const [pesan, setPesan] = useState("Memuat data mahasiswa...");
    const [loading, setLoading] = useState(true);

    const ambilMahasiswa = async () => {
        try {
            const response = await apiFiber.get("/mahasiswa");

            const data = response.data.data;

            setMahasiswa(Array.isArray(data) ? data : []);
            setPesan("");
        } catch (error) {
            setPesan(
                error.response?.data?.pesan ||
                "Gagal mengambil data mahasiswa."
            );
        } finally {
            setLoading(false);
        }
    };

    useEffect(() => {
        ambilMahasiswa();
    }, []);

    const handleLogout = () => {
            localStorage.removeItem("token");
            window.location.href = "/";
        };
        const handleHapus = async (id) => {
        const konfirmasi = window.confirm(
            "Apakah Anda yakin ingin menghapus data mahasiswa ini?"
        );

        if (!konfirmasi) {
            return;
        }

        try {
            await apiFiber.delete(`/mahasiswa/${id}`);

            alert("Data mahasiswa berhasil dihapus!");

            // Memperbarui daftar mahasiswa setelah penghapusan
            await ambilMahasiswa();
        } catch (error) {
            alert(
                error.response?.data?.pesan ||
                "Gagal menghapus data mahasiswa."
            );
        }
    };

    return (
        <div className="container mt-4">
            <div className="d-flex justify-content-between align-items-center mb-4">
                <h2>Data Mahasiswa</h2>

                <div>
                    <button
                        className="btn btn-primary me-2"
                        onClick={() => navigate("/mahasiswa/tambah")}
                    >
                        Tambah Mahasiswa
                    </button>

                    <button
                        className="btn btn-danger"
                        onClick={handleLogout}
                    >
                        Logout
                    </button>
                </div>
            </div>

            {pesan && (
                <div className="alert alert-info">
                    {pesan}
                </div>
            )}

            {!loading && mahasiswa.length > 0 && (
                <div className="table-responsive">
                    <table className="table table-bordered table-striped">
                        <thead className="table-dark">
                            <tr>
                                <th>No</th>
                                <th>NIM</th>
                                <th>Nama</th>
                                <th>IPK</th>
                                <th>Aksi</th>
                            </tr>
                        </thead>

                        <tbody>
                            {mahasiswa.map((item, index) => (
                                <tr key={item.id || index}>
                                    <td>{index + 1}</td>
                                    <td>{item.nim}</td>
                                    <td>{item.nama}</td>
                                    <td>{item.ipk}</td>
                                    <td>
                                        <button
                                            className="btn btn-danger btn-sm"
                                            onClick={() => handleHapus(item.id)}
                                        >
                                            Hapus
                                        </button>
                                    </td>
                                </tr>
                            ))}
                        </tbody>
                    </table>
                </div>
            )}

            {!loading && mahasiswa.length === 0 && !pesan && (
                <div className="alert alert-info">
                    Belum ada data mahasiswa.
                </div>
            )}
        </div>
    );
}

export default Mahasiswa;
