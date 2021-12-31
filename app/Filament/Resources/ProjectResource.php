<?php

namespace App\Filament\Resources;

use App\Filament\Resources\ProjectResource\Pages;
use App\Filament\Resources\ProjectResource\RelationManagers;
use App\Models\Project;
use Filament\Forms;
use Filament\Resources\Form;
use Filament\Resources\Resource;
use Filament\Resources\Table;
use Filament\Tables;

class ProjectResource extends Resource
{
    protected static ?string $model = Project::class;

    protected static ?string $navigationIcon = 'heroicon-o-collection';

    public static function form(Form $form): Form
    {
        return $form
            ->schema([
                Forms\Components\TextInput::make('title')
                    ->label('Title')
                    ->required(),
                Forms\Components\TextInput::make('excerpt')->label('Description'),
                Forms\Components\Select::make('media_upload_id')
                    ->label('Featured image')
                    ->options(\App\Models\MediaUpload::all()->pluck('id', 'id')),
                Forms\Components\SpatieTagsInput::make('tags')
                    ->label('Tags')
                    ->type('categories'),
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
            ])
            ->filters([
                //
            ]);
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
            'index' => Pages\ListProjects::route('/'),
            'create' => Pages\CreateProject::route('/create'),
            'edit' => Pages\EditProject::route('/{record}/edit'),
        ];
    }
}
