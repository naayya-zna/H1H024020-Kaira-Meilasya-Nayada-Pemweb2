<?php

namespace App\Http\Controllers\Api;

use App\Http\Controllers\Controller;
use App\Http\Requests\StoreMatakuliahRequest;
use App\Http\Requests\UpdateMatakuliahRequest;
use App\Http\Resources\MatakuliahResource;
use App\Models\Matakuliah;
use Illuminate\Http\JsonResponse;

class MatakuliahController extends Controller
{
    public function index()
    {
        $matakuliah = Matakuliah::query()
            ->orderBy('semester')
            ->orderBy('kode')
            ->paginate(10);

        return MatakuliahResource::collection($matakuliah);
    }

    public function store(StoreMatakuliahRequest $request): JsonResponse
    {
        $matakuliah = Matakuliah::create($request->validated());

        return response()->json([
            'sukses' => true,
            'pesan' => 'Data mata kuliah berhasil dibuat',
            'data' => new MatakuliahResource($matakuliah),
        ], 201);
    }

    public function show(Matakuliah $matakuliah): MatakuliahResource
    {
        return new MatakuliahResource($matakuliah);
    }

    public function update(
        UpdateMatakuliahRequest $request,
        Matakuliah $matakuliah
    ): MatakuliahResource {
        $matakuliah->update($request->validated());

        return new MatakuliahResource($matakuliah->refresh());
    }

    public function destroy(Matakuliah $matakuliah): JsonResponse
    {
        $matakuliah->delete();

        return response()->json([
            'sukses' => true,
            'pesan' => 'Data mata kuliah berhasil dihapus',
        ]);
    }
}