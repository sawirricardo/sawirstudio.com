<div>
    <x-layouts.app>
        <section id="banner" class="banner" role="banner">
            <div class="p-8">
                <h1 class="text-3xl font-bold">
                    Latest projects
                </h1>
                <p class="text-xl">
                    A collection of the latest projects from {{ config('app.name') }}
                </p>
            </div>
        </section>
        <section>
            <div class="p-8 space-y-4">
                <div class="flex space-x-4">
                    @foreach (['application', 'website', 'design'] as $category)
                        <div>
                            <input id="{{ $category }}" type="checkbox" wire:model='selectedCategories'
                                value="{{ $category }}" />
                            <label for="{{ $category }}">{{ Str::upper($category) }}</label>
                        </div>
                    @endforeach
                </div>
                <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
                    @forelse ($projects as $project)
                        <div>
                            <a href="{{ route('front.projects.show', ['publishedProject' => $project]) }}">
                                <div class="bg-white rounded shadow overflow-hidden">
                                    @if ($project->mediaUploads->count())
                                        {{ $project->mediaUpload->getFirstMedia()->img()->attributes([
                                                'class' => 'w-full h-64 object-cover',
                                                'alt' => $project->mediaUpload->alt,
                                            ]) }}
                                    @endif
                                    <div class="p-6">
                                        <h3 class="text-xl font-bold">{{ $project->title }}</h3>
                                        <p class="text-sm">{{ Str::limit($project->excerpt, 100) }}</p>
                                    </div>
                                </div>
                            </a>
                        </div>
                    @empty
                        <div class="col-span-3 text-center">No Projects (yet)</div>
                    @endforelse
                </div>
            </div>
        </section>
    </x-layouts.app>
</div>
