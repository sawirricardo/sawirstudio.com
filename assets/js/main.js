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

  document.querySelector(".theme-toggle").addEventListener("click", () => {
    const theme = root.dataset.theme === "dark" ? "light" : "dark";
    root.dataset.theme = theme;
    try {
      localStorage.setItem("theme", theme);
    } catch (e) {}
  });

  // Follow system changes until the visitor picks a theme themselves.
  systemDark.addEventListener("change", (e) => {
    if (!savedTheme()) root.dataset.theme = e.matches ? "dark" : "light";
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
