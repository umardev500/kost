create table user_details (
    id uuid primary key,
    user_id uuid,
    full_name varchar(50),
    dob date,
    gender gender_enum,
    phone varchar(12),
    avatar text,
    created_at timestamptz default current_timestamp,
    created_by uuid,
    updated_at timestamptz,
    updated_by uuid,
    deleted_at timestamptz,
    deleted_by uuid,
    doc_version int default 0,
    unique(user_id),
    foreign key (user_id)
        references users(id)
        on delete set null
        on update cascade
);

create index user_details_user_id_idx on user_details(user_id);

CREATE TRIGGER user_details_update_trigger
BEFORE UPDATE ON user_details
FOR EACH ROW
WHEN (
    OLD.user_id IS DISTINCT FROM NEW.user_id OR
    OLD.full_name IS DISTINCT FROM NEW.full_name OR
    OLD.dob IS DISTINCT FROM NEW.dob OR
    OLD.gender IS DISTINCT FROM NEW.gender OR
    OLD.phone IS DISTINCT FROM NEW.phone OR
    OLD.avatar IS DISTINCT FROM NEW.avatar
)
EXECUTE FUNCTION update_doc_and_last_update();
