CREATE TABLE IF NOT EXISTS ratings (
  rater_id INTEGER REFERENCES users(id),
  ratee_id INTEGER REFERENCES users(id),
  value INTEGER,
  PRIMARY KEY (rater_id, ratee_id)
);
