<?php

$db = json_decode(file_get_contents('./database/db.json'), true);

foreach ($db['projects'] as $index => $project) {
    file_put_contents(
        'public/assets/images/projects/' . $project['slug'] . '.png',
        file_get_contents($project['media_upload']['media'][0]['original_url'])
    );

    $client = array_values(array_filter($db['clients'], function ($client) use ($project) {
        return strtolower($client['name']) == strtolower($project['meta']['client']);
    }));

    if (! array_key_exists('Website URL', $project['meta'])) {
        $project['meta']['Website URL'] = 'null';
    }

    if (! array_key_exists('Google Playstore URL', $project['meta'])) {
        $project['meta']['Google Playstore URL'] = 'null';
    }

    if (! array_key_exists('Partner', $project['meta'])) {
        $project['meta']['Partner'] = 'null';
    }

    $categories = join('\n', array_map(function ($category) {
        return "- {$category['id']}";
    }, $project['tags']));

    file_put_contents(
        "pages/projects/{$project['slug']}.mdx",
        "
---
title: {$project['title']}
published_at: {$project['published_at']}
featured_media: /assets/images/projects/{$project['slug']}.png
client_id: {$client[0]['id']}
excerpt: {$project['excerpt']}
categories:
  {$categories}
meta:
  client: {$project['meta']['client']}
  partner: {$project['meta']['Partner']}
  website_url: {$project['meta']['Website URL']}
  google_playstore_url: {$project['meta']['Google Playstore URL']}
---
{$project['content']}"
    );
}
