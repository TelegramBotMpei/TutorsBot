CREATE TABLE IF NOT EXISTS tutors (
    tutor_id serial primary key,
    name varchar(50),
    surname varchar(50),
    password_hash text,
    salt text,
    email text,
    phone text,
    has_free_places boolean
);

CREATE INDEX IF NOT EXISTS tutors_has_free_places ON tutors(has_free_places);
CREATE INDEX IF NOT EXISTS tutors_name ON tutors(name);
CREATE INDEX IF NOT EXISTS tutors_surname ON tutors(surname);

CREATE TABLE IF NOT EXISTS students(
    student_id serial primary key,
    tutor_id int,
    name varchar(50),
    surname varchar(50),
    password_hash text,
    salt text,
    course_work_subject text,
    diploma_work_subject text,
    course int,
    FOREIGN KEY(tutor_id) REFERENCES tutors(tutor_id)
    ON UPDATE CASCADE
    ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS scientific_interests(
    subject_id serial primary key,
    tutor_id int,
    subject text,
    FOREIGN KEY(tutor_id) REFERENCES tutors(tutor_id)
    ON UPDATE CASCADE
    ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS scientific_interest_subject ON scientific_interests(subject);