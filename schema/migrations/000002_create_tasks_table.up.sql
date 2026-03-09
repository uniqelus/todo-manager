CREATE TABLE state.tasks (
  id uuid PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  description VARCHAR(1000) NOT NULL,
  priority VARCHAR(20) NOT NULL CHECK (priority IN ('undefined','low','medium','high')),
  due_date TIMESTAMPTZ,
  created_at TIMESTAMPTZ DEFAULT NOW(),
  updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE INDEX idx_tasks_priority ON state.tasks (priority);
CREATE INDEX idx_tasks_due_date ON state.tasks (due_date);
CREATE INDEX idx_tasks_created_at ON state.tasks (created_at DESC);
CREATE INDEX idx_tasks_priority_due_date ON state.tasks (priority, due_date);
CREATE INDEX idx_tasks_updated_at ON state.tasks (updated_at DESC);