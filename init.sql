CREATE TABLE IF NOT EXISTS users (
  id serial NOT NULL,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp NOT NULL DEFAULT now(),
  username varchar(255) NOT NULL,
  password varchar(255) NOT NULL,
  name varchar(255) NOT NULL,
  email varchar(255) NOT NULL,
  image varchar(255),
  PRIMARY KEY (id),
  UNIQUE (username),
  UNIQUE (email)
);


CREATE TABLE IF NOT EXISTS seller (
  id serial NOT NULL,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp NOT NULL DEFAULT now(),  
  username varchar(255) NOT NULL,
  password varchar(255) NOT NULL,
  name varchar(255) NOT NULL,
  email varchar(255) NOT NULL,
  image varchar(255) NOT NULL,
  PRIMARY KEY (id),
  UNIQUE (username),
  UNIQUE (email)
);

CREATE TABLE IF NOT EXISTS item (
  id serial NOT NULL,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp NOT NULL DEFAULT now(),
  name varchar(255) NOT NULL,
  description varchar(255) NOT NULL,
  is_auction boolean NOT NULL,
  is_active boolean NOT NULL,
  start_auct timestamp NOT NULL,
  end_auct timestamp NOT NULL,
  start_price bigint NOT NULL,
  seller_id int NOT NULL,
  PRIMARY KEY (id),
  CONSTRAINT fk_item_seller FOREIGN KEY(seller_id) REFERENCES seller(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS item_image (
  id serial NOT NULL,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp NOT NULL DEFAULT now(),
  image varchar(255) NOT NULL,
  item_id int NOT NULL,
  PRIMARY KEY (id),
  CONSTRAINT fk_item_image_item FOREIGN KEY(item_id) REFERENCES item(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS item_bid (
  id serial NOT NULL,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp NOT NULL DEFAULT now(),
  bid_price bigint NOT NULL,
  item_id int NOT NULL,
  user_id int NOT NULL,
  PRIMARY KEY (id),
  CONSTRAINT fk_item_bid_item FOREIGN KEY(item_id) REFERENCES item(id) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT fk_item_bid_user FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE
);

-- Path: init.sql

INSERT INTO users (username, password, name, email, image) VALUES 
('akimpi', '12345', 'Abdul Hakim', 'abdul.hakim@gmail.com', 'https://ucarecdn.com/0109f00c-d7f6-4a9f-8e3f-943605298025/-/preview/500x500/-/quality/smart/-/format/auto/');

INSERT INTO seller (username, password, name, email, image) VALUES 
('barang_antiq', '12345', 'Barang Antik', 'barang.antik@gmail.com', 'https://ucarecdn.com/c11d3018-8f6e-4d7b-96aa-389253e90c17/-/preview/500x500/-/quality/smart/-/format/auto/');

INSERT INTO item (name, description, is_auction, is_active, start_auct, end_auct, start_price, seller_id) VALUES 
('Jam Tangan', 'Jam tangan antik', true, true, '2023-12-07 22:00:00', '2023-12-14 22:00:00', 1000000, 1);

INSERT INTO item_image (image, item_id) VALUES 
('https://ucarecdn.com/86dc6ac8-c39f-4218-83e8-fce0cdf809af/jamtangan.jpeg', 1),
('https://ucarecdn.com/afe47101-9ac5-44f4-8888-74a6ea289c9c/jamtangan2.jpg', 1);

INSERT INTO item_bid (bid_price, item_id, user_id) VALUES 
(1000000, 1, 1),
(2000000, 1, 1);

INSERT INTO item (name, description, is_auction, is_active, start_auct, end_auct, start_price, seller_id) VALUES 
('Lukisan Raden Saleh', 'Lukisan raden saleh yang pernah dicuri', true, true, '2023-12-07 22:31:00', '2023-12-14 22:31:00', 3000000, 1),
('Sofa Kayu', 'Sofa kuno dari pohon jati', true, true, '2023-12-08 22:31:00', '2023-12-16 22:31:00', 1000000, 1),
('Sepeda Kebo', 'Sepeda unik dari jaman Belanda', true, true, '2023-12-07 22:31:00', '2023-12-14 22:31:00', 4000000, 1);

INSERT INTO item_image (image, item_id) VALUES 
('https://ucarecdn.com/2946ad43-ed1c-41ec-b60a-4f6d5f14f82d/lukisanradensaleh.jpg', 2),
('https://ucarecdn.com/c68fd25b-7be4-448d-834e-3c8053de5d9e/sofakayu.jpg', 3),
('https://ucarecdn.com/a6b030a3-bcd5-4c60-bfd1-27e14d891d3f/sepedakebo.jpg', 4);

INSERT INTO item_bid (bid_price, item_id, user_id) VALUES 
(3000000, 2, 1),
(1000000, 3, 1),
(4000000, 4, 1);
