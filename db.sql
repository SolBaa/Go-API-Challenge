-- public.users definition

-- Drop table

-- DROP TABLE public.users;

CREATE TABLE public.users (
	id bigserial NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	public_id text NULL,
	"name" text NULL,
	last_name text NULL,
	email text NULL,
	CONSTRAINT users_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_users_deleted_at ON public.users USING btree (deleted_at);

INSERT INTO public.users (created_at,updated_at,deleted_at,public_id,"name",last_name,email) VALUES
	 ('2021-12-07 08:00:52.039226-03','2021-12-07 08:00:52.039226-03',NULL,'105','Sol','Battaglia','sol@gmial.com'),
	 ('2021-12-07 08:03:07.112641-03','2021-12-07 08:03:07.112641-03',NULL,'196','Pepe','Lopez','pepe@gmail.com'),
	 ('2021-12-07 08:11:46.643136-03','2021-12-07 08:11:46.643136-03',NULL,'177','Lucas','Perez','lucas@gmail.com'),
	 ('2021-12-07 08:26:04.338211-03','2021-12-07 08:26:04.338211-03',NULL,'22','Matias','Rodriguez','mat@gmail.com'),
	 ('2021-12-07 08:27:19.041635-03','2021-12-07 08:27:19.041635-03',NULL,'104','Juan','Alonso','Juan@gmail.com'),
	 ('2021-12-07 08:56:36.1529-03','2021-12-07 08:56:36.1529-03',NULL,'11','Camila','Martinez','camila@gmail.com');



-- public.companies definition

-- Drop table

-- DROP TABLE public.companies;

CREATE TABLE public.companies (
	id bigserial NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	user_id int8 NULL,
	public_id text NULL,
	"name" text NULL,
	CONSTRAINT companies_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_companies_deleted_at ON public.companies USING btree (deleted_at);


-- public.companies foreign keys

ALTER TABLE public.companies ADD CONSTRAINT fk_users_company FOREIGN KEY (user_id) REFERENCES public.users(id);


INSERT INTO public.companies (created_at,updated_at,deleted_at,user_id,public_id,"name") VALUES
	 ('2021-12-08 08:29:38.905971-03','2021-12-08 08:29:38.905971-03',NULL,NULL,'152','Marvik'),
	 ('2021-12-08 08:31:33.14375-03','2021-12-08 08:31:33.14375-03',NULL,NULL,'67','Fake Company'),
	 ('2021-12-08 08:31:55.657208-03','2021-12-08 08:31:55.657208-03',NULL,NULL,'85','Company 1'),
	 ('2021-12-08 08:32:09.593674-03','2021-12-08 08:32:09.593674-03',NULL,NULL,'26','Unexistent Company');