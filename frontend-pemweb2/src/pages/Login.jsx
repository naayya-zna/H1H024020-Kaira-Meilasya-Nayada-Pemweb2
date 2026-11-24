
import { useState } from "react";
import { useNavigate } from "react-router-dom";
import apiFiber from "../services/api";

function Login() {
    const navigate = useNavigate();

    const [form, setForm] = useState({
        email: "",
        password: "",
    });

    const [pesan, setPesan] = useState("");
    const [loading, setLoading] = useState(false);

    const handleChange = (e) => {
        setForm({
            ...form,
            [e.target.name]: e.target.value,
        });
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        setPesan("");
        setLoading(true);

        try {
            const response = await apiFiber.post("/auth/login", {
                email: form.email,
                kata_sandi: form.password,
            });

            const token = response.data.data?.token || response.data.token;

            if (!token) {
                setPesan("Login berhasil, tetapi token tidak ditemukan.");
                return;
            }

            localStorage.setItem("token", token);
            navigate("/mahasiswa");
        } catch (error) {
            setPesan(
                error.response?.data?.pesan ||
                error.response?.data?.message ||
                "Login gagal. Periksa email dan password."
            );
        } finally {
            setLoading(false);
        }
    };

    return (
        <div className="container mt-5">
            <div className="row justify-content-center">
                <div className="col-md-5">
                    <div className="card shadow">
                        <div className="card-body p-4">
                            <h3 className="text-center mb-4">
                                Login
                            </h3>

                            {pesan && (
                                <div className="alert alert-danger">
                                    {pesan}
                                </div>
                            )}

                            <form onSubmit={handleSubmit}>
                                <div className="mb-3">
                                    <label className="form-label">
                                        Email
                                    </label>
                                    <input
                                        type="email"
                                        name="email"
                                        className="form-control"
                                        value={form.email}
                                        onChange={handleChange}
                                        required
                                    />
                                </div>

                                <div className="mb-3">
                                    <label className="form-label">
                                        Password
                                    </label>
                                    <input
                                        type="password"
                                        name="password"
                                        className="form-control"
                                        value={form.password}
                                        onChange={handleChange}
                                        required
                                    />
                                </div>

                                <button
                                    type="submit"
                                    className="btn btn-primary w-100"
                                    disabled={loading}
                                >
                                    {loading ? "Memproses..." : "Login"}
                                </button>
                            </form>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
}

export default Login;
