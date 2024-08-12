create table "profile"(
    "id" serial primary key,
    "picture" varchar (255),
    "full_name" varchar (80),
    "birth_date" date,
    "gender" smallint,
    "phone_number" varchar (15),
    "profession" varchar (50),
    "nationality_id" int references "nationalities"("id"),
    "user_id" int references "users"("id")
);