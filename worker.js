// Runs before the static files only to send plain-HTTP requests to HTTPS. The zone's
// "Always Use HTTPS" setting would do the same; until it's on, this keeps http:// from
// serving a second copy of every page.
export default {
  fetch(request, env) {
    const url = new URL(request.url);
    if (url.protocol === "http:") {
      url.protocol = "https:";
      return Response.redirect(url.toString(), 301);
    }
    return env.ASSETS.fetch(request);
  },
};
