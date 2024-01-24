CREATE DATABASE IF NOT EXISTS sample_database;
USE sample_database;

CREATE TABLE IF NOT EXISTS people (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    age INTEGER NOT NULL,
    gender TEXT NOT NULL,
    national_id TEXT NOT NULL,
    description TEXT,
    mobile TEXT,
    is_active BOOLEAN NOT NULL
);


INSERT INTO people (first_name, last_name, age, gender, national_id, description, mobile, is_active)
VALUES 
    ('Arsham', 'Karimi', 27, 'Male', '1234567890', 'I live in Iran', '123-456-7890', true),
    ('Alireza', 'Sabz', 25, 'Male', '0987654321', 'I live in Iran', '987-654-3210', false),
    ('Ayda', 'Rezaei', 14, 'female', '5555555555', NULL, '555-555-5555', true);
