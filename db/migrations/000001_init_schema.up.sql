CREATE TABLE "merchant"(
    "id" BIGSERIAL PRIMARY KEY,
    "code" VARCHAR(100) NOT NULL,
    "name" VARCHAR(50) NOT NULL,
    UNIQUE ("code")  
);

CREATE TABLE "team"(
     "id" BIGSERIAL PRIMARY KEY,
     "email" VARCHAR(50) NOT NULL,
     "merchant_id" INT NULL,
     UNIQUE ("email")
);

ALTER TABLE "team" ADD CONSTRAINT fk_merchant FOREIGN KEY ("merchant_id") REFERENCES "merchant" ("id") ON DELETE SET NULL;

CREATE INDEX "index_on_merchant" ON "merchant" ("code");
CREATE INDEX "index_on_team" ON "team" ("email", "merchant_id");