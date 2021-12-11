CREATE TABLE IF NOT EXISTS recipe (
	id serial NOT NULL,
	"name" varchar(255) NULL,
	description text NULL,
	author_id int NOT NULL
);

CREATE TABLE IF NOT EXISTS users (
	id serial NOT NULL,
	"name" varchar(45) NULL
);

CREATE TABLE IF NOT EXISTS recipe_category (
	id serial NOT NULL,
	"name" varchar(45) NOT NULL
);

CREATE TABLE IF NOT EXISTS recipe_category_recipe (
	recipe_category_id int NOT NULL,
	recipe_id int NOT NULL
);

CREATE TABLE IF NOT EXISTS step_ingredients (
	recipe_id int NOT NULL,
	ingredient_id int NOT NULL,
	step_id int NOT NULL,
	amount int NULL,
	unit varchar(25) NULL
);

CREATE TABLE IF NOT EXISTS ingredient (
	id serial NOT NULL,
	"name" varchar(45) NOT NULL,
	color int NOT NULL,
	img varchar(255) NULL
);

CREATE TABLE IF NOT EXISTS step (
	id serial NOT NULL,
	recipe_id int NOT NULL,
	step_number int NOT NULL,
	description text NOT NULL,
	timer int NULL,
	image varchar(100) NULL
);

CREATE TABLE IF NOT EXISTS ingredient_category (
	id serial NOT NULL,
	name varchar(255) NULL,
	description text NOT NULL
);

CREATE TABLE IF NOT EXISTS ingredient_category_ingredient (
	ingredient_category_id int NOT NULL,
	ingredient_id int NOT NULL
);

-- set primary key
ALTER TABLE recipe ADD CONSTRAINT recipe_pk PRIMARY KEY (id);

ALTER TABLE users ADD CONSTRAINT users_pk PRIMARY KEY (id);

ALTER TABLE recipe_category ADD CONSTRAINT recipe_category_pk PRIMARY KEY (id);

ALTER TABLE ingredient ADD CONSTRAINT ingredient_pk PRIMARY KEY (id);

ALTER TABLE step ADD CONSTRAINT step_pk PRIMARY KEY (id);

ALTER TABLE ingredient_category ADD CONSTRAINT ingredient_category_pk PRIMARY KEY (id);

-- set foreign key
ALTER TABLE recipe ADD CONSTRAINT recipe_fk FOREIGN KEY (author_id) REFERENCES users(id);

ALTER TABLE recipe_category_recipe ADD CONSTRAINT recipe_category_recipe_fk FOREIGN KEY (recipe_category_id) REFERENCES recipe_category(id);

ALTER TABLE recipe_category_recipe ADD CONSTRAINT recipe_category_recipe_fk1 FOREIGN KEY (recipe_id) REFERENCES recipe(id);

ALTER TABLE step_ingredients ADD CONSTRAINT step_ingredients_fk FOREIGN KEY (recipe_id) REFERENCES recipe(id);

ALTER TABLE step_ingredients ADD CONSTRAINT step_ingredients_fk1 FOREIGN KEY (ingredient_id) REFERENCES ingredient(id);

ALTER TABLE step_ingredients ADD CONSTRAINT step_ingredients_fk2 FOREIGN KEY (step_id) REFERENCES step(id);

ALTER TABLE ingredient_category_ingredient ADD CONSTRAINT ingredient_category_ingredient_fk FOREIGN KEY (ingredient_category_id) REFERENCES ingredient_category(id);

ALTER TABLE ingredient_category_ingredient ADD CONSTRAINT ingredient_category_ingredient_fk1 FOREIGN KEY (ingredient_id) REFERENCES ingredient(id);
