<?php

namespace App\Http\Livewire\Front\Projects;

use Artesaos\SEOTools\Facades\SEOTools;
use Livewire\Component;

class Index extends Component
{
    public $selectedCategories = [];

    public function mount()
    {
        SEOTools::setTitle('Projects');
    }

    public function render()
    {
        return view('livewire.front.projects.index', [
            'projects' => \App\Models\Project::query()
                ->with(['user', 'mediaUpload.media'])
                ->published()
                ->when(!empty($this->selectedCategories), function ($query) {
                    return $query->withAnyTags($this->selectedCategories, 'project');
                })
                ->orderByDesc('published_at')
                ->get(),
        ]);
    }
}
