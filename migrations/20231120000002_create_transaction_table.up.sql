CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    customer_id INT NOT NULL,
    total_points INT NOT NULL
);

CREATE TABLE transaction_vouchers (
    transaction_id INT NOT NULL,
    voucher_id INT NOT NULL,
    FOREIGN KEY (transaction_id) REFERENCES transactions(id),
    FOREIGN KEY (voucher_id) REFERENCES vouchers(id)
);
