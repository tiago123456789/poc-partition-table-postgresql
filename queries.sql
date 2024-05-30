CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(70) NOT NULL,
    email VARCHAR(120) NOT NULL,
    country_code VARCHAR(3)
);


CREATE TABLE IF NOT EXISTS users_partitioned (
    id SERIAL,
    name VARCHAR(70) NOT NULL,
    email VARCHAR(120) NOT NULL,
    country_code VARCHAR(3)
) partition by list(country_code);

CREATE TABLE IF NOT EXISTS users_partitioned_country_code_BR PARTITION OF users_partitioned
    FOR VALUES IN ('BR');

CREATE TABLE IF NOT EXISTS users_partitioned_country_code_CL PARTITION OF users_partitioned
    FOR VALUES IN ('CL');

CREATE TABLE IF NOT EXISTS users_partitioned_country_code_FR PARTITION OF users_partitioned
    FOR VALUES IN ('FR');

CREATE TABLE IF NOT EXISTS users_partitioned_country_code_IN PARTITION OF users_partitioned
    FOR VALUES IN ('IN');

CREATE TABLE IF NOT EXISTS users_partitioned_country_code_JP PARTITION OF users_partitioned
    FOR VALUES IN ('JP');

CREATE TABLE IF NOT EXISTS users_partitioned_country_code_KW PARTITION OF users_partitioned
    FOR VALUES IN ('KW');

CREATE TABLE IF NOT EXISTS users_partitioned_country_code_MX PARTITION OF users_partitioned
    FOR VALUES IN ('MX');

CREATE TABLE IF NOT EXISTS users_partitioned_country_code_USA PARTITION OF users_partitioned
    FOR VALUES IN ('USA');
