<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use App\Models\Mahasiswa; 
use Illuminate\Support\Facades\DB;

class MahasiswaWebController extends Controller
{
    public function index()
    {
        DB::listen(function ($kueri) {
            logger($kueri->sql);
        });

        $daftarMahasiswa = Mahasiswa::with('programStudi')
            ->orderBy('nama')
            ->paginate(10);

        return view('mahasiswa.data', [
            'daftarMahasiswa' => $daftarMahasiswa
        ]);
    }

    public function store(Request $request) 
    { 
        $data = $request->validate([ 
            'program_studi_id' => ['required', 'exists:program_studis,id'], 
            'nim' => ['required', 'string', 'max:20', 'unique:mahasiswas,nim'], 
            'nama' => ['required', 'string', 'max:100'], 
            'email' => ['required', 'email', 'unique:mahasiswas,email'], 
            'angkatan' => ['required', 'integer', 'min:2000'], 
        ]); 

        Mahasiswa::create($data); 

        return redirect()->route('mahasiswa.data')->with('sukses', 'Data mahasiswa berhasil disimpan'); 
    }
    
    public function show($id)
    {
        $mahasiswa = Mahasiswa::with(['programStudi', 'matakuliah'])->findOrFail($id);

        return view('mahasiswa.detail', compact('mahasiswa'));
    }

    public function edit(string $id)
    {
        //
    }

    public function update(Request $request, string $id)
    {
        //
    }

    public function destroy(string $id)
    {
        //
    }
}