-- TABLE: pets

CREATE TABLE public.pets (
  id uuid DEFAULT gen_random_uuid() NOT NULL,
  api_id bigint NOT NULL,

  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
  deleted_at TIMESTAMP WITH TIME ZONE NULL,

  name TEXT NOT NULL,
  tag TEXT,

  PRIMARY KEY (id)
);

SELECT manage_auto_updated_at('public.pets');
SELECT manage_auto_soft_delete('public.pets');

CREATE SEQUENCE public.pets_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.pets_id_seq OWNED BY public.pets.api_id;

ALTER TABLE ONLY public.pets ALTER COLUMN api_id SET DEFAULT nextval('public.pets_id_seq'::regclass);
