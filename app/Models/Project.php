<?php

namespace App\Models;

use Cviebrock\EloquentSluggable\Sluggable;
use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\SoftDeletes;
use Spatie\Tags\HasTags;

class Project extends Model
{
    use HasFactory;
    use Sluggable;
    use HasTags;
    use SoftDeletes;
    protected $guarded = [];
    public function sluggable(): array
    {
        return ['slug' => ['source' => 'title']];
    }

    public function mediaUpload()
    {
        return $this->belongsTo(MediaUpload::class);
    }
}
