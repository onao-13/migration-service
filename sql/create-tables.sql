CREATE TABLE products(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255),
    price DECIMAL,
    image_url VARCHAR(255),
    article INT,
    rating FLOAT DEFAULT 0.0,
    brand VARCHAR(50),
    country VARCHAR(50),
    weight INT,
    sale INT,
    category_id BIGINT
);

CREATE TABLE reviews(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(1000),
    description VARCHAR(255),
    date DATE,
    rating INT,
    product_id BIGINT
);

ALTER TABLE reviews
ADD CONSTRAINT FK_Reviews_Product FOREIGN KEY (product_id)
    REFERENCES products(id);

CREATE TABLE categories(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(50),
    image_url VARCHAR(255)
);

CREATE TABLE articles(
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR(100),
    description VARCHAR(1000),
    date DATE,
    image_url VARCHAR(255)
);

CREATE TABLE special_offers(
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR(100),
    description VARCHAR(1000),
    image_url VARCHAR(255)
);

ALTER TABLE products
ADD CONSTRAINT FK_Products_Categories FOREIGN KEY(category_id)
    REFERENCES categories(id);

CREATE TABLE users(
    id BIGSERIAL PRIMARY KEY,
    uuid UUID NOT NULL,
    name VARCHAR(50) NOT NULL,
    password BYTEA NOT NULL,
    number VARCHAR(12) NOT NULL
);

CREATE TABLE orders(
    id BIGSERIAL PRIMARY KEY,
    uuid UUID NOT NULL,
    user_uuid UUID NOT NULL,
    total_price DECIMAL NOT NULL,
    buy_date DATE DEFAULT NOW(),
    payment_type VARCHAR(10) NOT NULL
);

CREATE TABLE order_product(
    id BIGSERIAL PRIMARY KEY,
    order_uuid UUID NOT NULL,
    product_id BIGINT NOT NULL
);

ALTER TABLE order_product
ADD CONSTRAINT FK_Order_Product_Product_Id FOREIGN KEY(product_id)
    REFERENCES products(id);