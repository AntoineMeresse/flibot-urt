-- Players
CREATE TABLE IF NOT EXISTS player (
    id SERIAL PRIMARY KEY,
    guid TEXT NOT NULL UNIQUE,
    role INTEGER NOT NULL DEFAULT 0,
    name TEXT NOT NULL,
    ip_address TEXT NOT NULL,
    time_joined TIMESTAMP,
    aliases TEXT NOT NULL
);


--  Runs
CREATE TABLE IF NOT EXISTS runs (
    id SERIAL PRIMARY KEY,
    guid TEXT NOT NULL REFERENCES player(guid),
    utj TEXT NOT NULL,
    mapname TEXT NOT NULL,
    way TEXT NOT NULL,
    runtime INTEGER NOT NULL,
    checkpoints TEXT NOT NULL,
    run_date TIMESTAMP NOT NULL,
    demopath TEXT NOT NULL
);


--  Goto
CREATE TABLE IF NOT EXISTS goto (
    id SERIAL PRIMARY KEY,
    mapname TEXT NOT NULL,
    jumpname TEXT NOT NULL,
    pos_x DOUBLE PRECISION NOT NULL,
    pos_y DOUBLE PRECISION NOT NULL,
    pos_z DOUBLE PRECISION NOT NULL,
    angle_v DOUBLE PRECISION NOT NULL,
    angle_h DOUBLE PRECISION NOT NULL,
    UNIQUE (mapname, jumpname)
);


--  Pen
CREATE TABLE IF NOT EXISTS pen (
    id SERIAL PRIMARY KEY,
    guid TEXT NOT NULL REFERENCES player(guid),
    date DATE NOT NULL,
    size DOUBLE PRECISION NOT NULL,
    UNIQUE (guid, date)
);
