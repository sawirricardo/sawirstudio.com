// Google Analytics
window.dataLayer = window.dataLayer || [];
function gtag() {
  window.dataLayer.push(arguments);
}
gtag("js", new Date());
gtag("config", "UA-135085180-1");

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
