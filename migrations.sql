create table employee(
id uuid not null,
clock_in TIMESTAMP ,
clock_out TIMESTAMP,
total_hours bigint,
primary key (id,clock_in)
)
