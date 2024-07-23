CREATE TABLE categories
(
    category_id smallint NOT NULL,
    category_name character varying(15) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT pk_categories PRIMARY KEY (category_id)
)