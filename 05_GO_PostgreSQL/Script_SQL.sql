CREATE TABLE students (
    id SERIAL NOT NULL,
    name VARCHAR(50),
    age SMALLINT,
    active BOOLEAN,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP
)