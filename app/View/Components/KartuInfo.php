<?php

namespace App\View\Components;

use Closure;
use Illuminate\Contracts\View\View;
use Illuminate\View\Component;

class KartuInfo extends Component
{
   
    public function __construct(public string $judul) 
    {
    }

    public function render(): View|Closure|string
    {
        return view('components.kartu-info');
    }
}
