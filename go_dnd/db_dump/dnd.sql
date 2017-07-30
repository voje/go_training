--
-- PostgreSQL database dump
--

-- Dumped from database version 9.5.7
-- Dumped by pg_dump version 9.5.7

SET statement_timeout = 0;
SET lock_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET search_path = public, pg_catalog;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: npc; Type: TABLE; Schema: public; Owner: kristjan
--

CREATE TABLE npc (
    id integer NOT NULL,
    name character varying(75) NOT NULL
);


ALTER TABLE npc OWNER TO kristjan;

--
-- Name: npc_id_seq; Type: SEQUENCE; Schema: public; Owner: kristjan
--

CREATE SEQUENCE npc_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE npc_id_seq OWNER TO kristjan;

--
-- Name: npc_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: kristjan
--

ALTER SEQUENCE npc_id_seq OWNED BY npc.id;


--
-- Name: race; Type: TABLE; Schema: public; Owner: kristjan
--

CREATE TABLE race (
    id integer NOT NULL,
    name character varying(75) NOT NULL
);


ALTER TABLE race OWNER TO kristjan;

--
-- Name: race_id_seq; Type: SEQUENCE; Schema: public; Owner: kristjan
--

CREATE SEQUENCE race_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE race_id_seq OWNER TO kristjan;

--
-- Name: race_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: kristjan
--

ALTER SEQUENCE race_id_seq OWNED BY race.id;


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: kristjan
--

ALTER TABLE ONLY npc ALTER COLUMN id SET DEFAULT nextval('npc_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: kristjan
--

ALTER TABLE ONLY race ALTER COLUMN id SET DEFAULT nextval('race_id_seq'::regclass);


--
-- Data for Name: npc; Type: TABLE DATA; Schema: public; Owner: kristjan
--

COPY npc (id, name) FROM stdin;
1	Terminator
2	Django
\.


--
-- Name: npc_id_seq; Type: SEQUENCE SET; Schema: public; Owner: kristjan
--

SELECT pg_catalog.setval('npc_id_seq', 2, true);


--
-- Data for Name: race; Type: TABLE DATA; Schema: public; Owner: kristjan
--

COPY race (id, name) FROM stdin;
1	Humes
\.


--
-- Name: race_id_seq; Type: SEQUENCE SET; Schema: public; Owner: kristjan
--

SELECT pg_catalog.setval('race_id_seq', 1, true);


--
-- Name: npc_pkey; Type: CONSTRAINT; Schema: public; Owner: kristjan
--

ALTER TABLE ONLY npc
    ADD CONSTRAINT npc_pkey PRIMARY KEY (id);


--
-- Name: race_pkey; Type: CONSTRAINT; Schema: public; Owner: kristjan
--

ALTER TABLE ONLY race
    ADD CONSTRAINT race_pkey PRIMARY KEY (id);


--
-- Name: public; Type: ACL; Schema: -; Owner: postgres
--

REVOKE ALL ON SCHEMA public FROM PUBLIC;
REVOKE ALL ON SCHEMA public FROM postgres;
GRANT ALL ON SCHEMA public TO postgres;
GRANT ALL ON SCHEMA public TO PUBLIC;


--
-- Name: npc; Type: ACL; Schema: public; Owner: kristjan
--

REVOKE ALL ON TABLE npc FROM PUBLIC;
REVOKE ALL ON TABLE npc FROM kristjan;
GRANT ALL ON TABLE npc TO kristjan;
GRANT ALL ON TABLE npc TO gopher;


--
-- Name: npc_id_seq; Type: ACL; Schema: public; Owner: kristjan
--

REVOKE ALL ON SEQUENCE npc_id_seq FROM PUBLIC;
REVOKE ALL ON SEQUENCE npc_id_seq FROM kristjan;
GRANT ALL ON SEQUENCE npc_id_seq TO kristjan;
GRANT ALL ON SEQUENCE npc_id_seq TO gopher;


--
-- Name: race; Type: ACL; Schema: public; Owner: kristjan
--

REVOKE ALL ON TABLE race FROM PUBLIC;
REVOKE ALL ON TABLE race FROM kristjan;
GRANT ALL ON TABLE race TO kristjan;
GRANT ALL ON TABLE race TO gopher;


--
-- Name: race_id_seq; Type: ACL; Schema: public; Owner: kristjan
--

REVOKE ALL ON SEQUENCE race_id_seq FROM PUBLIC;
REVOKE ALL ON SEQUENCE race_id_seq FROM kristjan;
GRANT ALL ON SEQUENCE race_id_seq TO kristjan;
GRANT ALL ON SEQUENCE race_id_seq TO gopher;


--
-- PostgreSQL database dump complete
--

