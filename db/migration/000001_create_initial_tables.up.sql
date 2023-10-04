CREATE TABLE IF NOT EXISTS purchase_transaction
(
    id BIGSERIAL PRIMARY KEY,
    description VARCHAR(50) NOT NULL,
    transaction_date TIMESTAMP NOT NULL,
    purchase_amount DECIMAL(64, 2) NOT NULL
)