-- migrate:up
-- Caught by the form's hidden trap fields: kept, so nothing a person sent is lost, but not emailed.
alter table inquiries add column spam boolean not null default false;

-- migrate:down
alter table inquiries drop column spam;
