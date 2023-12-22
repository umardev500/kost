create table role_users (
    id uuid primary key,
    user_id uuid not null,
    role_id uuid not null,
    status status_enum not null,
    created_at timestamptz default current_timestamp,
    created_by uuid,
    updated_at timestamptz,
    updated_by uuid,
    deleted_at timestamptz,
    deleted_by uuid,
    doc_version int default 0,
    unique(user_id, role_id),
    foreign key (user_id)
        references users(id)
        on delete set null
        on update cascade,
    foreign key (role_id)
        references roles(id)
        on delete set null
        on update cascade
);

CREATE TRIGGER role_users_update_trigger
BEFORE UPDATE ON role_users
FOR EACH ROW
WHEN (
    OLD.user_id IS DISTINCT FROM NEW.user_id OR
    OLD.role_id IS DISTINCT FROM NEW.role_id OR
    OLD.status IS DISTINCT FROM NEW.status
)
EXECUTE FUNCTION update_doc_and_last_update();
