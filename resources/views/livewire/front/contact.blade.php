<div>
    <x-layouts.app>

        <section class="text-gray-700 relative">
            <div class="container px-5 py-8 mx-auto space-y-8">
                <div class="flex flex-col text-center w-full mb-4">
                    <h1 class="sm:text-3xl text-2xl font-medium title-font mb-4 text-gray-900">
                        Contact Us
                    </h1>
                    <p class="lg:w-2/3 mx-auto leading-relaxed text-base">
                        We're glad that you have found us. Let's talk.
                    </p>
                </div>
                <div class="lg:w-1/2 md:w-2/3 mx-auto space-y-4">
                    @if (session()->has('message'))
                        <div class="px-4 py-3 leading-normal text-green-700 bg-green-100 rounded-lg text-center w-full"
                            role="alert">
                            <p class="font-bold">Success</p>
                            <p>{{ session('message') }}</p>
                        </div>
                    @endif
                    <div class="flex flex-wrap -m-2">
                        <div class="p-2 w-1/2">
                            <div class="relative">
                                <label for="name" class="leading-7 text-sm text-gray-600">
                                    Name
                                </label>
                                <input wire:model.lazy='data.name' type="text" id="name" name="name"
                                    class="w-full bg-gray-100 rounded border border-gray-300 focus:border-blue-500 text-base outline-none text-gray-700 py-1 px-3 leading-8 transition-colors duration-200 ease-in-out" />
                            </div>
                        </div>
                        <div class="p-2 w-1/2">
                            <div class="relative">
                                <label for="email" class="leading-7 text-sm text-gray-600">
                                    Email
                                </label>
                                <input wire:model.lazy='data.email' type="email" id="email" name="email"
                                    class="@error('data.email')
border-red-500
                                    @enderror w-full bg-gray-100 rounded border border-gray-300 focus:border-blue-500 text-base outline-none text-gray-700 py-1 px-3 leading-8 transition-colors duration-200 ease-in-out" />
                                @error('data.email')
                                    <div class="text-red-500">{{ $message }}</div>
                                @enderror
                            </div>
                        </div>
                        <div class="p-2 w-full">
                            <div class="relative">
                                <label for="subject" class="leading-7 text-sm text-gray-600">
                                    Subject
                                </label>
                                <input wire:model.lazy='data.subject' id="subject" name="subject"
                                    class="w-full bg-gray-100 rounded border border-gray-300 focus:border-blue-500 text-base outline-none text-gray-700 py-1 px-3 resize-none leading-6 transition-colors duration-200 ease-in-out" />
                            </div>
                        </div>
                        <div class="p-2 w-full">
                            <div class="relative">
                                <label for="message" class="leading-7 text-sm text-gray-600">
                                    Message
                                </label>
                                <textarea wire:model.lazy='data.message' id="message" name="message"
                                    class="w-full bg-gray-100 rounded border border-gray-300 focus:border-blue-500 h-32 text-base outline-none text-gray-700 py-1 px-3 resize-none leading-6 transition-colors duration-200 ease-in-out"></textarea>
                            </div>
                        </div>
                        <div class="p-2 w-full">
                            <button wire:click.prevent='create'
                                class="flex mx-auto text-white bg-blue-500 border-0 py-2 px-8 focus:outline-none hover:bg-blue-600 rounded text-lg">
                                Submit
                            </button>
                        </div>
                    </div>
                </div>
            </div>
        </section>
    </x-layouts.app>
</div>
