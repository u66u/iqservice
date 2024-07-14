CREATE TYPE test_type AS ENUM ('WAIS IV', 'mensa norway');

CREATE TABLE iq_tests (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id),
    test_type test_type,
    correct_answers BOOLEAN[],
    finished BOOLEAN,
    result INT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
