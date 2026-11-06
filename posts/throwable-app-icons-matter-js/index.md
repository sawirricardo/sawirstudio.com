---
title: Throwable App Icons with Matter.js | SawirStudio
url: https://sawirstudio.com/posts/throwable-app-icons-matter-js/
description: How we built the box of app icons you can throw on our home page: Matter.js for physics, DOM elements instead of a canvas, a shadow that stays put while icons spin, and fallbacks.
---

# The Throwable App Icons on Our Home Page

Our [home page](https://sawirstudio.com/) has a box with our app icons in it. They drop in when you arrive, and you can pick them up and throw them. Tap one without dragging and it opens that app's page. It took an afternoon, and a few of the details were more interesting than we expected.

## Why a box

The first idea was to let the icons fly around the whole top of the page. But our headline and buttons are there, and icons landing on top of “Start a project” is a bad trade for a bit of fun. A box beside the intro, with a label on its edge, keeps the play contained. It also fit a gap that was already there: the empty space to the right of the intro and buttons on a wide screen. On phones it sits under the buttons.

## Real physics, from Matter.js

Circles bumping into each other are easy to write yourself. But app icons are rounded squares, and a square that rolls like a coin looks wrong straight away. Proper tumbling and stacking needs a real physics engine, so we used [Matter.js](https://brm.io/matter-js/). Each icon is a rectangle with rounded corners (`chamfer`) matching the icon's 22% corner radius, inside four static walls.

The engine is 83KB, so it loads only on the home page, only after the page is up, and only when the visitor hasn't asked for reduced motion.

## DOM elements, not a canvas

Matter.js can draw to a canvas, but we only use it for the maths. The icons are ordinary links with images, positioned every frame with a CSS transform:

```
st.setProperty("--x", (b.position.x - size / 2) + "px");
st.setProperty("--y", (b.position.y - size / 2) + "px");
st.setProperty("--a", b.angle + "rad");
```

So a tap is a real link, the images are crisp at any zoom, and the icons keep the same border and style as everywhere else on the site.

## The shadow that stays put

Our style uses a hard shadow down and to the right of everything. Put that on an element that rotates and the shadow rotates with it, so a spinning icon's shadow swings around like a clock hand. The fix is to rotate the shadow's offset the opposite way every frame:

```
// keep the shadow at (3px, 3px) in screen space while the icon turns by a
st.setProperty("--shadow-x", 3 * (Math.cos(a) + Math.sin(a)) + "px");
st.setProperty("--shadow-y", 3 * (Math.cos(a) - Math.sin(a)) + "px");
```

## Grabbing and throwing

On pointer down we attach a spring (a Matter.js constraint) between the pointer and the exact spot on the icon you grabbed, not its centre, so it swings naturally when you hold it by a corner. Let go and it keeps its speed, capped so a hard flick can't send it through a wall. If the pointer moved less than six pixels, it counts as a tap and the link opens.

The box's top is open while the icons drop in from above, then closes once they're all inside, so a throw can't lose one off the top.

## Little things

- A hard landing sends a ripple through the dotted background of the page.
- The simulation sleeps when everything is still and pauses when the box is off screen, so it costs nothing while you read.
- The icons only drop in once the box is on screen. On a phone, that's when you scroll to it.
- Without JavaScript, or with reduced motion on, the icons sit still in a tidy heap, and they're still links.

It's a small thing, but it says more about the apps we make than a sentence would. If you want something like this for your own site, [tell us](https://sawirstudio.com/contact/).

**Keep reading:** [What It Took to Make Our Sites Agent-Ready](https://sawirstudio.com/posts/making-our-sites-agent-ready/), or [WordPress 7.2 Lands December 8: What We're Testing in Our Plugins](https://sawirstudio.com/posts/wordpress-7-2-what-plugin-authors-should-test/).
