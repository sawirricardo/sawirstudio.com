<div class="flex flex-col min-h-screen {{ $attributes->merge() }}">
    <header class="pt-8 flex-none z-10 | md:bg-white md:shadow-light md:py-8 | print:bg-transparent print:shadow-none">
        <div class="px-8 leading-loose | md:leading-none md:flex md:items-stretch">
            <a class="flex-shrink-0 logo h-8 w-20 mr-16 mb-8 block | md:mb-0 md:w-48 md:h-auto" href="/" title="Home">
                <span class="text-3xl md:text-6xl font-bold">SawirStudio</span>
            </a>
            <div
                class="grid grid-cols-2 items-start col-gap-8 | md:grid-cols-1 md:row-gap-6 md:justify-end md:justify-items-end md:ml-auto">
                <nav class="flex links links-black | md:row-start-2">
                    <ul class=" md:grid md:grid-flow-col md:gap-12 md:justify-between | print:hidden">
                        <li><a href="/">Home</a></li>
                        <li><a href="/projects">Projects</a></li>
                        <li><a href="/posts">Blog</a></li>
                        <li><a href="/about-us">About us</a></li>
                        <li><a href="/contact-us">Contact Us</a></li>
                    </ul>

                </nav>
                {{-- <nav
                    class="grid links links-black | md:opacity-75 md:row-start-1 md:grid-flow-col md:items-center md:gap-6 md:text-xs | print:hidden">
                    <a href="https://spatie.be/about-us">About</a>
                    <a href="https://spatie.be/blog">Blog</a>
                    <a href="https://spatie.be/docs">Docs</a>
                    <a href="https://spatie.be/guidelines">Guidelines</a>
                </nav> --}}
            </div>
        </div>
        <div class="wrap | md:hidden | print:hidden">
            <hr class="mt-8 h-2px opacity-25 rounded text-gray">
        </div>
    </header>
    {{ $slot }}
    <div class="flex-grow"></div>
    <footer class="bg-gray-50 gradient shadow-inner-light | print:shadow-none print:bg-transparent print:gradient-none"
        style="--gradient-angle: 120deg; --gradient-from:#f3efea; --gradient-to:#e1ded9;">
        <div class="flex-none pt-16 pb-8 | print:pb-2" role="navigation">
            <div class="px-8 links links-gray text-gray leading-loose | md:leading-normal">
                <div class="grid grid-cols-2 items-start text-sm | md:flex md:justify-between">
                    {{-- <ul class=" md:grid md:grid-flow-col md:gap-12 md:justify-between | print:hidden">
                        <li><a href="https://spatie.be/products">Products</a></li>
                        <li><a href="https://spatie.be/open-source">Open Source</a></li>
                        <li><a href="https://spatie.be/videos">Videos</a></li>
                        <li><a href="https://spatie.be/web-development">Web Development</a></li>
                    </ul> --}}


                    <div class="grid | md:grid-flow-col ml:items-center md:gap-12">
                        <a href="/about-us">About</a>
                        <a href="/posts">Blog</a>
                        <a href="/projects">Projects</a>
                        <a href="/contact-us">Contact Us</a>

                    </div>
                </div>
                <hr class="my-8 h-2px text-gray opacity-25 rounded | print:text-black" style="page-break-after: avoid;">
                <div class="grid gap-4 text-sm | sm:grid-cols-2 sm:gap-8 | md:flex flex-row-reverse justify-between">
                    <address class="grid gap-4 | sm:gap-0 | md:grid-flow-col md:gap-8 md:text-right">
                        <div>
                            {{-- <a class="group flex items-end | md:flex-row-reverse"
                                href="https://goo.gl/maps/A2zoLK3nVF9V8jydA" target="_blank"
                                rel="nofollow noreferrer noopener">
                                <span>
                                    Kruikstraat 22, Box 12
                                    <br>
                                    2018 Antwerp, Belgium
                                </span>
                                <span
                                    class="icon mb-px px-1 fill-current text-gray-lighter group-hover:text-pink transition-all transition-100 | print:hidden">
                                    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 384 512">
                                        <path
                                            d="M172.268 501.67C26.97 291.031 0 269.413 0 192 0 85.961 85.961 0 192 0s192 85.961 192 192c0 77.413-26.97 99.031-172.268 309.67-9.535 13.774-29.93 13.773-39.464 0zM192 272c44.183 0 80-35.817 80-80s-35.817-80-80-80-80 35.817-80 80 35.817 80 80 80z">
                                        </path>
                                    </svg>
                                    <!--
    Font Awesome Pro 5.2.0 by @fontawesome - https://fontawesome.com
    License - https://fontawesome.com/license (Commercial License)
    -->
                                </span>
                            </a> --}}
                        </div>
                        {{-- <div>
                            <a href="mailto:info@spatie.be">info@spatie.be</a>
                            <br>
                            <a href="#tel">+32 3 292 56 79</a>
                        </div> --}}
                    </address>
                    <ul class="hidden | md:grid md:grid-flow-col md:gap-6 | print:hidden">
                        <li>
                            <a href="https://github.com/sawirricardo" target="_blank"
                                rel="nofollow noreferrer noopener">
                                GitHub
                            </a>
                        </li>
                        <li>
                            <a href="https://www.instagram.com/rsawir" target="_blank"
                                rel="nofollow noreferrer noopener">
                                Instagram
                            </a>
                        </li>
                        <li>
                            <a href="https://twitter.com/RicardoSawir" target="_blank"
                                rel="nofollow noreferrer noopener">
                                Twitter
                            </a>
                        </li>
                        {{-- <li>
                            <a href="https://www.youtube.com/channel/UCoBbei3S9JLTcS2VeEOWDeQ" target="_blank"
                                rel="nofollow noreferrer noopener">
                                YouTube
                            </a>
                        </li> --}}
                    </ul>
                </div>
            </div>
        </div>

        {{-- <div class="px-8">
            <ul
                class="grid md:grid-flow-col justify-start links links-gray text-xs py-4 opacity-50 | md:justify-end md:gap-6 | print:hidden">
                <li><a href="https://spatie.be/privacy">Privacy</a></li>
                <li><a href="https://spatie.be/disclaimer">Disclaimer</a></li>
            </ul>
        </div> --}}
    </footer>
</div>
