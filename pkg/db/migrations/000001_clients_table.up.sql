CREATE TABLE IF NOT EXISTS client (
    id SERIAL PRIMARY KEY,
    cpf VARCHAR(32) NOT NULL,
    private BOOLEAN NOT NULL,
    incomplete BOOLEAN NOT NULL,
    date_last_purchase DATE,
    average_ticket BIGINT,
    ticket_last_purchase BIGINT,
    frequent_store VARCHAR(32),
    last_store VARCHAR(32)
);
