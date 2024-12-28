-- 20241226151614_messages.up.sql
CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    task VARCHAR(255) NOT NULL,
    is_done BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP DEFAULT NULL,
    user_id INTEGER,  -- добавляем колонку user_id
    CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE  -- внешнй ключ с ON DELETE CASCADE
);

-- Индекс на user_id для улучшения производительности
CREATE INDEX idx_user_id ON tasks(user_id);


