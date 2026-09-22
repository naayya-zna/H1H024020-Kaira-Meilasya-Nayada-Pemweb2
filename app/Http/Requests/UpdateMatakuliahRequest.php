<?php

namespace App\Http\Requests;

use Illuminate\Foundation\Http\FormRequest;
use Illuminate\Validation\Rule;

class UpdateMatakuliahRequest extends FormRequest
{
    public function authorize(): bool
    {
        return true;
    }

    public function rules(): array
    {
        $matakuliahId = $this->route('matakuliah');

        return [
            'kode' => [
                'sometimes',
                'string',
                'max:20',
                Rule::unique('matakuliahs', 'kode')->ignore($matakuliahId),
            ],
            'nama' => ['sometimes', 'string', 'max:100'],
            'sks' => ['sometimes', 'integer', 'min:1', 'max:6'],
            'semester' => ['sometimes', 'integer', 'min:1', 'max:14'],
        ];
    }
}