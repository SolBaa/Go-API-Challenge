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
	 ('2021-12-07 08:56:36.1529-03','2021-12-07 08:56:36.1529-03',NULL,'11','Camila','Martinez','camila@gmail.com')
	 ('2021-12-09 07:31:28.773317-03','2021-12-09 07:31:28.773317-03',NULL,'180','Joaquin','Sargenti','joaquin@gmail.com'),
	 ('2021-12-09 07:32:08.322706-03','2021-12-09 07:32:08.322706-03',NULL,'25','Pablo','Lopez','pablo@gmail.com'),
	 ('2021-12-09 07:32:41.219826-03','2021-12-09 07:32:41.219826-03',NULL,'172','Lucas','Rosello','lucas@gmail.com'),
	 ('2021-12-09 07:33:30.125898-03','2021-12-09 07:33:30.125898-03',NULL,'73','Solana','Blois','solana@gmail.com');



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


public.companies foreign keys

ALTER TABLE public.companies ADD CONSTRAINT fk_users_company FOREIGN KEY (user_id) REFERENCES public.users(id);


INSERT INTO public.companies (created_at,updated_at,deleted_at,user_id,public_id,"name") VALUES
	 ('2021-12-09 07:27:04.066144-03','2021-12-09 07:27:04.066144-03',NULL,2,'','Marvik'),
	 ('2021-12-09 07:28:09.748504-03','2021-12-09 07:28:09.748504-03',NULL,3,'','Fake Company'),
	 ('2021-12-09 07:28:31.363687-03','2021-12-09 07:28:31.363687-03',NULL,4,'','Company 1'),
	 ('2021-12-09 07:29:26.044479-03','2021-12-09 07:29:26.044479-03',NULL,5,'','Company'),
	 ('2021-12-09 07:29:26.044479-03','2021-12-09 07:29:26.044479-03',NULL,5,'','Company 2'),
	 ('2021-12-09 07:31:28.773442-03','2021-12-09 07:31:28.773442-03',NULL,7,'','Company 3'),
	 ('2021-12-09 07:32:08.322865-03','2021-12-09 07:32:08.322865-03',NULL,8,'','Company 4'),
	 ('2021-12-09 07:32:41.219954-03','2021-12-09 07:32:41.219954-03',NULL,9,'','Company 5');