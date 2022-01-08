<?php

namespace App\Http\Controllers\Front;

use App\Http\Controllers\Controller;
use Illuminate\Http\Request;
use Laravelium\Sitemap\Sitemap;

class SitemapController extends Controller
{
    /**
     * Display a listing of the resource.
     *
     * @return \Illuminate\Http\Response
     */
    public function index()
    {
        /**@var Sitemap */
        $sitemap = app()->make('sitemap');

        // set cache key (string), duration in minutes (Carbon|Datetime|int), turn on/off (boolean)
        // by default cache is disabled
        $sitemap->setCache('laravel.sitemap', 60);

        // check if there is cached sitemap and build new only if is not
        if (!$sitemap->isCached()) {
            $sitemap->add(
                url()->to('/'),
                \Carbon\Carbon::createFromDate('2022', '01', '08')->toISOString(),
                '1.0',
                'daily'
            );
            collect(['projects', 'posts', 'about-us', 'contact-us'])
                ->each(function ($route) use ($sitemap) {
                    $sitemap->add(
                        url()->to($route),
                        \Carbon\Carbon::createFromDate('2022', '01', '08')->toISOString(),
                        '0.9',
                        'monthly'
                    );
                });

            // // add item with translations (url, date, priority, freq, images, title, translations)
            // $translations = [
            //     ['language' => 'fr', 'url' => URL::to('pageFr')],
            //     ['language' => 'de', 'url' => URL::to('pageDe')],
            //     ['language' => 'bg', 'url' => URL::to('pageBg')],
            // ];
            // $sitemap->add(URL::to('pageEn'), '2015-06-24T14:30:00+02:00', '0.9', 'monthly', [], null, $translations);

            // add item with images
            // $images = [
            //     ['url' => URL::to('images/pic1.jpg'), 'title' => 'Image title', 'caption' => 'Image caption', 'geo_location' => 'Plovdiv, Bulgaria'],
            //     ['url' => URL::to('images/pic2.jpg'), 'title' => 'Image title2', 'caption' => 'Image caption2'],
            //     ['url' => URL::to('images/pic3.jpg'), 'title' => 'Image title3'],
            // ];
            // $sitemap->add(URL::to('post-with-images'), '2015-06-24T14:30:00+02:00', '0.9', 'monthly', $images);

            // add every post to the sitemap
            foreach (\App\Models\Post::query()->get() as $post) {
                $sitemap->add(route('front.posts.show', ['publishedPost' => $post->slug]), $post->updated_at->toString(), '0.9', 'monthly', [], $post->title);
            }

            foreach (\App\Models\Project::query()->get() as $project) {
                $sitemap->add(route('front.projects.show', ['publishedProject' => $project->slug]), $project->updated_at->toString(), '0.9', 'monthly', [], $project->title);
            }
        }

        // show your sitemap (options: 'xml' (default), 'html', 'txt', 'ror-rss', 'ror-rdf')
        return $sitemap->render('xml');
    }
}
