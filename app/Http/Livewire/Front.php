<?php

namespace App\Http\Livewire;

use Artesaos\SEOTools\Facades\SEOTools;
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
            ->orderBy('name')
            ->get();

        SEOTools::setDescription('Crafting applications, websites, digital designs with heart.');
    }
    public function render()
    {
        return view('livewire.front');
    }
}
