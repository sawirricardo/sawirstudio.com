<div>
    <x-layouts.app>
        <section id="banner" class="banner" role="banner">
            <div class="p-8 space-y-8">
                <h1 class="text-4xl md:text-8xl">
                    Glad you found us
                </h1>
                <p class="text-xl md:text-2xl">
                    {{ config('app.name') }} helps you to achieve your business goals with applications, websites, and
                    designs.
                </p>
            </div>
        </section>
        <section>
            <div class="p-8 space-y-8">
                <div class="grid grid-cols-1 md:grid-cols-12 gap-8">
                    <div class="hidden md:col-span-2"></div>
                    <div class="prose mx-auto md:col-span-8">
                        <p>We love to work in partnership: as advisors, as architects of technically challenging
                            projects.</p>
                        <p>We'd really like to know what kind of project or problem you're dealing with. Feed us with as
                            much input as you have, so we can get accurate early on.</p>
                        <h3>Our technology stacks</h3>
                        @foreach ($techStacks as $techStack)
                            <h4>
                                {{ strtoupper($techStack['type']) }}
                                </h5>
                                <ul>
                                    @foreach ($techStack['items'] as $item)
                                        <li>
                                            {{ $item }}
                                        </li>
                                    @endforeach
                                </ul>
                        @endforeach
                    </div>
                </div>
            </div>
        </section>
        <section id="cta">
            <div class="p-8">
                <div class="bg-gradient-to-r from-green-600 to-green-400 text-white rounded-lg shadow-xl">
                    <div class="p-8 grid md:grid-cols-2 md:items-end">
                        <div class="links-underline links-white">
                            <p class="text-xl font-bold">
                                TOPICS TO DISCUSS
                            </p>
                            <ul>
                                <li>Your main objective</li>
                                <li>The initial budget & planning</li>
                                <li>Who's doing content creation?</li>
                                <li>All external services / APIs</li>
                            </ul>
                        </div>
                        <h2 class="md:ml-auto text-2xl md:text-6xl md:text-right">
                            <a class="hover:text-gray-200 underline" href="{{ route('front.contact') }}">Brief
                                us</a> your project.
                        </h2>
                    </div>
                </div>
            </div>
        </section>
    </x-layouts.app>
</div>
