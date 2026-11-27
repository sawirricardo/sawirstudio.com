-- migrate:up
-- When the inquiry's email notification went out (null: not sent, or it failed).
alter table inquiries add column notified_at timestamptz;

-- migrate:down
alter table inquiries drop column notified_at;
