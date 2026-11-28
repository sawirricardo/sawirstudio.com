-- migrate:up
-- The off-screen trap that set it gave way to Cloudflare Turnstile the same day.
alter table inquiries drop column spam;

-- migrate:down
alter table inquiries add column spam boolean not null default false;
