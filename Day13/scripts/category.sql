CREATE TABLE public.categories
(
    category_id smallint NOT NULL,
    category_name character varying(15) COLLATE pg_catalog."default" NOT NULL,
    description text COLLATE pg_catalog."default",
    picture bytea,
    CONSTRAINT pk_categories PRIMARY KEY (category_id)
)