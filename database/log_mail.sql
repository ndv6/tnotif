DROP SEQUENCE IF EXISTS log_mail_id_seq;
CREATE SEQUENCE log_mail_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 START 1 CACHE 1;

CREATE TABLE IF NOT EXISTS "public"."log_mail" (
    id integer DEFAULT nextval('log_mail_id_seq') NOT NULL,
    email character varying(100),
    sent_at timestamp
)   WITH (oids = false);