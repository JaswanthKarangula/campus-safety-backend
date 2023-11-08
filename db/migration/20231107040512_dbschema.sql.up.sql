CREATE TABLE "Users" (
                         "user_id" BIGSERIAL PRIMARY KEY,
                         "username" varchar UNIQUE NOT NULL,
                         "email" varchar UNIQUE NOT NULL,
                         "hashedpassword" varchar NOT NULL,
                         "role" varchar NOT NULL,
                         "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "Admins" (
    "admin_id" bigserial UNIQUE NOT NULL
);

CREATE TABLE "Customers" (
                             "customer_id" bigserial UNIQUE NOT NULL,
                             "customer_name" varchar UNIQUE NOT NULL,
                             "admin_id" bigserial NOT NULL,
                             "resource_usage" int
);

CREATE TABLE "SecuritOfficers" (
                                   "officer_id" bigserial UNIQUE NOT NULL,
                                   "customer_id" bigserial NOT NULL
);

CREATE TABLE "SecurityOfficerSchedule" (
                                           "schedule_id" BIGSERIAL PRIMARY KEY,
                                           "officer_id" bigserial NOT NULL,
                                           "start_time" timestamptz NOT NULL,
                                           "end_time" timestamptz NOT NULL,
                                           "day" varchar NOT NULL
);

CREATE TABLE "Drones" (
                          "drone_id" BIGSERIAL PRIMARY KEY,
                          "model" varchar NOT NULL,
                          "customer_id" bigserial NOT NULL,
                          "status" varchar NOT NULL,
                          "last_active" timestamptz
);

CREATE TABLE "Feeds" (
                         "stream_id" BIGSERIAL PRIMARY KEY,
                         "drone_id" bigserial,
                         "status" varchar NOT NULL
);

CREATE TABLE "DetectionTypes" (
                                  "detection_type_id" BIGSERIAL PRIMARY KEY,
                                  "type" varchar NOT NULL,
                                  "description" varchar
);

CREATE TABLE "SafetyDetectionAlerts" (
                                         "alert_id" BIGSERIAL PRIMARY KEY,
                                         "timestamp" timestamptz NOT NULL DEFAULT (now()),
                                         "frame_id" varchar NOT NULL,
                                         "drone_id" bigserial NOT NULL,
                                         "stream_id" bigserial NOT NULL,
                                         "detection_type_id" bigserial NOT NULL,
                                         "description" varchar,
                                         "status" varchar NOT NULL
);

CREATE TABLE "AlertsAssignment" (
                                    "alert_id" BIGSERIAL PRIMARY KEY,
                                    "officer_id" bigserial NOT NULL
);

CREATE TABLE "Notifications" (
                                 "notification_id" BIGSERIAL PRIMARY KEY,
                                 "user_id" bigserial NOT NULL,
                                 "message" varchar NOT NULL,
                                 "timestamp" timestamptz NOT NULL DEFAULT (now()),
                                 "status" varchar NOT NULL,
                                 "alert_id" bigserial NOT NULL
);

CREATE TABLE "AlertHistory" (
                                "history_id" BIGSERIAL PRIMARY KEY,
                                "alert_id" bigserial NOT NULL,
                                "status" varchar NOT NULL,
                                "resolution_description" varchar,
                                "resolved_by" bigserial NOT NULL,
                                "resolution_date" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "Issues" (
                          "issue_id" BIGSERIAL PRIMARY KEY,
                          "description" varchar NOT NULL,
                          "status" varchar NOT NULL,
                          "comments" varchar NOT NULL,
                          "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "SecurityOfficerIssues" (
                                         "customer_id" bigserial NOT NULL,
                                         "officer_id" bigserial NOT NULL,
                                         "issue_id" bigserial NOT NULL
);

CREATE TABLE "CustomerIssues" (
                                  "customer_id" bigserial NOT NULL,
                                  "admin_id" bigserial NOT NULL,
                                  "issue_id" bigserial NOT NULL
);

COMMENT ON COLUMN "Users"."role" IS 'Admin,Customer,Security Officer';

COMMENT ON COLUMN "SecurityOfficerSchedule"."day" IS 'e.g., Mon,Tue,Wed';

COMMENT ON COLUMN "Drones"."status" IS 'active,inactive,streaming';

COMMENT ON COLUMN "Feeds"."status" IS 'active,inactive';

COMMENT ON COLUMN "DetectionTypes"."type" IS 'critical, high, medium, or low';

ALTER TABLE "Admins" ADD FOREIGN KEY ("admin_id") REFERENCES "Users" ("user_id");

ALTER TABLE "Customers" ADD FOREIGN KEY ("admin_id") REFERENCES "Admins" ("admin_id");

ALTER TABLE "Customers" ADD FOREIGN KEY ("customer_id") REFERENCES "Users" ("user_id");

ALTER TABLE "SecuritOfficers" ADD FOREIGN KEY ("customer_id") REFERENCES "Users" ("user_id");

ALTER TABLE "SecuritOfficers" ADD FOREIGN KEY ("officer_id") REFERENCES "Users" ("user_id");

ALTER TABLE "SecurityOfficerSchedule" ADD FOREIGN KEY ("officer_id") REFERENCES "SecuritOfficers" ("officer_id");

ALTER TABLE "Drones" ADD FOREIGN KEY ("customer_id") REFERENCES "Customers" ("customer_id");

ALTER TABLE "Feeds" ADD FOREIGN KEY ("drone_id") REFERENCES "Drones" ("drone_id");

ALTER TABLE "SafetyDetectionAlerts" ADD FOREIGN KEY ("drone_id") REFERENCES "Drones" ("drone_id");

ALTER TABLE "SafetyDetectionAlerts" ADD FOREIGN KEY ("detection_type_id") REFERENCES "DetectionTypes" ("detection_type_id");

ALTER TABLE "SafetyDetectionAlerts" ADD FOREIGN KEY ("stream_id") REFERENCES "Feeds" ("stream_id");

ALTER TABLE "AlertsAssignment" ADD FOREIGN KEY ("alert_id") REFERENCES "SafetyDetectionAlerts" ("alert_id");

ALTER TABLE "AlertsAssignment" ADD FOREIGN KEY ("officer_id") REFERENCES "SecuritOfficers" ("officer_id");

ALTER TABLE "Notifications" ADD FOREIGN KEY ("user_id") REFERENCES "Users" ("user_id");

ALTER TABLE "Notifications" ADD FOREIGN KEY ("alert_id") REFERENCES "SafetyDetectionAlerts" ("alert_id");

ALTER TABLE "AlertHistory" ADD FOREIGN KEY ("alert_id") REFERENCES "SafetyDetectionAlerts" ("alert_id");

ALTER TABLE "AlertHistory" ADD FOREIGN KEY ("resolved_by") REFERENCES "Users" ("user_id");

ALTER TABLE "SecurityOfficerIssues" ADD FOREIGN KEY ("issue_id") REFERENCES "Issues" ("issue_id");

ALTER TABLE "CustomerIssues" ADD FOREIGN KEY ("issue_id") REFERENCES "Issues" ("issue_id");

ALTER TABLE "CustomerIssues" ADD FOREIGN KEY ("customer_id") REFERENCES "Customers" ("customer_id");

ALTER TABLE "CustomerIssues" ADD FOREIGN KEY ("admin_id") REFERENCES "Admins" ("admin_id");
