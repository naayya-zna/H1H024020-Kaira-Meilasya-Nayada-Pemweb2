<?php

namespace App\Http\Controllers\Api;

use App\Http\Controllers\Controller;
use App\Http\Resources\MahasiswaResource;
use App\Models\ProgramStudi;

class ProgramStudiController extends Controller
{
    public function mahasiswa($id)
    {
        $programStudi = ProgramStudi::findOrFail($id);

        $mahasiswa = $programStudi->mahasiswa()
            ->with('programStudi')
            ->paginate(10);

        return MahasiswaResource::collection($mahasiswa);
    }
}