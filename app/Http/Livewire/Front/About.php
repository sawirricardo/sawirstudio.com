<?php

namespace App\Http\Livewire\Front;

use Artesaos\SEOTools\Facades\SEOTools;
use Livewire\Component;

class About extends Component
{
    public function mount()
    {
        SEOTools::setTitle('About');
    }

    public function render()
    {
        return view('livewire.front.about', [
            'techStacks' => [
                ['type' => 'backend', 'items' => ['PHP', 'Laravel', 'Javascript (NodeJS)', 'Ruby on Rails', 'MySQL', 'PostgreSQL', 'MongoDB', 'WordPress']],
                ['type' => 'frontend', 'items' => ['Flutter', 'Livewire', 'AlpineJS', 'Vue.js', 'React.js', 'TailwindCSS']]
            ],
        ]);
    }
}
