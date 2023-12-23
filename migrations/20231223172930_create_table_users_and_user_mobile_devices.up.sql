CREATE TABLE "users" (
  "id" bigint PRIMARY KEY,
  "username" varchar(255) UNIQUE,
  "email" varchar(255) UNIQUE,
  "password" varchar(255),
  "bio" varchar(255),
  "created_at" timestamp,
  "updated_at" timestamp,
  "is_deleted" boolean,
  "is_active" boolean
);

CREATE TABLE "user_devices_mobile" (
  "id" bigint PRIMARY KEY,
  "user_id" bigint,
  "firebase_token" varchar(255),
  "is_active" boolean,
  "created_at" timestamp,
  "updated_at" timestamp
);

ALTER TABLE "user_devices_mobile" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
