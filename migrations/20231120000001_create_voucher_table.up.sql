CREATE TABLE vouchers (
    id SERIAL PRIMARY KEY,
    brand_id INT NOT NULL,
    name VARCHAR(255) NOT NULL,
    cost_in_point INT NOT NULL,
    FOREIGN KEY (brand_id) REFERENCES brands(id)
);
