-- migrate:up
-- Inquiries sent with the contact form (POST /api/inquiries).
create table inquiries (
  id bigint generated always as identity primary key,
  name text not null,
  email text not null,
  company text not null default '',
  needs text[] not null default '{}', -- Application, Website, Design, Not sure yet
  message text not null,
  ip inet,
  user_agent text not null default '',
  created_at timestamptz not null default now()
);

create index inquiries_created_at on inquiries (created_at desc);

-- migrate:down
drop table inquiries;
