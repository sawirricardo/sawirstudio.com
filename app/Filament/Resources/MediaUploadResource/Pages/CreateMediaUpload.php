<?php

namespace App\Filament\Resources\MediaUploadResource\Pages;

use App\Filament\Resources\MediaUploadResource;
use Filament\Resources\Pages\CreateRecord;

class CreateMediaUpload extends CreateRecord
{
    protected static string $resource = MediaUploadResource::class;

    protected function mutateFormDataBeforeCreate(array $data): array
    {
        $data['user_id'] = auth()->id();
        return $data;
    }
}
