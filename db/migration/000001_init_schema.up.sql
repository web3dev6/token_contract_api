CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "hashed_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "wallet_address" varchar NOT NULL, -- need to make this UNIQUE post CreateWallet() impl
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
  "status" VARCHAR NOT NULL DEFAULT 'PENDING', -- can be an enum instead of string
  "created_at" timestamptz NOT NULL DEFAULT (now())
);
ALTER TABLE "transactions" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");
CREATE INDEX ON "transactions" ("context");
CREATE TABLE tokens (
  "id" BIGSERIAL PRIMARY KEY,
  "username" varchar NOT NULL,
  "address" VARCHAR UNIQUE NOT NULL,
  "name" VARCHAR NOT NULL,
  "symbol" VARCHAR NOT NULL,
  "amount" VARCHAR NOT NULL,
  "owner" VARCHAR NOT NULL,
  "authority" VARCHAR NOT NULL
);
ALTER TABLE "tokens" ADD FOREIGN KEY ("username") REFERENCES "users" ("username"); -- TODO: can also link token to its CREATE_TOKEN txn here
CREATE INDEX ON "tokens" ("address");
