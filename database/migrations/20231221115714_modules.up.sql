create table modules (
    id uuid primary key,
    name varchar(50) not null,
    features jsonb,
    status status_enum default 'active'::status_enum,
    level module_level not null,
    created_at timestamptz default current_timestamp,
    created_by uuid,
    updated_at timestamptz,
    updated_by uuid,
    deleted_at timestamptz,
    deleted_by uuid,
    doc_version int default 0
);

create index modules_id_idx on modules(id);

CREATE TRIGGER modules_update_trigger
BEFORE UPDATE ON modules
FOR EACH ROW
WHEN (
    OLD.name IS DISTINCT FROM NEW.name OR
    OLD.features IS DISTINCT FROM NEW.features OR
    OLD.status IS DISTINCT FROM NEW.status OR
    OLD.level IS DISTINCT FROM NEW.level
)
EXECUTE FUNCTION update_doc_and_last_update();
