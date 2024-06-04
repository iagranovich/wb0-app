CREATE TABLE orders (
    order_uid VARCHAR(255) PRIMARY KEY,
    track_number VARCHAR(255) UNIQUE,
    entry VARCHAR(255),
    locale VARCHAR(255),
    internal_signature VARCHAR(255),
    customer_id VARCHAR(255),
    delivery_service VARCHAR(255),
    shardkey VARCHAR(255),
    sm_id INT,
    date_created VARCHAR(255),
    oof_shard VARCHAR(255)
);

CREATE TABLE deliveries (
    order_uid VARCHAR(255) REFERENCES orders(order_uid),
    name VARCHAR(255),
    phone VARCHAR(255),
    zip VARCHAR(255),
    city VARCHAR(255),
    address VARCHAR(255),
    region VARCHAR(255),
    email VARCHAR(255),
    PRIMARY KEY (order_uid)
);

CREATE TABLE payments (
    order_uid VARCHAR(255) REFERENCES orders(order_uid),
    transaction VARCHAR(255),
    request_id VARCHAR(255),
    currency VARCHAR(255),
    provider VARCHAR(255),
    amount INT,
    payment_dt INT,
    bank VARCHAR(255),
    delivery_cost INT,
    goods_total INT,
    custom_fee INT,
    PRIMARY KEY (order_uid)
);

CREATE TABLE items (
    chrt_id INT PRIMARY KEY,
    track_number VARCHAR(255) REFERENCES orders(track_number),
    price INT,
    rid VARCHAR(255),
    name VARCHAR(255),
    sale INT,
    size VARCHAR(255),
    total_price INT,
    nm_id INT,
    brand VARCHAR(255),
    status INT
);