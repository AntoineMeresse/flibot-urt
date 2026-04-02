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


--  Map options
CREATE TABLE IF NOT EXISTS map_options (
    id SERIAL PRIMARY KEY,
    mapname TEXT NOT NULL UNIQUE,
    options TEXT NOT NULL
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

--  Pen Counter
CREATE TABLE IF NOT EXISTS pen_counter (
    id SERIAL PRIMARY KEY,
    guid TEXT NOT NULL REFERENCES player(guid),
    year INTEGER NOT NULL,
    attempts INTEGER NOT NULL DEFAULT 0,
    UNIQUE (guid, year)
);

-- PB Pencoin daily limit
CREATE TABLE IF NOT EXISTS pb_pencoin_daily (
    guid TEXT NOT NULL REFERENCES player(guid),
    date DATE NOT NULL,
    count INTEGER NOT NULL DEFAULT 0,
    PRIMARY KEY (guid, date)
);

-- Ignore list
CREATE TABLE IF NOT EXISTS ignore_list (
    id SERIAL PRIMARY KEY,
    guid TEXT NOT NULL REFERENCES player(guid),
    ignored_guid TEXT NOT NULL REFERENCES player(guid),
    UNIQUE(guid, ignored_guid)
);

-- Bans
CREATE TABLE IF NOT EXISTS bans (
    id SERIAL PRIMARY KEY,
    guid TEXT NOT NULL,
    ip TEXT NOT NULL,
    reason TEXT NOT NULL DEFAULT '',
    banned_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Quotes
CREATE TABLE IF NOT EXISTS quotes (
    id SERIAL PRIMARY KEY,
    text TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Player preferences
CREATE TABLE IF NOT EXISTS player_preferences (
    guid TEXT NOT NULL REFERENCES player(guid) PRIMARY KEY,
    commands TEXT NOT NULL DEFAULT ''
);

-- Servers
CREATE TABLE IF NOT EXISTS server (
    ip TEXT NOT NULL,
    port INTEGER NOT NULL,
    rconpassword TEXT NOT NULL,
    channel_id BIGINT NOT NULL DEFAULT 0,
    name TEXT NOT NULL DEFAULT 'Server',
    PRIMARY KEY (ip, port)
);
