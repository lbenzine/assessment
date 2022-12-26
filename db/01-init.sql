-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS news_expenses_id_seq;

-- Table Definition
CREATE TABLE "expenses" (
    "id" int4 NOT NULL DEFAULT nextval('news_expenses_id_seq'::regclass),
    "title" text,
    "amount" float,
    "note" text,
    "tags" text[],
    PRIMARY KEY ("id")
);

INSERT INTO "expenses" ("id", "title", "amount", "note", "tags") VALUES (1, 'strawberry smoothie', 79, 'night market promotion discount 10 bath', '{"food","beverage"}');