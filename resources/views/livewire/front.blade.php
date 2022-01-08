<div>
    <x-layouts.app class="space-y-8">
        <section id="banner" class="banner" role="banner">
            <div class="p-8 space-y-8">
                <h1 class="text-4xl md:text-8xl">
                    Solid expertise <br>in Application / Website / Digital Design
                </h1>
                <p class="text-xl md:text-2xl">
                    Crafting applications, websites, digital designs with heart.
                </p>
            </div>
        </section>
        <section>
            <div class="p-8">
                <h2 class="text-xl font-bold">Some of our clients</h2>
                <div class="grid grid-cols-1 md:grid-cols-12 md:gap-4">
                    @forelse ($clients->chunk(4) as $clientChunk)
                        <div class="md:col-span-4">
                            <ul>
                                @foreach ($clientChunk as $client)
                                    <li>
                                        <h4 class="font-bold text-gray-800 text-sm">{{ $client->name }}</h4>
                                        <p class="text-gray-600 text-xs">{{ $client->excerpt }}</p>
                                    </li>
                                @endforeach
                            </ul>
                        </div>
                    @empty
                        <div>No clients yet</div>
                    @endforelse
                </div>
            </div>
        </section>
        <section id="portfolio">
            <div class="p-8 grid grid-cols-1 gap-8">
                @forelse ($projects as $project)
                    <div class="grid grid-cols-1 md:grid-cols-8 gap-8 px-8">
                        <div class="sm:col-span-4 | md:col-start-2 md:col-span-3 md:-ml-32 md:mr-16">
                            <a href="{{ route('front.projects.show', ['publishedProject' => $project]) }}"
                                class="illustration is-left h-full" title="Laravel package training" target="_blank"
                                rel="nofollow noreferrer noopener">
                                <div class="relative">
                                    <div style="padding-bottom: 56.25%" class="bg-gray-50"></div>
                                    {{ optional($project->mediaUpload)->getFirstMedia()->img()->attributes(['class' => 'absolute inset-0 h-full w-full shadow-xl']) }}
                                </div>
                            </a>
                        </div>
                        <div class="sm:col-span-3 sm:col-start-5 | self-center">
                            <a href="{{ route('front.projects.show', ['publishedProject' => $project]) }}">
                                <div class="markup links-blue links-underline | sm:grid-text-right">
                                    <h3 class="text-xl md:text-3xl font-bold">
                                        {{ $project->title }}
                                        {{-- <span class="title-subtext text-pink-dark block">
                                        <span class="font-normal">4 hours of </span>premium video content
                                    </span> --}}
                                    </h3>
                                    <p class="text-lg">
                                        {{ $project->excerpt }}
                                    </p>
                                    {{-- <ul class="text-lg mt-8">
                                    <li>
                                        <a href="https://laravelpackage.training" target="_blank"
                                            rel="nofollow noreferrer noopener">laravelpackage.training</a>
                                    </li>
                                </ul> --}}
                                </div>
                            </a>
                        </div>
                    </div>
                @empty
                    <div>No projects yet</div>
                @endforelse
            </div>
        </section>
        <section>
            <div class="py-24 text-center">
                <a href="{{ route('front.projects.index') }}" class="text-2xl text-blue-700 hover:text-blue-500">See
                    all
                    projects</a>
            </div>
        </section>
    </x-layouts.app>
</div>
