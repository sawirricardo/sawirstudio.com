<?php

namespace App\Http\Livewire\Front\Posts;

use Artesaos\SEOTools\Facades\SEOTools;
use Livewire\Component;

class Index extends Component
{
    public function mount()
    {
        SEOTools::setTitle('Blog');
    }

    public function render()
    {
        return view('livewire.front.posts.index', [
            'posts' => \App\Models\Post::query()
                ->with(['user'])
                ->published()
                ->orderByDesc('published_at')
                ->get(),
        ]);
    }
}
