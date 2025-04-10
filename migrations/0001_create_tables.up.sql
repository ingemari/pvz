CREATE TABLE users (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  email VARCHAR(256) NOT NULL UNIQUE,
  role VARCHAR(256) NOT NULL CHECK (role IN ('employee', 'moderator')),
  password VARCHAR(256) NOT NULL
);

CREATE TABLE pvz (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  registration_date TIMESTAMPTZ DEFAULT NOW(),
  city VARCHAR(256) NOT NULL CHECK (city IN ('Москва', 'Санкт-Петербург', 'Казань'))
);

CREATE TABLE receptions (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  date_time TIMESTAMPTZ NOT NULL,
  pvz_id UUID NOT NULL REFERENCES pvz(id) ON DELETE CASCADE,
  status VARCHAR(256) NOT NULL CHECK (status IN ('in_progress', 'close'))
);

CREATE TABLE products (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  date_time TIMESTAMPTZ DEFAULT NOW(),
  type VARCHAR(256) NOT NULL CHECK (type IN ('электроника', 'одежда', 'обувь')),
  reception_id UUID NOT NULL REFERENCES receptions(id) ON DELETE CASCADE
);
