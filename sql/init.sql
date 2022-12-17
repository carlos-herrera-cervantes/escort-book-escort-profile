CREATE TABLE "public"."nationality" (
    "id" varchar(100) NOT NULL,
    "name" varchar(100) NOT NULL,
    "active" bool NOT NULL,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp NOT NULL,
    PRIMARY KEY ("id")
);

CREATE TABLE "public"."profile_status_category" (
    "id" varchar(100) NOT NULL,
    "name" varchar(100) NOT NULL,
    "active" bool NOT NULL,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp NOT NULL,
    PRIMARY KEY ("id")
);

CREATE TABLE "public"."profile" (
    "id" varchar(100) NOT NULL,
    "escort_id" varchar(100) NOT NULL UNIQUE,
    "first_name" varchar(50),
    "last_name" varchar(50),
    "email" varchar(100) NOT NULL,
    "phone_number" varchar(20),
    "gender" varchar(15),
    "nationality_id" varchar(100),
    "birthdate" date,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp NOT NULL,
    CONSTRAINT "profile_nationality_id_fkey" FOREIGN KEY ("nationality_id") REFERENCES "public"."nationality"("id"),
    PRIMARY KEY ("id")
);

CREATE TABLE "public"."profile_status" (
    "id" varchar(100) NOT NULL,
    "escort_id" varchar(100) NOT NULL,
    "profile_status_category_id" varchar(100) NOT NULL,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp NOT NULL,
    CONSTRAINT "profile_status_profile_status_category_id_fkey" FOREIGN KEY ("profile_status_category_id") REFERENCES "public"."profile_status_category"("id"),
    CONSTRAINT "profile_status_escort_id_fkey" FOREIGN KEY ("escort_id") REFERENCES "public"."profile"("escort_id") ON DELETE CASCADE,
    PRIMARY KEY ("id")
);

CREATE TABLE "public"."attention_site_category" (
    "id" varchar(100) NOT NULL,
    "name" varchar(100) NOT NULL,
    "active" bool NOT NULL,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp NOT NULL,
    PRIMARY KEY ("id")
);

CREATE TABLE "public"."attention_site" (
    "id" varchar NOT NULL,
    "escort_id" varchar,
    "attention_site_category_id" varchar,
    "created_at" timestamp,
    "updated_at" timestamp,
    CONSTRAINT "attention_site_escort_id_fkey" FOREIGN KEY ("escort_id") REFERENCES "public"."profile"("escort_id") ON DELETE CASCADE,
    CONSTRAINT "attention_site_attention_site_category_id_fkey" FOREIGN KEY ("attention_site_category_id") REFERENCES "public"."attention_site_category"("id"),
    PRIMARY KEY ("id")
);

CREATE TABLE "public"."avatar" (
    "id" varchar(100) NOT NULL,
    "path" varchar(100) NOT NULL,
    "escort_id" varchar(100) NOT NULL,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp NOT NULL,
    CONSTRAINT "avatar_escort_id_fkey" FOREIGN KEY ("escort_id") REFERENCES "public"."profile"("escort_id") ON DELETE CASCADE,
    PRIMARY KEY ("id")
);

CREATE TABLE "public"."biography" (
    "id" varchar(100) NOT NULL,
    "description" text NOT NULL,
    "escort_id" varchar(100) NOT NULL,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp NOT NULL,
    CONSTRAINT "biography_escort_id_fkey" FOREIGN KEY ("escort_id") REFERENCES "public"."profile"("escort_id") ON DELETE CASCADE,
    PRIMARY KEY ("id")
);

CREATE TABLE "public"."day" (
    "id" varchar(100) NOT NULL,
    "name" varchar(100) NOT NULL,
    "active" bool NOT NULL,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp NOT NULL,
    PRIMARY KEY ("id")
);

CREATE TABLE "public"."identification_part" (
    "id" varchar(100) NOT NULL,
    "name" varchar(100) NOT NULL,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp NOT NULL,
    PRIMARY KEY ("id")
);

CREATE TABLE "public"."identification" (
    "id" varchar(100) NOT NULL,
    "path" varchar(100) NOT NULL,
    "escort_id" varchar(100) NOT NULL,
    "identification_part_id" varchar(100) NOT NULL,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp NOT NULL,
    CONSTRAINT "identification_identification_part_id_fkey" FOREIGN KEY ("identification_part_id") REFERENCES "public"."identification_part"("id"),
    CONSTRAINT "identification_escort_id_fkey" FOREIGN KEY ("escort_id") REFERENCES "public"."profile"("escort_id") ON DELETE CASCADE,
    PRIMARY KEY ("id")
);

CREATE TABLE "public"."photo" (
    "id" varchar(100) NOT NULL,
    "path" varchar(100) NOT NULL,
    "escort_id" varchar(100) NOT NULL,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp NOT NULL,
    CONSTRAINT "photo_escort_id_fkey" FOREIGN KEY ("escort_id") REFERENCES "public"."profile"("escort_id") ON DELETE CASCADE,
    PRIMARY KEY ("id")
);

CREATE TABLE "public"."time_category" (
    "id" varchar(100) NOT NULL,
    "name" varchar(100) NOT NULL,
    "measurement_unit" varchar(100) NOT NULL,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp NOT NULL,
    PRIMARY KEY ("id")
);

CREATE TABLE "public"."price" (
    "id" varchar(100) NOT NULL,
    "cost" numeric(10,2) NOT NULL,
    "escort_id" varchar(100) NOT NULL,
    "time_category_id" varchar(100) NOT NULL,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp NOT NULL,
    "quantity" int4,
    CONSTRAINT "price_time_category_id_fkey" FOREIGN KEY ("time_category_id") REFERENCES "public"."time_category"("id"),
    CONSTRAINT "price_escort_id_fkey" FOREIGN KEY ("escort_id") REFERENCES "public"."profile"("escort_id") ON DELETE CASCADE,
    PRIMARY KEY ("id")
);

CREATE TABLE "public"."schedule" (
    "id" varchar(100) NOT NULL,
    "from" time NOT NULL,
    "to" time NOT NULL,
    "escort_id" varchar(100) NOT NULL,
    "day_id" varchar(100) NOT NULL,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp NOT NULL,
    CONSTRAINT "schedule_day_id_fkey" FOREIGN KEY ("day_id") REFERENCES "public"."day"("id"),
    CONSTRAINT "schedule_escort_id_fkey" FOREIGN KEY ("escort_id") REFERENCES "public"."profile"("escort_id") ON DELETE CASCADE,
    PRIMARY KEY ("id")
);

CREATE TABLE "public"."service_category" (
    "id" varchar(100) NOT NULL,
    "name" varchar(100) NOT NULL,
    "active" bool NOT NULL,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp NOT NULL,
    PRIMARY KEY ("id")
);

CREATE TABLE "public"."service" (
    "id" varchar(100) NOT NULL,
    "escort_id" varchar(100) NOT NULL,
    "service_category_id" varchar(100) NOT NULL,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp NOT NULL,
    "cost" numeric(10,2) DEFAULT 0,
    CONSTRAINT "service_escort_id_fkey" FOREIGN KEY ("escort_id") REFERENCES "public"."profile"("escort_id"),
    CONSTRAINT "service_escort_id_fkey1" FOREIGN KEY ("escort_id") REFERENCES "public"."profile"("escort_id") ON DELETE CASCADE,
    CONSTRAINT "service_service_category_id_fkey" FOREIGN KEY ("service_category_id") REFERENCES "public"."service_category"("id"),
    PRIMARY KEY ("id")
);
