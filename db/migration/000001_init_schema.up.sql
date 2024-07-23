CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "hashed_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT('0001-01-01 00:00:00Z'),
  "created_at" timestamptz NOT NULL DEFAULT (now())
);
CREATE TABLE "sessions" (
  "id" uuid PRIMARY KEY,
  "username" varchar NOT NULL,
  "refresh_token" varchar NOT NULL,
  "user_agent" varchar NOT NULL,
  "client_ip" varchar NOT NULL,
  "is_blocked" boolean NOT NULL,
  "expires_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);
ALTER TABLE "sessions" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");
CREATE INDEX ON "sessions" ("username");
CREATE TABLE "transactions" (
  "id" BIGSERIAL PRIMARY KEY,
  "username" varchar NOT NULL,
  "context" varchar NOT NULL,
  "payload" JSONB NOT NULL,
  "is_confirmed" boolean NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);
ALTER TABLE "transactions" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");
CREATE INDEX ON "sessions" ("username");
