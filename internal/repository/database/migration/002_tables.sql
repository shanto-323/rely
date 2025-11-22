CREATE TABLE users(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL UNIQUE,
    user_type TEXT NOT NULL,
    token TEXT NOT NULL,
    blocked BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TABLE admins(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,

    issued_by UUID NOT NULL, 
    claimed_by UUID NOT NULL
);

CREATE TRIGGER set_updated_at_admins
    BEFORE UPDATE ON admins
    FOR EACH ROW
    EXECUTE FUNCTION trigger_set_updated_at();

CREATE TABLE students(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,

    student_id INT NOT NULL UNIQUE,
    fullname TEXT NOT NULL,
    email TEXT NOT NULL,
    phone TEXT NOT NULL,
    registration INT NOT NULL,
    department TEXT NOT NULL,
    shift TEXT NOT NULL,
    semester TEXT NOT NULL,
    section TEXT NOT NULL
);

CREATE TRIGGER set_updated_at_students
    BEFORE UPDATE ON students
    FOR EACH ROW
    EXECUTE FUNCTION trigger_set_updated_at();

CREATE TABLE teachers(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,

    teacher_id INT NOT NULL UNIQUE,
    fullname TEXT NOT NULL,
    email TEXT NOT NULL,
    phone TEXT NOT NULL
);

CREATE TRIGGER set_updated_at_teachers
    BEFORE UPDATE ON teachers
    FOR EACH ROW
    EXECUTE FUNCTION trigger_set_updated_at();


CREATE TABLE subjects(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,

    code INT NOT NULL UNIQUE,
    name TEXT NOT NULL,
    credits INT NOT NULL,
    semester TEXT NOT NULL 
);

CREATE TRIGGER set_updated_at_subjects
    BEFORE UPDATE ON subjects
    FOR EACH ROW
    EXECUTE FUNCTION trigger_set_updated_at();

CREATE TABLE tokens(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,

    token TEXT NOT NULL,
    issued_by UUID REFERENCES teachers(id) ON DELETE CASCADE,
    claimed_by UUID NOT NULL,
    valid BOOLEAN NOT NULL DEFAULT TRUE
);

CREATE TRIGGER set_updated_at_tokens
    BEFORE UPDATE ON tokens
    FOR EACH ROW
    EXECUTE FUNCTION trigger_set_updated_at();


CREATE TABLE attendance_sessions(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,

    teacher_id UUID REFERENCES teachers(id) ON DELETE CASCADE,
    subject_code INT NULL REFERENCES subjects(code) ON DELETE SET NULL,
    department TEXT NOT NULL,
    shift TEXT NOT NULL,
    semester TEXT NOT NULL,
    section TEXT NOT NULL
);

CREATE TRIGGER set_updated_at_attendance_sessions
    BEFORE UPDATE ON attendance_sessions
    FOR EACH ROW
    EXECUTE FUNCTION trigger_set_updated_at();


CREATE TABLE attendance_records(
    session_id UUID REFERENCES attendance_sessions(id) ON DELETE CASCADE,
    student_id INT REFERENCES students(student_id) ON DELETE CASCADE,
    PRIMARY key (session_id, student_id)
);
