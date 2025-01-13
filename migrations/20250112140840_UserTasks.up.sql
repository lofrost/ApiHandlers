ALTER TABLE users
ADD COLUMN tasks UUID[] DEFAULT '{}';

UPDATE users
SET tasks = (
    SELECT ARRAY_AGG(id)
    FROM tasks
    WHERE tasks.user_id = users.id
);