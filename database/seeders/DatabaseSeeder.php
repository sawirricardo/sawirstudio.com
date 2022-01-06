<?php

namespace Database\Seeders;

use Illuminate\Database\Seeder;

class DatabaseSeeder extends Seeder
{
    /**
     * Seed the application's database.
     *
     * @return void
     */
    public function run()
    {
        \App\Models\User::factory(10)
            ->has(\App\Models\Post::factory(10))
            ->has(\App\Models\Project::factory(10))
            ->create();
        \App\Models\Client::factory(10)->create();
        \App\Models\Project::query()->get()->each(function ($project) {
            $project->attachTag(['application', 'website', 'design'][rand(0, 2)], 'project');
        });
    }
}
