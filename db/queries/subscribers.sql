-- name: GetSubscriberByEmail :one
select * from subscribers where lower(email) = lower($1);

-- name: GetSubscriberByToken :one
select * from subscribers where token = $1;

-- name: CreateSubscriber :one
insert into subscribers (email, token, lang, source, ip, user_agent)
values ($1, $2, $3, $4, $5, $6)
returning *;

-- name: TouchSubscriber :exec
-- Another signup with an address we have: when it was, and from where.
update subscribers set requested_at = now(), ip = $2, user_agent = $3 where id = $1;

-- name: ResubscribeSubscriber :one
-- An address that unsubscribed signs up again: pending again, with a new token.
update subscribers
set token = $2, lang = $3, source = $4, ip = $5, user_agent = $6, requested_at = now(),
    confirmed_at = null, unsubscribed_at = null, confirmation_sent_at = null
where id = $1
returning *;

-- name: CountRecentSignupsFromIP :one
select count(*) from subscribers
where ip = $1 and requested_at > now() - make_interval(mins => sqlc.arg(minutes)::int);

-- name: MarkConfirmationSent :exec
update subscribers set confirmation_sent_at = now() where id = $1;

-- name: ConfirmSubscriber :one
-- Confirms a pending signup; a confirmed one stays as it was.
update subscribers set confirmed_at = coalesce(confirmed_at, now())
where token = $1 and unsubscribed_at is null
returning *;

-- name: UnsubscribeSubscriber :one
update subscribers set unsubscribed_at = coalesce(unsubscribed_at, now())
where token = $1
returning *;

-- name: ImportSubscriber :execrows
-- A confirmed subscriber brought over from ConvertKit. Addresses we already have are left alone.
insert into subscribers (email, token, source, confirmed_at, created_at, requested_at)
values ($1, $2, 'convertkit', $3, $3, $3)
on conflict (lower(email)) do nothing;

-- name: ListUnsentSubscribers :many
-- Confirmed subscribers who haven't had this issue yet.
select s.* from subscribers s
where s.confirmed_at is not null and s.unsubscribed_at is null
  and not exists (select 1 from newsletter_sends n where n.issue = $1 and n.subscriber_id = s.id)
order by s.id;

-- name: RecordNewsletterSend :exec
insert into newsletter_sends (issue, subscriber_id) values ($1, $2)
on conflict do nothing;

-- name: CountSubscribers :one
select
  count(*) filter (where confirmed_at is not null and unsubscribed_at is null) as confirmed,
  count(*) filter (where confirmed_at is null and unsubscribed_at is null) as pending,
  count(*) filter (where unsubscribed_at is not null) as unsubscribed
from subscribers;
