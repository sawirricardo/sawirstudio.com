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
        <section id="portfolio">
            <div class="p-8 grid grid-cols-1 gap-8">
                @foreach ($projects as $project)
                    <div class="grid grid-cols-1 md:grid-cols-8 gap-8 px-8">
                        <div class="sm:col-span-4 | md:col-start-2 md:col-span-3 md:-ml-32 md:mr-16">
                            <a href="https://laravelpackage.training" class="illustration is-left h-full"
                                title="Laravel package training" target="_blank" rel="nofollow noreferrer noopener">
                                <div class="relative">
                                    <div style="padding-bottom: 56.25%" class="bg-gray-50"></div>
                                    {{ optional($project->mediaUpload)->getFirstMedia() }}
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
                @endforeach
            </div>
        </section>
    </x-layouts.app>
</div>
