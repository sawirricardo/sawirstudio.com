<?php

namespace App\Http\Livewire;

use Illuminate\Support\Collection;
use Livewire\Component;

class Front extends Component
{
    public Collection $projects;
    public Collection $clients;
    public function mount()
    {
        $this->projects = \App\Models\Project::query()
            ->with(['mediaUpload.media'])
            ->published()
            ->orderByDesc('published_at')
            ->take(6)
            ->get();

        $this->clients = \App\Models\Client::query()
            ->with(['mediaUpload.media'])
            ->get();
    }
    public function render()
    {
        return view('livewire.front');
    }
}
