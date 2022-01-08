<?php

namespace App\Http\Livewire\Front\Projects;

use App\Models\Project;
use Artesaos\SEOTools\Facades\SEOTools;
use Livewire\Component;

class Show extends Component
{
    public Project $project;
    public function mount(Project $publishedProject)
    {
        $this->project = $publishedProject;
        SEOTools::setTitle($publishedProject->title);
        if (!is_null($publishedProject->excerpt)) {
            SEOTools::setDescription($publishedProject->excerpt);
        } else {
            SEOTools::setDescription($publishedProject->title);
        }
        if (!is_null($publishedProject->media_upload_id)) {
            SEOTools::addImages([$publishedProject->mediaUpload->getFirstMediaUrl()]);
        }
        visits($publishedProject)->increment();
    }

    public function render()
    {
        return view('livewire.front.projects.show');
    }
}
