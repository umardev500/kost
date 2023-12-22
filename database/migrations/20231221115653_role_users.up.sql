create table role_users (
    id varchar(255) primary key,
    user_id varchar(255) not null,
    role_id varchar(255) not null,
    status status_enum not null,
    created_at timestamptz default current_timestamp,
    created_by varchar(50),
    updated_at timestamptz,
    updated_by varchar(255),
    deleted_at timestamptz,
    deleted_by varchar(255),
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
