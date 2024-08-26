--
-- PostgreSQL database dump
--

-- Dumped from database version 15.4
-- Dumped by pg_dump version 15.4

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: uuid-ossp; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;


--
-- Name: EXTENSION "uuid-ossp"; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';


--
-- Name: update_products_updated_at_column(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.update_products_updated_at_column() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$;


ALTER FUNCTION public.update_products_updated_at_column() OWNER TO postgres;

--
-- Name: update_timestamp(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.update_timestamp() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$;


ALTER FUNCTION public.update_timestamp() OWNER TO postgres;

--
-- Name: update_updated_at_column(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.update_updated_at_column() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
   NEW.updated_at = now();
RETURN NEW;
END;
$$;


ALTER FUNCTION public.update_updated_at_column() OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: categories; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.categories (
    name character varying(255) NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    slug character varying(255),
    description character varying(255),
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    version integer DEFAULT 1
);


ALTER TABLE public.categories OWNER TO postgres;

--
-- Name: colors; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.colors (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    name character varying(255) NOT NULL,
    slug character varying(255) NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    css_variables character varying(255),
    version integer DEFAULT 1
);


ALTER TABLE public.colors OWNER TO postgres;

--
-- Name: goose_db_version; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.goose_db_version (
    id integer NOT NULL,
    version_id bigint NOT NULL,
    is_applied boolean NOT NULL,
    tstamp timestamp without time zone DEFAULT now()
);


ALTER TABLE public.goose_db_version OWNER TO postgres;

--
-- Name: goose_db_version_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.goose_db_version_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.goose_db_version_id_seq OWNER TO postgres;

--
-- Name: goose_db_version_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.goose_db_version_id_seq OWNED BY public.goose_db_version.id;


--
-- Name: product_colors; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.product_colors (
    product_id uuid NOT NULL,
    color_id uuid NOT NULL
);


ALTER TABLE public.product_colors OWNER TO postgres;

--
-- Name: product_sizes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.product_sizes (
    product_id uuid NOT NULL,
    size_id uuid NOT NULL
);


ALTER TABLE public.product_sizes OWNER TO postgres;

--
-- Name: products; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.products (
    name character varying(255) NOT NULL,
    price integer NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    version integer DEFAULT 1,
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    category_id uuid,
    description text,
    images text[] DEFAULT ARRAY[]::text[]
);


ALTER TABLE public.products OWNER TO postgres;

--
-- Name: size; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.size (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    name character varying(255) NOT NULL,
    slug character varying(255) NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    version integer DEFAULT 1
);


ALTER TABLE public.size OWNER TO postgres;

--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    name character varying(255) NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    phone character varying(15) NOT NULL,
    email text NOT NULL,
    avatar_path text DEFAULT 'https://www.google.com/url?sa=i&url=https%3A%2F%2Fru.dreamstime.com%2F%25D0%25B0%25D0%25BD%25D0%25BE%25D0%25BD%25D0%25B8%25D0%25BC%25D0%25BD%25D1%258B%25D0%25B9-%25D0%25B3%25D0%25B5%25D0%25BD%25D0%25B4%25D0%25B5%25D1%2580%25D0%25BD%25D0%25BE-%25D0%25BD%25D0%25B5%25D0%25B9%25D1%2582%25D1%2580%25D0%25B0%25D0%25BB%25D1%258C%25D0%25BD%25D1%258B%25D0%25B9-%25D0%25B0%25D0%25B2%25D0%25B0%25D1%2582%25D0%25B0%25D1%2580-%25D1%2581%25D0%25B8%25D0%25BB%25D1%2583%25D1%258D%25D1%2582-%25D0%25B3%25D0%25BE%25D0%25BB%25D0%25BE%25D0%25B2%25D1%258B-%25D0%25B8%25D0%25BD%25D0%25BA%25D0%25BE%25D0%25B3%25D0%25BD%25D0%25B8%25D1%2582%25D0%25BE-image227531366&psig=AOvVaw3ne_h-rK4XxvxDPC-lzVwc&ust=1717639583815000&source=images&cd=vfe&opi=89978449&ved=0CBIQjRxqFwoTCLi2oZyww4YDFQAAAAAdAAAAABAE'::text,
    role character varying(50) NOT NULL,
    password character varying(255) DEFAULT ''::character varying NOT NULL
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: goose_db_version id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.goose_db_version ALTER COLUMN id SET DEFAULT nextval('public.goose_db_version_id_seq'::regclass);


--
-- Data for Name: categories; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.categories (name, created_at, updated_at, slug, description, id, version) FROM stdin;
Худи	2024-05-20 14:47:13	2024-05-20 14:47:11	худи	худи на все цвета и размеры	3be89d88-e234-4c24-afb5-7015cedf6780	1
Футболки	2024-05-20 15:06:30	2024-05-20 15:06:31	футболки	футболки на любой вкус и цвет	3445eef1-4db1-4bcf-b8e5-c7a46d8b2443	1
Кофты	2024-05-20 16:15:37.993115	2024-05-20 16:15:37.993115	кофты	test	aa7f868b-e1d2-44ac-95dc-aa0f61adf58c	1
Штаны	2024-06-16 08:37:40.866485	2024-06-16 08:37:40.866485	штаны	штаны	07f5dcb2-e5db-44f4-8d48-62c433c46671	1
\.


--
-- Data for Name: colors; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.colors (id, name, slug, created_at, updated_at, css_variables, version) FROM stdin;
d77c9760-8a64-492e-99db-cece0bbb6faf	Белый	белый	2024-06-17 18:27:16.923144	2024-06-17 18:27:16.923144	#FFFAFA	1
7451ae9b-e1d2-4005-bf85-1fe0dbc620b4	Черный	черный	2024-06-17 18:27:48.083144	2024-06-17 18:27:48.083144	#000	1
ee405ed1-a834-44cd-8430-fc5d0570f99d	Красный	красный	2024-06-17 18:28:15.054577	2024-06-17 18:28:15.054577	#DC143C	1
dbdd9fd3-6e48-4596-8d30-4035aa8bfbaf	Зеленый	зеленый	2024-06-17 18:28:30.016756	2024-06-17 18:28:30.016756	#00FF00	1
\.


--
-- Data for Name: goose_db_version; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.goose_db_version (id, version_id, is_applied, tstamp) FROM stdin;
1	0	t	2024-05-20 10:47:52.540407
2	20240520074705	t	2024-05-20 10:51:13.493054
3	20240520080323	t	2024-05-20 11:09:01.69538
4	20240520081040	t	2024-05-20 11:17:35.597634
5	20240529112207	t	2024-05-29 14:41:44.905302
8	20240530174846	t	2024-05-30 22:11:53.948367
9	20240530191411	t	2024-05-30 22:14:26.419503
10	20240530191551	t	2024-05-30 22:16:42.541371
11	20240530200629	t	2024-05-30 23:10:11.427321
12	20240530201159	t	2024-05-30 23:14:56.821626
13	20240530201752	t	2024-05-30 23:24:42.482633
14	20240530202554	t	2024-05-30 23:26:25.851109
15	20240531175503	t	2024-05-31 20:58:41.287651
16	20240604191538	t	2024-06-05 05:07:33.163717
17	20240606223726	t	2024-06-07 01:40:03.251596
18	20240614030420	t	2024-06-14 06:08:35.006916
19	20240614031928	t	2024-06-14 06:21:12.667832
20	20240617063636	t	2024-06-17 10:00:18.933716
21	20240617070108	t	2024-06-17 10:14:43.765442
22	20240617142501	t	2024-06-17 17:26:26.307608
23	20240618053215	t	2024-06-18 08:35:31.497671
\.


--
-- Data for Name: product_colors; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.product_colors (product_id, color_id) FROM stdin;
\.


--
-- Data for Name: product_sizes; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.product_sizes (product_id, size_id) FROM stdin;
\.


--
-- Data for Name: products; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.products (name, price, created_at, updated_at, version, id, category_id, description, images) FROM stdin;
test	1000	2024-06-02 05:15:48.922685	2024-06-14 06:10:57.372118	1	6e874b18-dd9e-41a8-9a97-90b1188d62fd	aa7f868b-e1d2-44ac-95dc-aa0f61adf58c	test	{https://i.ibb.co/b2yb9Cx/photo-2023-09-29-04-16-21.jpg,https://i.ibb.co/gSLkXM6/photo-2023-09-29-04-16-22.jpg}
tshirts	2000	2024-05-20 10:56:46.25085	2024-06-14 06:13:41.334396	1	c664a1f5-11e0-4dce-b25e-a3e73a9d89de	aa7f868b-e1d2-44ac-95dc-aa0f61adf58c	sdqweasd 	{https://i.ibb.co/djfF1Y0/photo-2023-09-29-04-16-08-2.jpg,https://i.ibb.co/PNmxVnC/photo-2023-09-29-04-16-23.jpg}
Красная футболка	1500	2024-05-20 15:36:14.574797	2024-06-14 06:13:41.334396	1	861e8bce-a554-4986-aa5e-c48625f51b17	3be89d88-e234-4c24-afb5-7015cedf6780	asdawd	{https://i.ibb.co/88RwVZz/EL0A3923.jpg,https://i.ibb.co/syp6L34/EL0A3919.jpg}
Футболка	1500	2024-05-20 15:09:10.936012	2024-06-14 06:13:41.334396	1	92aa2dc6-87c9-4388-9cfa-b9be48141f6c	3445eef1-4db1-4bcf-b8e5-c7a46d8b2443	sdsadad	{https://i.ibb.co/r2vKWdL/EL0A3967.jpg,https://i.ibb.co/rQfBvS0/EL0A3969.jpg}
\.


--
-- Data for Name: size; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.size (id, name, slug, created_at, updated_at, version) FROM stdin;
3ebe052b-a022-422a-8508-6a1b765da193	XL	xl	2024-06-17 12:14:21.171819	2024-06-17 12:14:21.171819	1
4431b555-306b-4cf6-93ec-4df738ce17fd	L	l	2024-06-17 12:23:35.536249	2024-06-17 12:23:35.536249	1
c946b53b-b556-43ff-a810-b70447aa7c6e	S	s	2024-06-17 12:23:38.779019	2024-06-17 12:23:38.779019	1
71dea29a-fe6b-479f-9dfe-5261bbc3c804	M	m	2024-06-17 12:23:42.187672	2024-06-17 12:23:42.187672	1
f35edc7f-d63c-49a1-9a3d-bb6d94d10fbb	XS	xs	2024-06-17 12:23:47.210389	2024-06-17 12:23:47.210389	1
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, name, created_at, updated_at, phone, email, avatar_path, role, password) FROM stdin;
e565e9c3-5f89-493f-b8bc-f220dc064426	sanchir	2024-06-15 14:33:56.413918	2024-06-15 14:33:56.413918	98765	kaori@mail.ru	https://www.google.com/url?sa=i&url=https%3A%2F%2Fru.dreamstime.com%2F%25D0%25B0%25D0%25BD%25D0%25BE%25D0%25BD%25D0%25B8%25D0%25BC%25D0%25BD%25D1%258B%25D0%25B9-%25D0%25B3%25D0%25B5%25D0%25BD%25D0%25B4%25D0%25B5%25D1%2580%25D0%25BD%25D0%25BE-%25D0%25BD%25D0%25B5%25D0%25B9%25D1%2582%25D1%2580%25D0%25B0%25D0%25BB%25D1%258C%25D0%25BD%25D1%258B%25D0%25B9-%25D0%25B0%25D0%25B2%25D0%25B0%25D1%2582%25D0%25B0%25D1%2580-%25D1%2581%25D0%25B8%25D0%25BB%25D1%2583%25D1%258D%25D1%2582-%25D0%25B3%25D0%25BE%25D0%25BB%25D0%25BE%25D0%25B2%25D1%258B-%25D0%25B8%25D0%25BD%25D0%25BA%25D0%25BE%25D0%25B3%25D0%25BD%25D0%25B8%25D1%2582%25D0%25BE-image227531366&psig=AOvVaw3ne_h-rK4XxvxDPC-lzVwc&ust=1717639583815000&source=images&cd=vfe&opi=89978449&ved=0CBIQjRxqFwoTCLi2oZyww4YDFQAAAAAdAAAAABAE	ADMIN	vgO+A1f4abYdBY8ZHtSH/cY17W6hhgEDCBfv3dQlBZI
\.


--
-- Name: goose_db_version_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.goose_db_version_id_seq', 23, true);


--
-- Name: categories categories_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.categories
    ADD CONSTRAINT categories_pkey PRIMARY KEY (id);


--
-- Name: categories categories_slug_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.categories
    ADD CONSTRAINT categories_slug_key UNIQUE (slug);


--
-- Name: colors colors_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.colors
    ADD CONSTRAINT colors_pkey PRIMARY KEY (id);


--
-- Name: colors colors_slug_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.colors
    ADD CONSTRAINT colors_slug_key UNIQUE (slug);


--
-- Name: goose_db_version goose_db_version_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.goose_db_version
    ADD CONSTRAINT goose_db_version_pkey PRIMARY KEY (id);


--
-- Name: product_colors product_colors_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.product_colors
    ADD CONSTRAINT product_colors_pkey PRIMARY KEY (product_id, color_id);


--
-- Name: product_sizes product_sizes_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.product_sizes
    ADD CONSTRAINT product_sizes_pkey PRIMARY KEY (product_id, size_id);


--
-- Name: products products_new_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_new_pkey PRIMARY KEY (id);


--
-- Name: size size_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.size
    ADD CONSTRAINT size_pkey PRIMARY KEY (id);


--
-- Name: size size_slug_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.size
    ADD CONSTRAINT size_slug_key UNIQUE (slug);


--
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: users users_phone_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_phone_key UNIQUE (phone);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: colors set_timestamp; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER set_timestamp BEFORE UPDATE ON public.colors FOR EACH ROW EXECUTE FUNCTION public.update_timestamp();


--
-- Name: products update_products_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER update_products_updated_at BEFORE UPDATE ON public.products FOR EACH ROW EXECUTE FUNCTION public.update_products_updated_at_column();


--
-- Name: users update_user_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER update_user_updated_at BEFORE UPDATE ON public.users FOR EACH ROW EXECUTE FUNCTION public.update_updated_at_column();


--
-- Name: product_colors product_colors_color_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.product_colors
    ADD CONSTRAINT product_colors_color_id_fkey FOREIGN KEY (color_id) REFERENCES public.colors(id) ON DELETE CASCADE;


--
-- Name: product_colors product_colors_product_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.product_colors
    ADD CONSTRAINT product_colors_product_id_fkey FOREIGN KEY (product_id) REFERENCES public.products(id) ON DELETE CASCADE;


--
-- Name: product_sizes product_sizes_product_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.product_sizes
    ADD CONSTRAINT product_sizes_product_id_fkey FOREIGN KEY (product_id) REFERENCES public.products(id) ON DELETE CASCADE;


--
-- Name: product_sizes product_sizes_size_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.product_sizes
    ADD CONSTRAINT product_sizes_size_id_fkey FOREIGN KEY (size_id) REFERENCES public.size(id) ON DELETE CASCADE;


--
-- Name: products products_category_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_category_id_fkey FOREIGN KEY (category_id) REFERENCES public.categories(id);


--
-- PostgreSQL database dump complete
--

