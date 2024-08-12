create table "events"(
    "id" serial primary key,
    "image" varchar (255),
    "title" varchar (80),
    "date" date,
    "description" text,
    "location_id" int references "locations"("id"),
    "created_by" int references "users"("id")
);