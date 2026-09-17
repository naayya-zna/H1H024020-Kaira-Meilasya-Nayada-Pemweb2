<?php

use Illuminate\Support\Facades\Route;
use App\Http\Controllers\MahasiswaController;
use App\Http\Controllers\MatakuliahController;
use App\Http\Controllers\MahasiswaWebController;

Route::get('/', function () {
    return view('welcome');
});

Route::get('/salam', function () { 
    return 'Selamat datang di Pemrograman Web II'; 
}); 

// Rute ini dihapus/ditiadakan agar tidak bentrok dengan rute detail mahasiswa:
// Route::get('/mahasiswa/{nim}', function (string $nim) { ... });

Route::get('/semester/{angka}', function (int $angka) { 
    return 'Semester ke ' . $angka; 
})->whereNumber('angka');

Route::get('/data-mahasiswa', [MahasiswaController::class, 'index'])->name('mahasiswa.index'); 
Route::get('/data-mahasiswa/{nim}', [MahasiswaController::class, 'show'])->name('mahasiswa.show'); 

Route::get('/cari-mahasiswa', [MahasiswaController::class, 'cari']); 

Route::get('/matakuliah', [MatakuliahController::class, 'index'])->name('matakuliah.index');
Route::get('/matakuliah/{kode}', [MatakuliahController::class, 'show'])->name('matakuliah.show');

// Rute Tugas Praktikum Web
Route::get('/mahasiswa-data', [MahasiswaWebController::class, 'index'])->name('mahasiswa.data');
Route::get('/mahasiswa/{id}', [MahasiswaWebController::class, 'show'])->name('mahasiswa.detail');