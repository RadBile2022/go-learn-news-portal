CREATE TABLE IF NOT EXISTS "app_news_contents" (
    id SERIAL PRIMARY KEY,
    title VARCHAR(200) NOT NULL,
    excerpt VARCHAR(250) NOT NULL,
    description text NOT NULL,
    image text NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'PUBLISH',
    tags text NOT NULL,
    created_by_id INT REFERENCES  app_news_users(id) ON DELETE CASCADE,
    category_id INT REFERENCES app_news_categories(id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_contents_created_by_id ON app_news_contents(created_by_id);
CREATE INDEX idx_contents_category_id ON app_news_contents(category_id);