CREATE TABLE public.item_category (
	item_category_id int4 NOT NULL DEFAULT nextval('item_category_item_category_id_seq'::regclass),
	item_category_name varchar(100) NOT NULL,
	created_by varchar(50) NULL,
	created_time timestamp NULL,
	modified_by varchar(50) NULL,
	modified_time timestamp NULL,
	CONSTRAINT item_category_pkey PRIMARY KEY (item_category_id)
)
WITH (
	OIDS=FALSE
);

CREATE TABLE public.user_access (
	user_id int4 NOT NULL DEFAULT nextval('user_access_user_id_seq'::regclass),
	username varchar(50) NOT NULL,
	password varchar(255) NOT NULL,
	first_name varchar(255) NULL,
	last_name varchar(255) NULL,
	birthday date NULL,
	gender varchar(1) NULL,
	mobile_phone varchar(15) NULL,
	email_address varchar(255) NULL,
	ip_address varchar(50) NULL,
	country_id int4 NULL,
	state_id int4 NULL,
	city_id int4 NULL,
	CONSTRAINT user_access_pkey PRIMARY KEY (user_id)
)
WITH (
	OIDS=FALSE
);

INSERT INTO public.item_category
(item_category_id, item_category_name, created_by, created_time, modified_by, modified_time)
VALUES(1, 'Women''s Clothing', 'bagongkia', '2017-06-11 14:36:23.002', 'bagongkia', '2017-06-11 14:59:57.841');