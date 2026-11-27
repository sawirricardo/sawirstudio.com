-- migrate:up
-- Newsletter subscribers, from the form at the end of each blog post (POST /api/subscribe).
-- A signup is pending until the address confirms it from the email we send (double opt-in).
create table subscribers (
  id bigint generated always as identity primary key,
  email text not null,
  -- In the confirmation and unsubscribe links. Random; a new one when an unsubscribed
  -- address signs up again.
  token text not null unique,
  lang text not null default 'en',  -- the page's language when they signed up
  source text not null default '',  -- the page they signed up on, or 'convertkit' for imports
  ip inet,
  user_agent text not null default '',
  created_at timestamptz not null default now(),
  requested_at timestamptz not null default now(),  -- the latest signup with this address
  confirmation_sent_at timestamptz,
  confirmed_at timestamptz,
  unsubscribed_at timestamptz
);

create unique index subscribers_email on subscribers (lower(email));
create index subscribers_ip_requested on subscribers (ip, requested_at);

-- One row per newsletter email sent to a subscriber, so a send can stop and pick up again
-- without emailing anyone twice.
create table newsletter_sends (
  issue text not null,  -- the post's path, like /posts/services-is-next/
  subscriber_id bigint not null references subscribers (id) on delete cascade,
  sent_at timestamptz not null default now(),
  primary key (issue, subscriber_id)
);

-- migrate:down
drop table newsletter_sends;
drop table subscribers;
