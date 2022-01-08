<?php

namespace App\Models;

use Cviebrock\EloquentSluggable\Sluggable;
use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class Client extends Model
{
    use HasFactory;
    use Sluggable;

    protected $guarded = [];

    public function sluggable(): array
    {
        return ['slug' => ['source' => 'name']];
    }

    public function mediaUpload()
    {
        return $this->belongsTo(MediaUpload::class);
    }
}
