INSERT INTO currencies (id, name) VALUES (
    'ccfdcd34-7885-4f83-8d71-dcc26b0a9be7',
    'USD'
);

INSERT INTO currencies (id, name) VALUES (
    '4bc4dfb8-51ba-41ce-93a3-8cf369357603',
    'EUR'
);

INSERT INTO currency_companies (company_id, currency_id, exchange_rate) VALUES
    ('b918deaf-92ab-485d-9a69-ee7a2a5f4aef', 'ccfdcd34-7885-4f83-8d71-dcc26b0a9be7', 1), -- Primary company -> USD (primary currency)
    ('b918deaf-92ab-485d-9a69-ee7a2a5f4aef', '4bc4dfb8-51ba-41ce-93a3-8cf369357603', 4); -- Primary company -> EUR (secondary currency)