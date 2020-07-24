CREATE TABLE IF NOT EXISTS "public"."log_mail" (
    id SERIAL NOT NULL,
    email character varying(100),
    sent_at timestamp
)   WITH (oids = false);