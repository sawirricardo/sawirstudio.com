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
  const isList = (path) => path === "/" || path === "/projects/";
  const isProject = (path) => /^\/projects\/[^/]+\/$/.test(path);

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
    button.textContent = "Sending…";
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
      button.textContent = "Send inquiry";
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
  new MutationObserver(colors).observe(document.documentElement, { attributes: true, attributeFilter: ["data-theme"] });
  if ("IntersectionObserver" in window) {
    new IntersectionObserver(function (entries) { visible = entries[0].isIntersecting; kick(); }).observe(hero);
  }
  kick();
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
