<?php

namespace App\Filament\Resources;

use App\Filament\Resources\MediaUploadResource\Pages;
use App\Filament\Resources\MediaUploadResource\RelationManagers;
use App\Models\MediaUpload;
use Filament\Forms;
use Filament\Resources\Form;
use Filament\Resources\Resource;
use Filament\Resources\Table;
use Filament\Tables;

class MediaUploadResource extends Resource
{
    protected static ?string $model = MediaUpload::class;

    protected static ?string $navigationIcon = 'heroicon-o-collection';

    public static function form(Form $form): Form
    {
        return $form
            ->schema([
                Forms\Components\TextInput::make('title')
                    ->label('Title'),
                Forms\Components\TextInput::make('alt')
                    ->label('Alt Image'),
                Forms\Components\TextInput::make('caption')
                    ->label('Caption'),
                Forms\Components\TextInput::make('description')
                    ->label('Description'),
                Forms\Components\SpatieMediaLibraryFileUpload::make('file')
                    ->label('File')
                    ->required(),
            ]);
    }

    public static function table(Table $table): Table
    {
        return $table
            ->columns([
                Tables\Columns\SpatieMediaLibraryImageColumn::make('file')
                    ->label('File'),
                Tables\Columns\TextColumn::make('id')->label('ID'),
                Tables\Columns\TextColumn::make('alt')
                    ->label('Alt Image')
                    ->required(),
                Tables\Columns\TextColumn::make('title')
                    ->label('Title'),
                Tables\Columns\TextColumn::make('caption')
                    ->label('Caption'),
                Tables\Columns\TextColumn::make('media')
                    ->label('URL')
                    ->formatStateUsing(function (MediaUpload $record) {
                        return $record->getFirstMediaUrl();
                    }),
            ])
            ->filters([
                //
            ])
            ->defaultSort('created_at', 'desc');
    }

    public static function getRelations(): array
    {
        return [
            //
        ];
    }

    public static function getPages(): array
    {
        return [
            'index' => Pages\ListMediaUploads::route('/'),
            'create' => Pages\CreateMediaUpload::route('/create'),
            'edit' => Pages\EditMediaUpload::route('/{record}/edit'),
        ];
    }
}
