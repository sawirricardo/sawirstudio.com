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
