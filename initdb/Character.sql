-- Create characters table
CREATE TABLE characters (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

-- Create emissions table
CREATE TABLE emissions (
    id SERIAL PRIMARY KEY,
    character_id INTEGER NOT NULL,
    rarity INTEGER NOT NULL,
    FOREIGN KEY (character_id) REFERENCES characters(id)
);

-- Insert characters
INSERT INTO characters (name) VALUES
    ('a'), ('b'), ('c'),
    ('d'), ('e'), ('f'),
    ('g'), ('h'), ('i'),
    ('j'), ('k'), ('l'),
    ('m'), ('n'), ('o'),
    ('p'),('q');

-- Insert emissions
INSERT INTO emissions (character_id, rarity) VALUES
    ((SELECT id FROM characters WHERE name = 'a'), 1),
    ((SELECT id FROM characters WHERE name = 'b'), 1),
    ((SELECT id FROM characters WHERE name = 'c'), 1),
    ((SELECT id FROM characters WHERE name = 'd'), 2),
    ((SELECT id FROM characters WHERE name = 'e'), 2),
    ((SELECT id FROM characters WHERE name = 'f'), 2),
    ((SELECT id FROM characters WHERE name = 'g'), 3),
    ((SELECT id FROM characters WHERE name = 'h'), 3),
    ((SELECT id FROM characters WHERE name = 'i'), 3),
    ((SELECT id FROM characters WHERE name = 'j'), 4),
    ((SELECT id FROM characters WHERE name = 'k'), 4),
    ((SELECT id FROM characters WHERE name = 'l'), 4),
    ((SELECT id FROM characters WHERE name = 'm'), 5),
    ((SELECT id FROM characters WHERE name = 'n'), 5),
    ((SELECT id FROM characters WHERE name = 'o'), 5),
    ((SELECT id FROM characters WHERE name = 'p'), 5),
    ((SELECT id FROM characters WHERE name = 'q'), 5);