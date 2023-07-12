CREATE TABLE IF NOT EXISTS manifold_store
(
    id serial PRIMARY KEY,
    topic_name text NOT NULL,
    data bytea NOT NULL,
);
