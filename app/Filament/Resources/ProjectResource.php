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
                Forms\Components\DateTimePicker::make('published_at')
                    ->label('Publish Date')
                    ->required(),
                Forms\Components\Select::make('media_upload_id')
                    ->label('Featured image')
                    ->options(\App\Models\MediaUpload::all()->pluck('id', 'id')),
                Forms\Components\SpatieTagsInput::make('tags')
                    ->label('Tags')
                    ->type('project'),
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
                    ->label('Title')
                    ->sortable(),
                Tables\Columns\SpatieTagsColumn::make('tags')
                    ->label('Category')
                    ->type('project'),
                Tables\Columns\TextColumn::make('published_at')
                    ->label('Published at')
                    ->sortable(),
                Tables\Columns\TextColumn::make('created_at')
                    ->label('Created at'),
            ])
            ->filters([
                Tables\Filters\MultiSelectFilter::make('tags')
                    ->label('Category')
                    ->options(\Spatie\Tags\Tag::all()->pluck('name', 'name'))
                    ->query(function ($query, $value) {
                        if (!empty($value['values'])) $query->withAnyTags($value['values'], 'project');
                    }),
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
            'index' => Pages\ListProjects::route('/'),
            'create' => Pages\CreateProject::route('/create'),
            'edit' => Pages\EditProject::route('/{record}/edit'),
        ];
    }
}
