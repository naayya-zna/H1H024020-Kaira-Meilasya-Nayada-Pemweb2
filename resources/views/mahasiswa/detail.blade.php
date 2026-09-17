@extends('layouts.app')

@section('judul', 'Detail Mahasiswa')

@section('konten')
<h1 class="h3 mb-3">Detail Mahasiswa</h1>

<div class="card mb-4">
    <div class="card-body">
        <h5>{{ $mahasiswa->nama }} ({{ $mahasiswa->nim }})</h5>
        <p class="mb-1">Prodi: {{ $mahasiswa->programStudi->nama }}</p>
        <p class="mb-0">IPK: {{ $mahasiswa->ipk }}</p>
    </div>
</div>

<h2 class="h5 mb-3">Mata Kuliah yang Diambil</h2>
<table class="table table-bordered bg-white">
    <thead>
        <tr>
            <th>Kode</th>
            <th>Nama Matakuliah</th>
            <th>SKS</th>
            <th>Nilai</th>
        </tr>
    </thead>
    <tbody>
        @forelse ($mahasiswa->matakuliah as $mk)
            <tr>
                <td>{{ $mk->kode }}</td>
                <td>{{ $mk->nama }}</td>
                <td>{{ $mk->sks }}</td>
                <td>{{ $mk->pivot->nilai ?? '-' }}</td>
            </tr>
        @empty
            <tr>
                <td colspan="4" class="text-center">Belum ada mata kuliah yang diambil.</td>
            </tr>
        @endforelse
    </tbody>
</table>
@endsection