<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;

class MatakuliahController extends Controller
{
    private function getDaftarMatakuliah()
    {
        return [
            ['kode' => 'IF101', 'nama' => 'Pemrograman Web II', 'sks' => 3],
            ['kode' => 'IF102', 'nama' => 'Algoritma & Struktur Data', 'sks' => 4],
            ['kode' => 'IF103', 'nama' => 'Etika Profesi', 'sks' => 2],
            ['kode' => 'IF104', 'nama' => 'Basis Data', 'sks' => 3],
            ['kode' => 'IF105', 'nama' => 'Bahasa Inggris Tekno', 'sks' => 2],
        ];
    }

    public function index(Request $request)
    {
        $daftarMatakuliah = $this->getDaftarMatakuliah();
        $q = $request->query('q', '');

        if (!empty($q)) {
            $daftarMatakuliah = array_filter($daftarMatakuliah, function ($mk) use ($q) {
                return stripos($mk['nama'], $q) !== false || stripos($mk['kode'], $q) !== false;
            });
        }
        return view('matakuliah.index', [
            'daftarMatakuliah' => $daftarMatakuliah,
            'q' => $q
        ]);
    }

    public function show(string $kode)
    {
        $daftarMatakuliah = $this->getDaftarMatakuliah();
        $matakuliah = collect($daftarMatakuliah)->firstWhere('kode', strtoupper($kode));
        if (!$matakuliah) {
            abort(404, 'Matakuliah tidak ditemukan');
        }
        return view('matakuliah.show', ['matakuliah' => $matakuliah]);
    }
}