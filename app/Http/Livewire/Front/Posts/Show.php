<?php

namespace App\Http\Livewire\Front\Posts;

use App\Models\Post;
use Artesaos\SEOTools\Facades\SEOTools;
use Livewire\Component;

class Show extends Component
{
    public Post $post;

    public function mount(Post $publishedPost)
    {
        $this->post = $publishedPost;
        SEOTools::setTitle($publishedPost->title);
        if (!is_null($publishedPost->excerpt)) {
            SEOTools::setDescription($publishedPost->excerpt);
        } else {
            SEOTools::setDescription($publishedPost->title);
        }
        if (!is_null($publishedPost->media_upload_id)) {
            SEOTools::addImages([$publishedPost->mediaUpload->getFirstMediaUrl()]);
        }
        visits($publishedPost)->increment();
    }

    public function render()
    {
        return view('livewire.front.posts.show');
    }
}
