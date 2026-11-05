// Runs before the static files. It sends plain-HTTP requests to HTTPS (the zone's
// "Always Use HTTPS" setting would do the same; until it's on, this keeps http:// from
// serving a second copy of every page), and serves AI agents a page's Markdown copy
// (index.md, from scripts/markdown_pages.py) when they ask for Accept: text/markdown.
export default {
  async fetch(request, env) {
    const url = new URL(request.url);
    if (url.protocol === "http:") {
      url.protocol = "https:";
      return Response.redirect(url.toString(), 301);
    }

    const isPage = url.pathname.endsWith("/");
    if (!isPage) return env.ASSETS.fetch(request);

    const accept = request.headers.get("Accept") || "";
    if (/\btext\/markdown\b/i.test(accept)) {
      const md = await env.ASSETS.fetch(new URL(url.pathname + "index.md", url));
      if (md.ok) {
        const text = await md.text();
        return new Response(request.method === "HEAD" ? null : text, {
          headers: {
            "Content-Type": "text/markdown; charset=utf-8",
            "Vary": "Accept",
            // A rough count (about four characters a token), as Markdown for Agents sends.
            "x-markdown-tokens": String(Math.ceil(text.length / 4)),
            "Link": `<${url.pathname}>; rel="canonical"`,
          },
        });
      }
    }

    const res = await env.ASSETS.fetch(request);
    if (!res.ok || !(res.headers.get("Content-Type") || "").startsWith("text/html")) return res;
    const out = new Response(res.body, res);
    out.headers.append("Vary", "Accept");
    out.headers.append("Link", `<${url.pathname}index.md>; rel="alternate"; type="text/markdown"`);
    if (url.pathname === "/") {
      out.headers.append("Link", '</llms.txt>; rel="describedby"; type="text/plain"');
      out.headers.append("Link", '</sitemap.xml>; rel="sitemap"; type="application/xml"');
    }
    return out;
  },
};
