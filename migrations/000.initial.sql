CREATE TABLE IF NOT EXISTS currencies (
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS currency_companies (
    company_id VARCHAR(36) NOT NULL,
    currency_id VARCHAR(36) NOT NULL,
    exchange_rate REAL NOT NULL,
    PRIMARY KEY (company_id, currency_id)
);