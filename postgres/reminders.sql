CREATE TABLE IF NOT EXISTS reminders ( 
    reminders_id bigserial PRIMARY KEY,
    user_id bigint REFERENCES users(id),  
    appointment TIMESTAMP NOT NULL, 
    email TEXT NOT NULL,  
    note TEXT NOT NULL, 
    send BOOLEAN DEFAULT FALSE
); 