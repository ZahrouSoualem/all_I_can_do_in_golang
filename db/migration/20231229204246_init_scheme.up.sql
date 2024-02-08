CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "username" varchar UNIQUE NOT NULL,
  "hash_password" varchar NOT NULL
);


CREATE TABLE "sessions" (
  "id" uuid PRIMARY KEY ,
  "userid" bigint NOT NULL,
  "refresh_token" varchar NOT NULL,
  "isbloced" BOOLEAN NOT NULL DEFAULT false,
  "expires_at" timestamptz NOT NULL ,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);


ALTER TABLE "sessions" ADD FOREIGN KEY ("userid") REFERENCES "users" ("id");