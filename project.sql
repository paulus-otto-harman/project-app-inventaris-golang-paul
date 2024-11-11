--
-- PostgreSQL database dump
--

-- Dumped from database version 16.4 (Homebrew)
-- Dumped by pg_dump version 16.4 (Homebrew)

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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: categories; Type: TABLE; Schema: public; Owner: paul
--

CREATE TABLE public.categories (
    id smallint NOT NULL,
    name character varying(20) NOT NULL,
    description character varying(80) NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.categories OWNER TO paul;

--
-- Name: categories_id_seq; Type: SEQUENCE; Schema: public; Owner: paul
--

CREATE SEQUENCE public.categories_id_seq
    AS smallint
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.categories_id_seq OWNER TO paul;

--
-- Name: categories_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: paul
--

ALTER SEQUENCE public.categories_id_seq OWNED BY public.categories.id;


--
-- Name: items; Type: TABLE; Schema: public; Owner: paul
--

CREATE TABLE public.items (
    id smallint NOT NULL,
    name character varying(25) NOT NULL,
    category_id smallint NOT NULL,
    photo_url character varying(255),
    price money NOT NULL,
    purchase_date date NOT NULL,
    depreciation_rate numeric(3,0) DEFAULT 10 NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.items OWNER TO paul;

--
-- Name: items_id_seq; Type: SEQUENCE; Schema: public; Owner: paul
--

CREATE SEQUENCE public.items_id_seq
    AS smallint
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.items_id_seq OWNER TO paul;

--
-- Name: items_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: paul
--

ALTER SEQUENCE public.items_id_seq OWNED BY public.items.id;


--
-- Name: categories id; Type: DEFAULT; Schema: public; Owner: paul
--

ALTER TABLE ONLY public.categories ALTER COLUMN id SET DEFAULT nextval('public.categories_id_seq'::regclass);


--
-- Name: items id; Type: DEFAULT; Schema: public; Owner: paul
--

ALTER TABLE ONLY public.items ALTER COLUMN id SET DEFAULT nextval('public.items_id_seq'::regclass);


--
-- Data for Name: categories; Type: TABLE DATA; Schema: public; Owner: paul
--

COPY public.categories (id, name, description, created_at, updated_at, deleted_at) FROM stdin;
1	Peralatan Kantor	Kategori untuk peralatan kantor	2024-11-09 17:01:03.493129+07	\N	\N
3	Peralatan Elektronis	Peralatan elektronik yang digunakan di kantor	2024-11-09 17:15:24.087485+07	\N	\N
\.


--
-- Data for Name: items; Type: TABLE DATA; Schema: public; Owner: paul
--

COPY public.items (id, name, category_id, photo_url, price, purchase_date, depreciation_rate, created_at, updated_at, deleted_at) FROM stdin;
2	Item 002	3	/images/two.jpg	$150,000.00	2024-06-12	10	2024-06-12 14:23:18+07	\N	\N
3	Item 003	1	/images/three.jpg	$220,000.00	2024-04-10	10	2024-04-10 08:37:12+07	\N	\N
4	Item 004	3	/images/four.jpg	$185,000.00	2024-05-22	10	2024-05-22 11:40:30+07	\N	\N
5	Item 005	1	/images/five.jpg	$290,000.00	2024-07-15	10	2024-07-15 13:09:22+07	\N	\N
6	Item 006	3	/images/six.jpg	$130,000.00	2024-08-05	10	2024-08-05 09:15:12+07	\N	\N
7	Item 007	1	/images/seven.jpg	$170,000.00	2024-09-25	10	2024-09-25 17:23:33+07	\N	\N
8	Item 008	3	/images/eight.jpg	$250,000.00	2024-07-30	10	2024-07-30 15:42:45+07	\N	\N
9	Item 009	1	/images/nine.jpg	$275,000.00	2024-06-28	10	2024-06-28 14:58:27+07	\N	\N
10	Item 010	3	/images/ten.jpg	$120,000.00	2024-10-04	10	2024-10-04 12:11:11+07	\N	\N
11	Item 011	1	/images/eleven.jpg	$240,000.00	2024-09-19	10	2024-09-19 09:40:20+07	\N	\N
12	Item 012	3	/images/twelve.jpg	$155,000.00	2024-07-07	10	2024-07-07 10:25:30+07	\N	\N
13	Item 013	1	/images/thirteen.jpg	$210,000.00	2024-05-15	10	2024-05-15 18:32:14+07	\N	\N
14	Item 014	3	/images/fourteen.jpg	$285,000.00	2024-06-20	10	2024-06-20 14:16:19+07	\N	\N
15	Item 015	1	/images/fifteen.jpg	$300,000.00	2024-08-18	10	2024-08-18 11:27:59+07	\N	\N
16	Item 016	3	/images/sixteen.jpg	$100,000.00	2024-07-09	10	2024-07-09 13:50:48+07	\N	\N
17	Item 017	1	/images/seventeen.jpg	$180,000.00	2024-09-28	10	2024-09-28 16:05:25+07	\N	\N
18	Item 018	3	/images/eighteen.jpg	$200,000.00	2024-05-13	10	2024-05-13 17:45:31+07	\N	\N
19	Item 019	1	/images/nineteen.jpg	$250,000.00	2024-06-08	10	2024-06-08 08:33:50+07	\N	\N
20	Item 020	3	/images/twenty.jpg	$190,000.00	2024-10-06	10	2024-10-06 09:15:09+07	\N	\N
21	Item 021	1	/images/twentyone.jpg	$210,000.00	2024-05-27	10	2024-05-27 10:23:18+07	\N	\N
22	Item 022	3	/images/twentytwo.jpg	$280,000.00	2024-08-23	10	2024-08-23 07:44:29+07	\N	\N
23	Item 023	1	/images/twentythree.jpg	$145,000.00	2024-09-01	10	2024-09-01 09:38:13+07	\N	\N
24	Item 024	3	/images/twentyfour.jpg	$270,000.00	2024-07-25	10	2024-07-25 13:10:35+07	\N	\N
25	Item 025	1	/images/twentyfive.jpg	$220,000.00	2024-06-02	10	2024-06-02 14:55:21+07	\N	\N
26	Item 026	3	/images/twentysix.jpg	$170,000.00	2024-09-12	10	2024-09-12 08:22:41+07	\N	\N
27	Item 027	1	/images/twentyseven.jpg	$155,000.00	2024-08-08	10	2024-08-08 17:30:22+07	\N	\N
28	Item 028	3	/images/twentyeight.jpg	$130,000.00	2024-10-15	10	2024-10-15 11:58:34+07	\N	\N
29	Item 029	1	/images/twentynine.jpg	$295,000.00	2024-04-29	10	2024-04-29 09:14:29+07	\N	\N
30	Item 030	3	/images/thirty.jpg	$160,000.00	2024-05-19	10	2024-05-19 12:05:10+07	\N	\N
31	Item 031	1	/images/thirtyone.jpg	$145,000.00	2024-06-30	10	2024-06-30 11:25:45+07	\N	\N
32	Item 032	3	/images/thirtytwo.jpg	$170,000.00	2024-09-18	10	2024-09-18 16:41:22+07	\N	\N
33	Item 033	1	/images/thirtythree.jpg	$200,000.00	2024-08-27	10	2024-08-27 15:09:55+07	\N	\N
34	Item 034	3	/images/thirtyfour.jpg	$180,000.00	2024-07-14	10	2024-07-14 18:11:44+07	\N	\N
35	Item 035	1	/images/thirtyfive.jpg	$210,000.00	2024-10-01	10	2024-10-01 10:03:12+07	\N	\N
36	Item 036	3	/images/thirtysix.jpg	$140,000.00	2024-05-25	10	2024-05-25 08:28:35+07	\N	\N
37	Item 037	1	/images/thirtyseven.jpg	$190,000.00	2024-06-16	10	2024-06-16 09:47:52+07	\N	\N
38	Item 038	3	/images/thirtyeight.jpg	$275,000.00	2024-08-01	10	2024-08-01 14:59:13+07	\N	\N
39	Item 039	1	/images/thirtynine.jpg	$250,000.00	2024-10-10	10	2024-10-10 11:04:40+07	\N	\N
40	Item 040	3	/images/forty.jpg	$175,000.00	2024-07-19	10	2024-07-19 13:11:59+07	\N	\N
1	Item 001	1	/images/one.jpg	$120,000.00	2024-05-08	10	2024-05-08 10:15:45+07	\N	\N
\.


--
-- Name: categories_id_seq; Type: SEQUENCE SET; Schema: public; Owner: paul
--

SELECT pg_catalog.setval('public.categories_id_seq', 3, true);


--
-- Name: items_id_seq; Type: SEQUENCE SET; Schema: public; Owner: paul
--

SELECT pg_catalog.setval('public.items_id_seq', 40, true);


--
-- Name: categories categories_name_key; Type: CONSTRAINT; Schema: public; Owner: paul
--

ALTER TABLE ONLY public.categories
    ADD CONSTRAINT categories_name_key UNIQUE (name);


--
-- Name: categories categories_pkey; Type: CONSTRAINT; Schema: public; Owner: paul
--

ALTER TABLE ONLY public.categories
    ADD CONSTRAINT categories_pkey PRIMARY KEY (id);


--
-- Name: items items_pkey; Type: CONSTRAINT; Schema: public; Owner: paul
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT items_pkey PRIMARY KEY (id, created_at);


--
-- Name: items items_category_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: paul
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT items_category_id_fkey FOREIGN KEY (category_id) REFERENCES public.categories(id) NOT VALID;


--
-- PostgreSQL database dump complete
--

