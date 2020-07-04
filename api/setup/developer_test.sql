-- SEQUENCE: public.customer_companies_company_id_seq

-- DROP SEQUENCE public.customer_companies_company_id_seq;

CREATE SEQUENCE public.customer_companies_company_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 2147483647
    CACHE 1;

ALTER SEQUENCE public.customer_companies_company_id_seq
    OWNER TO postgres;

GRANT ALL ON SEQUENCE public.customer_companies_company_id_seq TO orderviewer WITH GRANT OPTION;

GRANT ALL ON SEQUENCE public.customer_companies_company_id_seq TO postgres WITH GRANT OPTION;

GRANT ALL ON SEQUENCE public.customer_companies_company_id_seq TO vanjoe WITH GRANT OPTION;

GRANT ALL ON SEQUENCE public.customer_companies_company_id_seq TO vanjoechua WITH GRANT OPTION;

-- SEQUENCE: public.customers_user_id_seq

-- DROP SEQUENCE public.customers_user_id_seq;

CREATE SEQUENCE public.customers_user_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 2147483647
    CACHE 1;

ALTER SEQUENCE public.customers_user_id_seq
    OWNER TO postgres;

GRANT ALL ON SEQUENCE public.customers_user_id_seq TO orderviewer WITH GRANT OPTION;

GRANT ALL ON SEQUENCE public.customers_user_id_seq TO postgres;

GRANT ALL ON SEQUENCE public.customers_user_id_seq TO vanjoe WITH GRANT OPTION;

GRANT ALL ON SEQUENCE public.customers_user_id_seq TO vanjoechua WITH GRANT OPTION;

-- SEQUENCE: public.deliveries_id_seq

-- DROP SEQUENCE public.deliveries_id_seq;

CREATE SEQUENCE public.deliveries_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 2147483647
    CACHE 1;

ALTER SEQUENCE public.deliveries_id_seq
    OWNER TO postgres;

GRANT ALL ON SEQUENCE public.deliveries_id_seq TO orderviewer WITH GRANT OPTION;

GRANT ALL ON SEQUENCE public.deliveries_id_seq TO postgres;

GRANT ALL ON SEQUENCE public.deliveries_id_seq TO vanjoe WITH GRANT OPTION;

GRANT ALL ON SEQUENCE public.deliveries_id_seq TO vanjoechua WITH GRANT OPTION;

-- SEQUENCE: public.order_items_id_seq

-- DROP SEQUENCE public.order_items_id_seq;

CREATE SEQUENCE public.order_items_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 2147483647
    CACHE 1;

ALTER SEQUENCE public.order_items_id_seq
    OWNER TO postgres;

GRANT ALL ON SEQUENCE public.order_items_id_seq TO orderviewer WITH GRANT OPTION;

GRANT ALL ON SEQUENCE public.order_items_id_seq TO postgres;

GRANT ALL ON SEQUENCE public.order_items_id_seq TO vanjoe WITH GRANT OPTION;

GRANT ALL ON SEQUENCE public.order_items_id_seq TO vanjoechua WITH GRANT OPTION;

-- SEQUENCE: public.orders_id_seq

-- DROP SEQUENCE public.orders_id_seq;

CREATE SEQUENCE public.orders_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 2147483647
    CACHE 1;

ALTER SEQUENCE public.orders_id_seq
    OWNER TO postgres;

GRANT ALL ON SEQUENCE public.orders_id_seq TO orderviewer WITH GRANT OPTION;

GRANT ALL ON SEQUENCE public.orders_id_seq TO postgres;

GRANT ALL ON SEQUENCE public.orders_id_seq TO vanjoe WITH GRANT OPTION;

GRANT ALL ON SEQUENCE public.orders_id_seq TO vanjoechua WITH GRANT OPTION;

-- Table: public.customer_companies

-- DROP TABLE public.customer_companies;

CREATE TABLE public.customer_companies
(
    company_id integer NOT NULL DEFAULT nextval('customer_companies_company_id_seq'::regclass),
    company_name character varying(100) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT customer_companies_pkey PRIMARY KEY (company_id)
)

TABLESPACE pg_default;

ALTER TABLE public.customer_companies
    OWNER to postgres;

GRANT ALL ON TABLE public.customer_companies TO orderviewer WITH GRANT OPTION;

GRANT ALL ON TABLE public.customer_companies TO postgres WITH GRANT OPTION;

GRANT ALL ON TABLE public.customer_companies TO vanjoe WITH GRANT OPTION;

GRANT ALL ON TABLE public.customer_companies TO vanjoechua WITH GRANT OPTION;

-- Table: public.customers

-- DROP TABLE public.customers;

CREATE TABLE public.customers
(
    user_id character varying COLLATE pg_catalog."default" NOT NULL,
    login character varying COLLATE pg_catalog."default" NOT NULL,
    password character varying COLLATE pg_catalog."default" NOT NULL,
    name character varying COLLATE pg_catalog."default" NOT NULL,
    company_id integer NOT NULL,
    credit_cards character varying COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT customers_pkey PRIMARY KEY (user_id)
)

TABLESPACE pg_default;

ALTER TABLE public.customers
    OWNER to postgres;

GRANT ALL ON TABLE public.customers TO orderviewer WITH GRANT OPTION;

GRANT ALL ON TABLE public.customers TO postgres WITH GRANT OPTION;

GRANT ALL ON TABLE public.customers TO vanjoe WITH GRANT OPTION;

GRANT ALL ON TABLE public.customers TO vanjoechua WITH GRANT OPTION;

-- Table: public.deliveries

-- DROP TABLE public.deliveries;

CREATE TABLE public.deliveries
(
    id integer NOT NULL DEFAULT nextval('deliveries_id_seq'::regclass),
    order_item_id integer,
    delivered_qty integer,
    CONSTRAINT deliveries_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE public.deliveries
    OWNER to postgres;

GRANT ALL ON TABLE public.deliveries TO orderviewer WITH GRANT OPTION;

GRANT ALL ON TABLE public.deliveries TO postgres WITH GRANT OPTION;

GRANT ALL ON TABLE public.deliveries TO vanjoe WITH GRANT OPTION;

GRANT ALL ON TABLE public.deliveries TO vanjoechua WITH GRANT OPTION;

-- Table: public.order_items

-- DROP TABLE public.order_items;

CREATE TABLE public.order_items
(
    id integer NOT NULL DEFAULT nextval('order_items_id_seq'::regclass),
    order_id integer NOT NULL,
    price_per_unit numeric(15,6) DEFAULT NULL::numeric,
    quantity integer NOT NULL,
    product character varying COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT order_items_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE public.order_items
    OWNER to postgres;

GRANT ALL ON TABLE public.order_items TO orderviewer WITH GRANT OPTION;

GRANT ALL ON TABLE public.order_items TO postgres WITH GRANT OPTION;

GRANT ALL ON TABLE public.order_items TO vanjoe WITH GRANT OPTION;

GRANT ALL ON TABLE public.order_items TO vanjoechua WITH GRANT OPTION;

-- Table: public.orders

-- DROP TABLE public.orders;

CREATE TABLE public.orders
(
    id integer NOT NULL DEFAULT nextval('orders_id_seq'::regclass),
    created_at timestamp with time zone NOT NULL,
    order_name character varying COLLATE pg_catalog."default" NOT NULL,
    customer_id character varying COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT orders_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE public.orders
    OWNER to postgres;

GRANT ALL ON TABLE public.orders TO orderviewer WITH GRANT OPTION;

GRANT ALL ON TABLE public.orders TO postgres WITH GRANT OPTION;

GRANT ALL ON TABLE public.orders TO vanjoe WITH GRANT OPTION;

GRANT ALL ON TABLE public.orders TO vanjoechua WITH GRANT OPTION;