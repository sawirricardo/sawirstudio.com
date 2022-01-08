<?php

namespace App\Filament\Resources;

use App\Filament\Resources\PostResource\Pages;
use App\Filament\Resources\PostResource\Pages\CreatePost;
use App\Filament\Resources\PostResource\RelationManagers;
use App\Models\Post;
use Filament\Forms;
use Filament\Resources\Form;
use Filament\Resources\Resource;
use Filament\Resources\Table;
use Filament\Tables;

class PostResource extends Resource
{
    protected static ?string $model = Post::class;

    protected static ?string $navigationIcon = 'heroicon-o-collection';

    public static function form(Form $form): Form
    {
        return $form
            ->schema([
                Forms\Components\TextInput::make('title')
                    ->label('Title')
                    ->required(),
                Forms\Components\TextInput::make('excerpt')->label('Description'),
                Forms\Components\DateTimePicker::make('published_at')
                    ->label('Publish Date')
                    ->afterStateHydrated(function ($set, $state, $livewire) {
                        if ($livewire instanceof CreatePost) {
                            if (is_null($state)) $set('published_at', now());
                        }
                    })
                    ->required(),
                Forms\Components\Select::make('media_upload_id')
                    ->label('Featured image')
                    ->options(\App\Models\MediaUpload::all()->pluck('id', 'id')),
                Forms\Components\SpatieTagsInput::make('tags')
                    ->label('Tags')
                    ->type('categories'),
                Forms\Components\KeyValue::make('meta')
                    ->label('Meta'),
                Forms\Components\Section::make('content')
                    ->label('Content')
                    ->schema([
                        Forms\Components\MarkdownEditor::make('content')
                            ->label('Content')
                            ->required(),
                    ]),
            ]);
    }

    public static function table(Table $table): Table
    {
        return $table
            ->columns([
                Tables\Columns\TextColumn::make('title')
                    ->label('Title'),
                Tables\Columns\SpatieTagsColumn::make('tags')
                    ->label('Tags')
                    ->type('categories'),
                Tables\Columns\TextColumn::make('published_at')
                    ->label('Published at'),
                Tables\Columns\TextColumn::make('created_at')
                    ->label('Created at'),
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
            'index' => Pages\ListPosts::route('/'),
            'create' => Pages\CreatePost::route('/create'),
            'edit' => Pages\EditPost::route('/{record}/edit'),
        ];
    }
}
