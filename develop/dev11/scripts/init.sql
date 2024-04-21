CREATE TABLE IF NOT EXISTS events (
    event_id SERIAL PRIMARY KEY,
    user_id TEXT,
    week_number INT NOT NULL,
    day_number INT NOT NULL,
    month_name TEXT NOT NULL,
    year_number INT NOT NULL,
    data_title TEXT NOT NULL,
    data_description TEXT
);
