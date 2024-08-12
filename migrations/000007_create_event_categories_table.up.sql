create table "event_categories"(
    "id" serial primary key,
    "event_id" int references "events"("id"),
    "category_id" int references "categories"("id")
);