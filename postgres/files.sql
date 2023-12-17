CREATE TABLE IF NOT EXISTS files ( 
    file_id bigserial PRIMARY KEY, 
    user_id bigint REFERENCES users(id), 
    created_at date DEFAULT NOW(), 
    name text NOT NULL, 
    ext varchar(12) NOT NULL, 
    cat text DEFAULT 'general', 
    prev_name text NOT NULL, 
    type text NOT NULL
); 