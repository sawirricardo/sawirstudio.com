// Dark mode toggle. The initial theme is set by an inline script in <head>.
(function () {
  const root = document.documentElement;
  const systemDark = matchMedia("(prefers-color-scheme: dark)");

  function savedTheme() {
    try {
      return localStorage.getItem("theme");
    } catch (e) {
      return null;
    }
  }

  // The phone's status bar and browser bar take the page's background, in the theme shown.
  function syncThemeColor() {
    const color = root.dataset.theme === "dark" ? "#1a1916" : "#fffcf5";
    document
      .querySelectorAll('meta[name="theme-color"]')
      .forEach((meta) => (meta.content = color));
  }
  syncThemeColor();

  document.querySelector(".theme-toggle").addEventListener("click", () => {
    const theme = root.dataset.theme === "dark" ? "light" : "dark";
    const apply = () => {
      root.dataset.theme = theme;
      syncThemeColor();
    };
    // Cross-fade where the browser can, so the whole page doesn't flash to the other theme.
    if (
      document.startViewTransition &&
      !matchMedia("(prefers-reduced-motion: reduce)").matches
    ) {
      document.startViewTransition(apply);
    } else {
      apply();
    }
    try {
      localStorage.setItem("theme", theme);
    } catch (e) {}
  });

  // Follow system changes until the visitor picks a theme themselves.
  systemDark.addEventListener("change", (e) => {
    if (!savedTheme()) root.dataset.theme = e.matches ? "dark" : "light";
    syncThemeColor();
  });
})();

// The pinned desktop header draws its bottom edge once the page scrolls under it.
(function () {
  const header = document.querySelector(".site-header");
  let queued = false;
  function update() {
    header.classList.toggle("is-stuck", scrollY > 0);
    queued = false;
  }
  addEventListener(
    "scroll",
    () => {
      if (!queued) {
        queued = true;
        requestAnimationFrame(update);
      }
    },
    { passive: true }
  );
  update();
})();

// Page transitions run only between a project card and its page (both ways);
// every other link loads without one. Browsers without them ignore all this.
(function () {
  // Translated pages live under /<lang>/ (like /zh-hans/projects/); compare without it.
  const bare = (path) => path.replace(/^\/[a-z]{2}(-[a-z]{4})?(?=\/)/, "");
  const isList = (path) => bare(path) === "/" || bare(path) === "/projects/";
  const isProject = (path) => /^\/projects\/[^/]+\/$/.test(bare(path));

  addEventListener("pageswap", (event) => {
    if (!event.viewTransition) return;
    const from = event.activation && event.activation.from;
    const to = event.activation && event.activation.entry;
    if (!from || !to || !from.url || !to.url) return event.viewTransition.skipTransition();
    const a = new URL(from.url).pathname;
    const b = new URL(to.url).pathname;
    if (!((isList(a) && isProject(b)) || (isProject(a) && isList(b)))) {
      event.viewTransition.skipTransition();
    }
  });
})();

// Projects page: search + category filter
(function () {
  const list = document.getElementById("project-list");
  if (!list) return;

  const search = document.getElementById("search");
  const checkboxes = document.querySelectorAll(".filters input[type=checkbox]");
  const empty = document.getElementById("no-projects");
  const cards = list.querySelectorAll("[data-categories]");

  function applyFilters() {
    const query = search.value.trim().toLowerCase();
    const selected = Array.from(checkboxes)
      .filter((checkbox) => checkbox.checked)
      .map((checkbox) => checkbox.value);
    let visible = 0;

    cards.forEach((card) => {
      const categories = card.dataset.categories.split(",");
      const matchesCategory =
        selected.length === 0 || categories.some((c) => selected.includes(c));
      const matchesSearch =
        query === "" || card.dataset.search.includes(query);
      const show = matchesCategory && matchesSearch;

      card.hidden = !show;
      if (show) visible++;
    });

    list.hidden = visible === 0;
    empty.hidden = visible !== 0;
  }

  search.addEventListener("input", applyFilters);
  checkboxes.forEach((checkbox) =>
    checkbox.addEventListener("change", applyFilters)
  );
  document.getElementById("clear-filters").addEventListener("click", () => {
    search.value = "";
    checkboxes.forEach((checkbox) => (checkbox.checked = false));
    applyFilters();
  });
})();

// Contact page: submit the inquiry to Google Forms without leaving the site.
// Without JS the form still posts directly and Google shows its own confirmation.
(function () {
  const form = document.getElementById("inquiry-form");
  if (!form) return;

  const button = form.querySelector("button[type=submit]");
  const error = document.getElementById("form-error");
  const success = document.getElementById("form-success");

  form.addEventListener("submit", async (event) => {
    event.preventDefault();
    button.disabled = true;
    const label = button.textContent;
    button.textContent = button.dataset.sending || "Sending…";
    error.hidden = true;

    try {
      // Google Forms doesn't send CORS headers, so the response is opaque;
      // a resolved request means it reached Google.
      await fetch(form.action, {
        method: "POST",
        mode: "no-cors",
        body: new URLSearchParams(new FormData(form)),
      });
      form.hidden = true;
      success.hidden = false;
      success.focus();
    } catch (e) {
      error.hidden = false;
      button.disabled = false;
      button.textContent = label;
    }
  });
})();

// App icons you can pick up and turn over. A hover tilts an icon toward the pointer, a drag turns
// it up to 60 degrees, and letting go springs it back. A drag never follows the link.
(function () {
  const row = document.querySelector(".app-row");
  if (!row || matchMedia("(prefers-reduced-motion: reduce)").matches) return;

  const HOVER = 16; // degrees at the icon's edge while hovering
  const DRAG = 60; // most a drag turns it
  const finePointer = matchMedia("(hover: hover) and (pointer: fine)");

  row.querySelectorAll("a").forEach((link) => {
    const img = link.querySelector("img");
    if (!img) return;
    const tilt = document.createElement("span");
    const card = document.createElement("span");
    const foil = document.createElement("span");
    tilt.className = "tilt";
    card.className = "tilt-card";
    foil.className = "foil";
    img.replaceWith(tilt);
    tilt.append(card);
    card.append(img, foil);

    const s = { rx: 0, ry: 0, vx: 0, vy: 0, tx: 0, ty: 0, scale: 1, ts: 1, frame: 0 };
    let drag = null;
    let dragged = false;

    function render() {
      card.style.setProperty("--rx", `${s.rx.toFixed(2)}deg`);
      card.style.setProperty("--ry", `${s.ry.toFixed(2)}deg`);
      card.style.setProperty("--s", s.scale.toFixed(3));
      // The light sits opposite the tilt, so the sheen slides across as the icon turns.
      foil.style.setProperty("--mx", `${50 - s.ry * 1.4}%`);
      foil.style.setProperty("--my", `${50 + s.rx * 1.4}%`);
      foil.style.setProperty("--hue", `${(s.ry * 3 - s.rx * 2).toFixed(1)}deg`);
    }

    // A damped spring toward the target, so a let-go icon overshoots a little and settles.
    function step() {
      s.vx = (s.vx + (s.tx - s.rx) * 0.12) * 0.74;
      s.vy = (s.vy + (s.ty - s.ry) * 0.12) * 0.74;
      s.rx += s.vx;
      s.ry += s.vy;
      s.scale += (s.ts - s.scale) * 0.2;
      render();
      const settled =
        Math.abs(s.tx - s.rx) < 0.05 && Math.abs(s.ty - s.ry) < 0.05 &&
        Math.abs(s.vx) < 0.05 && Math.abs(s.vy) < 0.05 && Math.abs(s.ts - s.scale) < 0.002;
      s.frame = settled ? 0 : requestAnimationFrame(step);
      if (settled && !tilt.classList.contains("is-active")) {
        card.style.removeProperty("--rx");
        card.style.removeProperty("--ry");
        card.style.removeProperty("--s");
      }
    }
    function go() {
      if (!s.frame) s.frame = requestAnimationFrame(step);
    }
    function aim(x, y, reach) {
      s.tx = Math.max(-reach, Math.min(reach, x));
      s.ty = Math.max(-reach, Math.min(reach, y));
      go();
    }
    function rest() {
      tilt.classList.remove("is-active", "is-grabbed");
      s.tx = s.ty = 0;
      s.ts = 1;
      go();
    }

    tilt.addEventListener("pointermove", (e) => {
      if (drag) {
        const dx = e.clientX - drag.x;
        const dy = e.clientY - drag.y;
        if (Math.hypot(dx, dy) > 6) dragged = true;
        aim(drag.rx - dy * 0.6, drag.ry + dx * 0.6, DRAG);
        return;
      }
      if (e.pointerType !== "mouse" || !finePointer.matches) return;
      const r = tilt.getBoundingClientRect();
      const px = (e.clientX - r.left) / r.width - 0.5;
      const py = (e.clientY - r.top) / r.height - 0.5;
      tilt.classList.add("is-active");
      s.ts = 1.08;
      aim(-py * 2 * HOVER, px * 2 * HOVER, HOVER);
    });
    tilt.addEventListener("pointerdown", (e) => {
      if (e.button !== 0) return;
      dragged = false;
      drag = { x: e.clientX, y: e.clientY, rx: s.rx, ry: s.ry };
      tilt.setPointerCapture(e.pointerId);
      tilt.classList.add("is-active", "is-grabbed");
      s.ts = 1.12;
      go();
    });
    function release(e) {
      if (!drag) return;
      drag = null;
      if (tilt.hasPointerCapture(e.pointerId)) tilt.releasePointerCapture(e.pointerId);
      rest();
    }
    tilt.addEventListener("pointerup", release);
    tilt.addEventListener("pointercancel", release);
    tilt.addEventListener("pointerleave", () => {
      if (!drag) rest();
    });
    // Turning an icon over is not a tap on it.
    link.addEventListener("click", (e) => {
      if (dragged) {
        e.preventDefault();
        dragged = false;
      }
    });
    // Links and images are draggable by default, and the browser's own drag would cancel ours.
    img.draggable = false;
    link.draggable = false;
    link.addEventListener("dragstart", (e) => e.preventDefault());
  });
  row.classList.add("tilt-ready");
})();

// Home hero: a field of paper dots, turned up loud. Dots flee the cursor or a finger, swell and
// heat from blue to yellow to coral; a click or tap sends a shockwave; and the field breathes on its
// own so it's alive before anyone touches it. Pauses off-screen; a still CSS pattern under reduced motion.
(function () {
  var canvas = document.querySelector(".dotfield");
  if (!canvas || !canvas.getContext) return;
  if (matchMedia("(prefers-reduced-motion: reduce)").matches) return;
  var hero = canvas.parentElement;
  var ctx = canvas.getContext("2d");
  var GAP = 24, RADIUS = 1.8, REACH = 240, PUSH = 70;
  var dots = [], waves = [], w = 0, h = 0, dpr = 1, t0 = performance.now();
  var pointer = { x: -9999, y: -9999, px: -9999, py: -9999, speed: 0, active: false };
  var running = false, visible = true, ink = "#1f1d1a", heat = ["#5b93ff", "#ffd84d", "#ff7a59"];

  function colors() {
    var cs = getComputedStyle(document.documentElement);
    ink = cs.getPropertyValue("--ink").trim() || ink;
    heat = [cs.getPropertyValue("--blue").trim() || heat[0], cs.getPropertyValue("--yellow").trim() || heat[1], cs.getPropertyValue("--coral").trim() || heat[2]];
  }
  function layout() {
    var r = hero.getBoundingClientRect();
    dpr = Math.min(window.devicePixelRatio || 1, 2);
    w = r.width; h = r.height;
    canvas.width = Math.round(w * dpr); canvas.height = Math.round(h * dpr);
    ctx.setTransform(dpr, 0, 0, dpr, 0, 0);
    dots = [];
    for (var y = GAP / 2; y < h; y += GAP) for (var x = GAP / 2; x < w; x += GAP) dots.push({ ox: x, oy: y, x: x, y: y, vx: 0, vy: 0, heat: 0 });
  }
  function step(now) {
    var t = (now - t0) / 1000;
    pointer.speed *= 0.9;
    var boost = 1 + Math.min(pointer.speed / 25, 1.6);
    ctx.clearRect(0, 0, w, h);
    for (var k = waves.length - 1; k >= 0; k--) { waves[k].r += 14; waves[k].life -= 0.022; if (waves[k].life <= 0) waves.splice(k, 1); }
    for (var i = 0; i < dots.length; i++) {
      var d = dots[i];
      // breathing: a slow diagonal swell across the field
      var breathe = Math.sin(t * 1.6 - (d.ox + d.oy) * 0.012) * 0.5 + 0.5;
      var tx = d.ox + Math.sin(t * 0.9 + d.oy * 0.03) * 2.5, ty = d.oy + Math.cos(t * 0.8 + d.ox * 0.03) * 2.5, hot = 0;
      if (pointer.active) {
        var dx = d.ox - pointer.x, dy = d.oy - pointer.y, dist = Math.sqrt(dx * dx + dy * dy);
        if (dist < REACH) {
          var f = 1 - dist / REACH; f = f * f * (3 - 2 * f);
          var n = dist || 1;
          tx += (dx / n) * PUSH * f * boost; ty += (dy / n) * PUSH * f * boost; hot = f;
        }
      }
      for (var k2 = 0; k2 < waves.length; k2++) {
        var wv = waves[k2], wx = d.ox - wv.x, wy = d.oy - wv.y, wd = Math.sqrt(wx * wx + wy * wy), band = 1 - Math.abs(wd - wv.r) / 70;
        if (band > 0) { var wn = wd || 1; tx += (wx / wn) * 55 * band * wv.life; ty += (wy / wn) * 55 * band * wv.life; hot = Math.max(hot, band * wv.life); }
      }
      d.vx = (d.vx + (tx - d.x) * 0.12) * 0.78;
      d.vy = (d.vy + (ty - d.y) * 0.12) * 0.78;
      d.x += d.vx; d.y += d.vy;
      d.heat += (hot - d.heat) * 0.16;
      var e = d.heat;
      ctx.globalAlpha = 0.14 + breathe * 0.08 + e * 0.8;
      ctx.fillStyle = e < 0.06 ? ink : e < 0.35 ? heat[0] : e < 0.65 ? heat[1] : heat[2];
      ctx.beginPath();
      ctx.arc(d.x, d.y, RADIUS + breathe * 0.6 + e * 7, 0, 6.2832);
      ctx.fill();
    }
    ctx.globalAlpha = 1;
    if (visible && !document.hidden) requestAnimationFrame(step); else running = false;
  }
  function kick() { if (!running && visible && !document.hidden) { running = true; requestAnimationFrame(step); } }
  function local(e) { var r = canvas.getBoundingClientRect(); return { x: e.clientX - r.left, y: e.clientY - r.top, r: r }; }
  function move(e) {
    var p = local(e);
    if (pointer.px > -9999) pointer.speed = Math.max(pointer.speed, Math.hypot(p.x - pointer.px, p.y - pointer.py));
    pointer.px = pointer.x = p.x; pointer.py = pointer.y = p.y;
    pointer.active = p.x >= 0 && p.y >= 0 && p.x <= p.r.width && p.y <= p.r.height;
    kick();
  }
  function burst(e) {
    if (e.target.closest && e.target.closest("a, button")) return;
    var p = local(e); waves.push({ x: p.x, y: p.y, r: 0, life: 1 }); if (waves.length > 4) waves.shift(); move(e);
  }
  function leave() { pointer.active = false; pointer.px = -9999; }

  colors(); layout();
  canvas.classList.add("is-live");
  hero.addEventListener("pointermove", move, { passive: true });
  hero.addEventListener("pointerdown", burst, { passive: true });
  hero.addEventListener("pointerleave", leave);
  hero.addEventListener("pointerup", function (e) { if (e.pointerType !== "mouse") leave(); });
  hero.addEventListener("pointercancel", leave);
  var resizeTimer;
  window.addEventListener("resize", function () { clearTimeout(resizeTimer); resizeTimer = setTimeout(layout, 150); });
  document.addEventListener("visibilitychange", kick);
  // An app icon landing hard in the box below sends a ripple through the dots.
  window.addEventListener("hero:thud", function (e) {
    var r = canvas.getBoundingClientRect();
    waves.push({ x: e.detail.x - r.left, y: e.detail.y - r.top, r: 0, life: e.detail.strength });
    if (waves.length > 4) waves.shift();
    kick();
  });
  new MutationObserver(colors).observe(document.documentElement, { attributes: true, attributeFilter: ["data-theme"] });
  if ("IntersectionObserver" in window) {
    new IntersectionObserver(function (entries) { visible = entries[0].isIntersecting; kick(); }).observe(hero);
  }
  kick();
})();

// Home hero: the highlights sweep in behind the headline and the doodles draw themselves, once.
(function () {
  var hero = document.querySelector(".hero");
  if (!hero || !hero.querySelector(".pile")) return;
  if (matchMedia("(prefers-reduced-motion: reduce)").matches) return;
  hero.classList.add("will-draw");
  requestAnimationFrame(function () {
    requestAnimationFrame(function () { hero.classList.add("is-drawn"); });
  });
})();

// Home hero: a box of our app icons. They drop in the first time the box is on screen, and can be
// picked up and thrown; a hard landing ripples the dot field. A tap without a drag opens the app's
// page. The physics engine (Matter.js) loads only here; without it the icons stay a still heap.
(function () {
  var pile = document.querySelector(".pile");
  if (!pile || matchMedia("(prefers-reduced-motion: reduce)").matches) return;
  pile.classList.add("is-loading");
  var script = document.createElement("script");
  script.src = "/assets/js/vendor/matter-0.20.0.min.js";
  script.onload = function () {
    try { start(window.Matter); } catch (e) { pile.classList.remove("is-loading", "is-live"); }
  };
  script.onerror = function () { pile.classList.remove("is-loading"); };
  document.head.appendChild(script);

  function start(M) {
    var engine = M.Engine.create({ enableSleeping: true, positionIterations: 10, velocityIterations: 8 });
    engine.gravity.scale = 0.002;
    var world = engine.world;
    var STEP = 1000 / 120, WALL = 400;
    var links = Array.prototype.slice.call(pile.querySelectorAll("a"));
    var items = [], walls = [], size = 0, W = 0, H = 0;
    var running = false, visible = false, dropped = false, pending = 0, last = 0, acc = 0;
    var drag = null, lastThud = 0, lidClosed = false;

    pile.classList.add("is-live");
    pile.classList.remove("is-loading");

    function measure() {
      W = pile.clientWidth;
      H = pile.clientHeight;
      return links[0].offsetWidth;
    }
    function buildWalls() {
      M.Composite.remove(world, walls);
      var o = { isStatic: true, friction: 0.6, restitution: 0.2 };
      // Open above while the icons drop in, then closed, so a throw can't lose one off the top.
      var top = lidClosed ? 0 : -size * 4, sideH = H - top + WALL * 2, sideY = (H + top) / 2;
      walls = [
        M.Bodies.rectangle(W / 2, H + WALL / 2, W + WALL * 2, WALL, o),
        M.Bodies.rectangle(W / 2, top - WALL / 2, W + WALL * 2, WALL, o),
        M.Bodies.rectangle(-WALL / 2, sideY, WALL, sideH, o),
        M.Bodies.rectangle(W + WALL / 2, sideY, WALL, sideH, o),
      ];
      M.Composite.add(world, walls);
    }

    function render() {
      for (var i = 0; i < items.length; i++) {
        var b = items[i].body, st = items[i].el.style, a = b.angle;
        var c = Math.cos(a), s = Math.sin(a);
        st.setProperty("--x", (b.position.x - size / 2).toFixed(1) + "px");
        st.setProperty("--y", (b.position.y - size / 2).toFixed(1) + "px");
        st.setProperty("--a", a.toFixed(4) + "rad");
        // The pop shadow stays down and to the right however the icon is turned.
        st.setProperty("--shadow-x", (3 * (c + s)).toFixed(2) + "px");
        st.setProperty("--shadow-y", (3 * (c - s)).toFixed(2) + "px");
      }
    }
    function frame(now) {
      if (!last) last = now;
      acc += Math.min(now - last, 50);
      last = now;
      while (acc >= STEP) {
        M.Engine.update(engine, STEP);
        acc -= STEP;
      }
      render();
      if (!lidClosed && dropped && pending === 0 && items.every(function (it) { return it.body.bounds.min.y > 0; })) {
        lidClosed = true;
        buildWalls();
      }
      var awake = drag || pending > 0 || items.some(function (it) { return !it.body.isSleeping; });
      if (awake && visible && !document.hidden) requestAnimationFrame(frame);
      else { running = false; last = 0; acc = 0; }
    }
    function kick() {
      if (!running && visible && !document.hidden) {
        running = true;
        requestAnimationFrame(frame);
      }
    }

    function drop() {
      dropped = true;
      var cols = Math.max(1, Math.floor(W / (size * 1.3)));
      links.forEach(function (el, i) {
        pending++;
        setTimeout(function () {
          var x = ((i % cols) + 0.5) * (W / cols) + (Math.random() - 0.5) * size * 0.3;
          var y = -size * 0.8 - Math.floor(i / cols) * size * 0.4;
          var body = M.Bodies.rectangle(x, y, size, size, {
            chamfer: { radius: size * 0.22 },
            restitution: 0.3,
            friction: 0.5,
            frictionStatic: 0.8,
            frictionAir: 0.012,
            density: 0.0015,
            sleepThreshold: 40,
          });
          M.Body.setAngle(body, (Math.random() - 0.5) * 0.8);
          M.Body.setAngularVelocity(body, (Math.random() - 0.5) * 0.08);
          M.Composite.add(world, body);
          items.push({ el: el, body: body });
          pending--;
          kick();
        }, 300 + i * 110);
      });
    }

    function local(e) {
      var r = pile.getBoundingClientRect();
      return {
        x: Math.max(-size, Math.min(W + size, e.clientX - r.left - pile.clientLeft)),
        y: Math.max(-size, Math.min(H + size, e.clientY - r.top - pile.clientTop)),
      };
    }
    function release(e) {
      if (!drag || e.pointerId !== drag.id) return;
      var it = drag.it;
      M.Composite.remove(world, drag.c);
      if (it.el.hasPointerCapture(e.pointerId)) it.el.releasePointerCapture(e.pointerId);
      it.el.classList.remove("is-grabbed");
      // Keep a hard throw from outrunning the walls.
      if (M.Body.getSpeed(it.body) > size * 0.45) M.Body.setSpeed(it.body, size * 0.45);
      it.thrown = drag.moved;
      if (drag.moved) pile.classList.add("was-thrown");
      drag = null;
      kick();
    }
    links.forEach(function (el) {
      el.draggable = false;
      el.addEventListener("dragstart", function (e) { e.preventDefault(); });
      el.addEventListener("pointerdown", function (e) {
        var it = items.find(function (x) { return x.el === el; });
        if (!it || drag || e.button !== 0) return;
        var p = local(e), b = it.body;
        M.Sleeping.set(b, false);
        drag = {
          it: it,
          id: e.pointerId,
          x0: e.clientX,
          y0: e.clientY,
          moved: false,
          c: M.Constraint.create({
            pointA: p,
            bodyB: b,
            pointB: { x: p.x - b.position.x, y: p.y - b.position.y },
            stiffness: 0.2,
            damping: 0.1,
            length: 0,
          }),
        };
        M.Composite.add(world, drag.c);
        el.setPointerCapture(e.pointerId);
        el.classList.add("is-grabbed");
        kick();
      });
      el.addEventListener("pointermove", function (e) {
        if (!drag || e.pointerId !== drag.id) return;
        if (Math.hypot(e.clientX - drag.x0, e.clientY - drag.y0) > 6) drag.moved = true;
        drag.c.pointA = local(e);
      });
      el.addEventListener("pointerup", release);
      el.addEventListener("pointercancel", release);
      // A throw is not a tap on the link.
      el.addEventListener("click", function (e) {
        var it = items.find(function (x) { return x.el === el; });
        if (!it || it.thrown) {
          e.preventDefault();
          if (it) it.thrown = false;
        }
      });
    });

    M.Events.on(engine, "collisionStart", function (ev) {
      var now = performance.now();
      if (now - lastThud < 140) return;
      for (var i = 0; i < ev.pairs.length; i++) {
        var pair = ev.pairs[i], va = M.Body.getVelocity(pair.bodyA), vb = M.Body.getVelocity(pair.bodyB);
        var speed = Math.hypot(va.x - vb.x, va.y - vb.y);
        if (speed < 9) continue;
        var pt = pair.collision.supports[0] || pair.bodyA.position, r = pile.getBoundingClientRect();
        lastThud = now;
        window.dispatchEvent(new CustomEvent("hero:thud", {
          detail: { x: r.left + pile.clientLeft + pt.x, y: r.top + pile.clientTop + pt.y, strength: Math.min(1, (speed - 6) / 14) },
        }));
        return;
      }
    });

    size = measure();
    buildWalls();
    if ("ResizeObserver" in window) {
      new ResizeObserver(function () {
        var next = measure();
        if (next && next !== size) {
          items.forEach(function (it) { M.Body.scale(it.body, next / size, next / size); });
          size = next;
        }
        buildWalls();
        items.forEach(function (it) {
          var p = it.body.position;
          M.Body.setPosition(it.body, {
            x: Math.max(size / 2, Math.min(W - size / 2, p.x)),
            y: Math.min(H - size / 2, p.y),
          });
          M.Sleeping.set(it.body, false);
        });
        render();
        kick();
      }).observe(pile);
    }
    document.addEventListener("visibilitychange", kick);
    var io = new IntersectionObserver(function (entries) {
      visible = entries[0].isIntersecting;
      if (visible && !dropped && entries[0].intersectionRatio >= 0.35) drop();
      kick();
    }, { threshold: [0, 0.35] });
    io.observe(pile);
  }
})();

// Signatures write themselves once, the first time they scroll into view.
(function () {
  var sigs = document.querySelectorAll(".signature");
  if (!sigs.length || !("IntersectionObserver" in window)) return;
  if (matchMedia("(prefers-reduced-motion: reduce)").matches) return;
  var io = new IntersectionObserver(function (entries) {
    entries.forEach(function (e) {
      if (!e.isIntersecting) return;
      e.target.classList.add("is-drawn");
      io.unobserve(e.target);
    });
  }, { threshold: 0.6 });
  sigs.forEach(function (s) {
    s.classList.add("will-draw");
    io.observe(s);
  });
})();

// Suggest the visitor's own language when this page exists in it. Never redirects:
// it asks once, and any choice (or a click on the footer's language list) is remembered.
(function () {
  const links = document.querySelectorAll(".lang-links a[hreflang]");
  if (!links.length || navigator.webdriver) return;
  const store = {
    get: (k) => { try { return localStorage.getItem(k); } catch (e) { return null; } },
    set: (k, v) => { try { localStorage.setItem(k, v); } catch (e) {} },
  };
  links.forEach((a) => a.addEventListener("click", () => store.set("lang", a.hreflang.toLowerCase())));
  if (store.get("lang")) return;

  const available = {};
  links.forEach((a) => (available[a.hreflang.toLowerCase()] = a));
  const match = (tag) => {
    tag = tag.toLowerCase();
    if (tag.startsWith("zh")) return /hant|-(tw|hk|mo)\b/.test(tag) ? "zh-hant" : "zh-hans";
    return tag.split("-")[0];
  };
  const wanted = (navigator.languages || [navigator.language]).map(match).find((c) => available[c]);
  const current = document.documentElement.lang.toLowerCase();
  if (!wanted || wanted === current) return;

  // Asked in the language being offered, since that's the one the visitor reads.
  const text = {
    en: ["Read this page in English?", "Switch", "No thanks"],
    "zh-hans": ["要以简体中文阅读本页吗？", "切换", "不用了"],
    "zh-hant": ["要以繁體中文閱讀本頁嗎？", "切換", "不用了"],
    ja: ["このページを日本語で表示しますか？", "切り替える", "閉じる"],
    de: ["Diese Seite auf Deutsch lesen?", "Wechseln", "Nein, danke"],
    fr: ["Lire cette page en français ?", "Changer", "Non merci"],
    nl: ["Deze pagina in het Nederlands lezen?", "Wisselen", "Nee, bedankt"],
    it: ["Leggere questa pagina in italiano?", "Cambia", "No, grazie"],
    ru: ["Открыть эту страницу на русском?", "Перейти", "Нет, спасибо"],
    hi: ["यह पेज हिंदी में पढ़ें?", "बदलें", "नहीं, धन्यवाद"],
    ar: ["هل تريد قراءة هذه الصفحة بالعربية؟", "التبديل", "لا، شكرًا"],
    he: ["לקרוא את הדף הזה בעברית?", "מעבר", "לא, תודה"],
  }[wanted];
  if (!text) return;

  const bar = document.createElement("div");
  bar.className = "lang-suggest";
  bar.lang = available[wanted].lang;
  bar.dir = ["ar", "he"].includes(wanted) ? "rtl" : "ltr";
  bar.setAttribute("role", "region");
  bar.setAttribute("aria-label", text[0]);
  const p = document.createElement("p");
  p.textContent = text[0];
  const go = document.createElement("a");
  go.className = "button";
  go.href = available[wanted].href;
  go.textContent = text[1];
  const no = document.createElement("button");
  no.type = "button";
  no.className = "lang-suggest-close";
  no.textContent = text[2];
  go.addEventListener("click", () => store.set("lang", wanted));
  no.addEventListener("click", () => {
    store.set("lang", current);
    bar.remove();
  });
  bar.append(p, go, no);
  document.body.append(bar);
})();
