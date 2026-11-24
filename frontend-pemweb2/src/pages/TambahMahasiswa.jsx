
import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import apiFiber from "../services/api";

function TambahMahasiswa() {
    const navigate = useNavigate();

    const [programStudi, setProgramStudi] = useState([]);
    const [form, setForm] = useState({
        program_studi_id: "",
        nim: "",
        nama: "",
        email: "",
        angkatan: "",
        ipk: "",
        aktif: true,
    });

    const [pesan, setPesan] = useState("");
    const [galat, setGalat] = useState({});
    const [loading, setLoading] = useState(false);

    useEffect(() => {
        const ambilProgramStudi = async () => {
            try {
                const response = await apiFiber.get("/program-studi");
                setProgramStudi(response.data.data || []);
            } catch (error) {
                setPesan(
                    error.response?.data?.pesan ||
                    "Gagal mengambil data program studi."
                );
            }
        };

        ambilProgramStudi();
    }, []);

    const handleChange = (e) => {
        const { name, value, type, checked } = e.target;

        setForm({
            ...form,
            [name]: type === "checkbox" ? checked : value,
        });
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        setPesan("");
        setGalat({});
        setLoading(true);

        try {
            await apiFiber.post("/mahasiswa", {
                ...form,
                program_studi_id: Number(form.program_studi_id),
                angkatan: Number(form.angkatan),
                ipk: Number(form.ipk),
            });

            alert("Mahasiswa berhasil ditambahkan!");
            navigate("/mahasiswa");
        } catch (error) {
            const response = error.response?.data;

            setPesan(
                response?.pesan ||
                "Gagal menambahkan mahasiswa."
            );

            setGalat(response?.galat || {});
        } finally {
            setLoading(false);
        }
    };

    return (
        <div className="container mt-4">
            <h2 className="mb-4">Tambah Mahasiswa</h2>

            {pesan && (
                <div className="alert alert-danger">
                    {pesan}
                </div>
            )}

            <div className="card shadow-sm">
                <div className="card-body">
                    <form onSubmit={handleSubmit}>
                        <div className="mb-3">
                            <label className="form-label">
                                Program Studi
                            </label>
                            <select
                                name="program_studi_id"
                                className="form-select"
                                value={form.program_studi_id}
                                onChange={handleChange}
                                required
                            >
                                <option value="">
                                    Pilih Program Studi
                                </option>

                                {programStudi.map((item) => (
                                    <option key={item.id} value={item.id}>
                                        {item.nama}
                                    </option>
                                ))}
                            </select>
                        </div>

                        {[
                            ["nim", "NIM", "text"],
                            ["nama", "Nama Lengkap", "text"],
                            ["email", "Email", "email"],
                            ["angkatan", "Angkatan", "number"],
                            ["ipk", "IPK", "number"],
                        ].map(([name, label, type]) => (
                            <div className="mb-3" key={name}>
                                <label className="form-label">
                                    {label}
                                </label>

                                <input
                                    type={type}
                                    name={name}
                                    className="form-control"
                                    value={form[name]}
                                    onChange={handleChange}
                                    step={name === "ipk" ? "0.01" : undefined}
                                    required
                                />

                                {galat[name] && (
                                    <div className="text-danger small mt-1">
                                        {Array.isArray(galat[name])
                                            ? galat[name].join(", ")
                                            : galat[name]}
                                    </div>
                                )}
                            </div>
                        ))}

                        <div className="form-check mb-3">
                            <input
                                type="checkbox"
                                name="aktif"
                                className="form-check-input"
                                checked={form.aktif}
                                onChange={handleChange}
                                id="aktif"
                            />

                            <label
                                className="form-check-label"
                                htmlFor="aktif"
                            >
                                Mahasiswa Aktif
                            </label>
                        </div>

                        <button
                            type="submit"
                            className="btn btn-primary me-2"
                            disabled={loading}
                        >
                            {loading ? "Menyimpan..." : "Simpan"}
                        </button>

                        <button
                            type="button"
                            className="btn btn-secondary"
                            onClick={() => navigate("/mahasiswa")}
                        >
                            Kembali
                        </button>
                    </form>
                </div>
            </div>
        </div>
    );
}

export default TambahMahasiswa;
