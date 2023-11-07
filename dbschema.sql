CREATE TABLE "Users" (
  "user_id" BIGSERIAL PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "hashedpassword" varchar NOT NULL,
  "role" enum NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "Drones" (
  "drone_id" BIGSERIAL PRIMARY KEY,
  "model" varchar NOT NULL,
  "status" enum NOT NULL,
  "last_active" timestamptz
);

CREATE TABLE "DroneLocations" (
  "location_id" BIGSERIAL PRIMARY KEY,
  "drone_id" bigserial NOT NULL,
  "latitude" float NOT NULL,
  "longitude" float NOT NULL,
  "timestamp" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "SafetyDetectionAlerts" (
  "alert_id" BIGSERIAL PRIMARY KEY,
  "timestamp" timestamptz NOT NULL DEFAULT (now()),
  "frame_id" varchar NOT NULL,
  "drone_id" bigserial NOT NULL,
  "detection_id" bigserial NOT NULL,
  "description" varchar,
  "status" enum NOT NULL,
  "reviewed_by" bigserial NOT NULL
);

CREATE TABLE "DetectionTypes" (
  "detection_id" BIGSERIAL PRIMARY KEY,
  "type" varchar NOT NULL,
  "description" varchar
);

CREATE TABLE "Notifications" (
  "notification_id" BIGSERIAL PRIMARY KEY,
  "user_id" bigserial NOT NULL,
  "message" varchar NOT NULL,
  "timestamp" timestamptz NOT NULL DEFAULT (now()),
  "status" enum NOT NULL,
  "alert_id" bigserial NOT NULL
);

CREATE TABLE "AlertHistory" (
  "history_id" BIGSERIAL PRIMARY KEY,
  "alert_id" bigserial NOT NULL,
  "status" enum NOT NULL,
  "resolution_description" varchar,
  "resolved_by" bigserial NOT NULL,
  "resolution_date" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "SecurityOfficerSchedule" (
  "schedule_id" BIGSERIAL PRIMARY KEY,
  "officer_id" bigserial NOT NULL,
  "start_time" timestamptz NOT NULL,
  "end_time" timestamptz NOT NULL,
  "days" varchar NOT NULL
);

CREATE TABLE "CustomerOfficerMapping" (
  "customer_id" bigserial NOT NULL,
  "officer_id" bigserial NOT NULL
);

CREATE TABLE "StaffCustomerMapping" (
  "staff_id" bigserial NOT NULL,
  "customer_id" bigserial NOT NULL
);

COMMENT ON COLUMN "SecurityOfficerSchedule"."days" IS 'e.g., Mon,Tue,Wed';

ALTER TABLE "Notifications" ADD FOREIGN KEY ("user_id") REFERENCES "Users" ("user_id");

ALTER TABLE "Notifications" ADD FOREIGN KEY ("alert_id") REFERENCES "SafetyDetectionAlerts" ("alert_id");

ALTER TABLE "AlertHistory" ADD FOREIGN KEY ("alert_id") REFERENCES "SafetyDetectionAlerts" ("alert_id");

ALTER TABLE "AlertHistory" ADD FOREIGN KEY ("resolved_by") REFERENCES "Users" ("user_id");

ALTER TABLE "SafetyDetectionAlerts" ADD FOREIGN KEY ("drone_id") REFERENCES "Drones" ("drone_id");

ALTER TABLE "SafetyDetectionAlerts" ADD FOREIGN KEY ("detection_id") REFERENCES "DetectionTypes" ("detection_id");

ALTER TABLE "DroneLocations" ADD FOREIGN KEY ("drone_id") REFERENCES "Drones" ("drone_id");

ALTER TABLE "SecurityOfficerSchedule" ADD FOREIGN KEY ("officer_id") REFERENCES "Users" ("user_id");

ALTER TABLE "CustomerOfficerMapping" ADD FOREIGN KEY ("customer_id") REFERENCES "Users" ("user_id");

ALTER TABLE "CustomerOfficerMapping" ADD FOREIGN KEY ("officer_id") REFERENCES "Users" ("user_id");

ALTER TABLE "StaffCustomerMapping" ADD FOREIGN KEY ("staff_id") REFERENCES "Users" ("user_id");

ALTER TABLE "StaffCustomerMapping" ADD FOREIGN KEY ("customer_id") REFERENCES "Users" ("user_id");
