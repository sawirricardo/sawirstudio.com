<div>
    <x-layouts.app>
        <section id="banner" class="banner" role="banner">
            <div class="p-8">
                <h1 class="text-3xl font-bold">
                    {{ $project->title }}
                </h1>
                <p class="text-xl">
                    {{ $project->excerpt }}
                </p>
            </div>
        </section>
        <section>
            <div class="p-8">
                <div class="grid grid-cols-12 gap-8">
                    <div class="col-span-2"></div>
                    <div class="prose mx-auto col-span-8">
                        {!! Str::markdown($project->content) !!}
                    </div>
                    @if (!is_null($project->meta))
                        <div class="col-span-2">
                            <h3>Meta informations</h3>
                            <ul>
                                @foreach ($project->meta as $key => $value)
                                    <li>
                                        <strong>{{ ucfirst($key) }}</strong> : {{ $value }}
                                    </li>
                                @endforeach
                            </ul>
                        </div>
                    @endif
                </div>
            </div>
        </section>
        <section id="cta">
            <div class="p-8">
                <div class="bg-gradient-to-r from-green-600 to-green-400 text-white rounded-lg shadow-xl">
                    <div class="p-8 grid md:grid-cols-2 md:items-end">
                        <div class="links-underline links-white">
                            <p class="text-xl font-bold">
                                Enough about us.
                            </p>
                        </div>
                        <h2 class="md:ml-auto text-2xl md:text-6xl md:text-right">
                            <a class="hover:text-gray-200 underline" href="{{ route('front.contact') }}">How can
                                we<br>
                                help you</a>?
                        </h2>
                    </div>
                </div>
            </div>
        </section>
    </x-layouts.app>

</div>
