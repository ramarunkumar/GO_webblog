-- Table: public.blog_user

-- DROP TABLE IF EXISTS public.blog_user;

CREATE TABLE IF NOT EXISTS public.blog_user
(
    id integer NOT NULL DEFAULT nextval('blog_user_id_seq'::regclass),
    name text COLLATE pg_catalog."default",
    username text COLLATE pg_catalog."default",
    password character varying(255) COLLATE pg_catalog."default",
    CONSTRAINT blog_user_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.blog_user
    OWNER to postgres;

SELECT content,count(content) from blogs group by content;
--Restart data---
TRUNCATE <table name> RESTART IDENTITY;



    -- Table: public.blogs

-- DROP TABLE IF EXISTS public.blogs;

CREATE TABLE IF NOT EXISTS public.blogs
(
    article_id integer NOT NULL DEFAULT nextval('blogs_article_id_seq'::regclass),
    blog_id integer NOT NULL DEFAULT nextval('blogs_blog_id_seq'::regclass),
    title text COLLATE pg_catalog."default",
    content text COLLATE pg_catalog."default",
    CONSTRAINT blogs_pkey PRIMARY KEY (article_id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.blogs
    OWNER to postgres;