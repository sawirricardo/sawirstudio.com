<?php

namespace App\Http\Livewire;

use Livewire\Component;

class Front extends Component
{
    public $projects;
    public function mount()
    {
        $this->projects = \App\Models\Project::query()
            ->with(['mediaUpload.media'])
            ->published()
            ->orderByDesc('published_at')
            ->take(6)
            ->get();
    }
    public function render()
    {
        return view('livewire.front');
    }
}
