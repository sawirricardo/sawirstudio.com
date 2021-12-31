<div>
    <x-layouts.app>
        <section id="banner" class="banner" role="banner">
            <div class="p-8">
                <h1 class="text-3xl font-bold">
                    Latest insights
                </h1>
                <p class="text-xl">
                    A collection of this journey.
                </p>
            </div>
        </section>
        <section>
            <div class="p-8">
                <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                    @forelse ($posts as $post)
                        <div>
                            <a href="{{ route('front.posts.show', ['publishedPost' => $post->slug]) }}">
                                <h3 class="font-bold text-xl">{{ $post->title }}</h3>
                                <p class="text-gray-800">{{ $post->excerpt }}</p>
                            </a>
                            <span class="text-gray-600">{{ $post->published_at->format('M d, Y') }}</span>
                        </div>
                    @empty
                        <p>No posts yet.</p>
                    @endforelse
                </div>
            </div>
        </section>
    </x-layouts.app>
</div>
