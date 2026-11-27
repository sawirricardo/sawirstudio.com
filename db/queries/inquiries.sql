-- name: CreateInquiry :one
insert into inquiries (name, email, company, needs, message, ip, user_agent)
values ($1, $2, $3, $4, $5, $6, $7)
returning id, created_at;

-- name: ListInquiries :many
select * from inquiries
order by created_at desc
limit $1;

-- name: CountRecentInquiriesFromIP :one
select count(*) from inquiries
where ip = $1 and created_at > now() - make_interval(mins => sqlc.arg(minutes)::int);

-- name: MarkInquiryNotified :exec
update inquiries set notified_at = now() where id = $1;

-- name: ListUnnotifiedInquiries :many
select * from inquiries
where notified_at is null and created_at > now() - interval '7 days'
order by created_at;
