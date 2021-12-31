<?php

use Illuminate\Support\Facades\Route;

/*
|--------------------------------------------------------------------------
| Web Routes
|--------------------------------------------------------------------------
|
| Here is where you can register web routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| contains the "web" middleware group. Now create something great!
|
*/

Route::name('front.')->group(function () {
    Route::get('/', \App\Http\Livewire\Front::class)->name('index');
    Route::get('/projects', \App\Http\Livewire\Front\Projects\Index::class)->name('projects.index');
    Route::get('/projects/{publishedProject:slug}', \App\Http\Livewire\Front\Projects\Show::class)->name('projects.show');
    Route::get('/posts', \App\Http\Livewire\Front\Posts\Index::class)->name('posts.index');
    Route::get('/posts/{publishedPost:slug}', \App\Http\Livewire\Front\Posts\Show::class)->name('posts.show');
    Route::get('/about-us', \App\Http\Livewire\Front\About::class)->name('about');
    Route::get('/contact-us', \App\Http\Livewire\Front\Contact::class)->name('contact');
});
