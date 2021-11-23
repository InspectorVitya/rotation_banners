-- +goose Up
CREATE TABLE banners (
                         id serial4 NOT NULL,
                         descr text NULL,
                         CONSTRAINT banners_pk PRIMARY KEY (id),
                         CONSTRAINT banners_un UNIQUE (id)
);
CREATE TABLE social_groups (
                               id serial4 NOT NULL,
                               descr text NOT NULL,
                               CONSTRAINT social_groups_pk PRIMARY KEY (id),
                               CONSTRAINT social_groups_un UNIQUE (id)
);
INSERT INTO social_groups (descr)
VALUES
    ('Young'),
    ('Old');
CREATE TABLE slots (
                       id serial4 NOT NULL,
                       descr text NULL,
                       CONSTRAINT slots_pk PRIMARY KEY (id),
                       CONSTRAINT slots_un UNIQUE (id)
);
INSERT INTO slots (descr)
VALUES
    ('top slot'),
    ('side slot'),
    ('bottom slot');

CREATE TABLE "statistics" (
                              slot_id int4 NOT NULL,
                              banner_id int4 NOT NULL,
                              group_id int4 NULL,
                              count_show int4 NULL DEFAULT 0,
                              count_click int4 NULL DEFAULT 0,
                              CONSTRAINT statistics_fk FOREIGN KEY (banner_id) REFERENCES banners(id) ON DELETE CASCADE,
                              CONSTRAINT statistics_fk_1 FOREIGN KEY (group_id) REFERENCES social_groups(id) ON DELETE CASCADE,
                              CONSTRAINT statistics_fk_2 FOREIGN KEY (slot_id) REFERENCES slots(id) ON DELETE CASCADE
);
CREATE TABLE rotation (
                          banner_id int4 NOT NULL,
                          slot_id int4 NOT NULL,
                          status bool NOT NULL DEFAULT false, -- статус, 0 - не удален, 1 - удален
                          CONSTRAINT rotation_fk FOREIGN KEY (banner_id) REFERENCES banners(id) ON DELETE CASCADE,
                          CONSTRAINT rotation_fk_1 FOREIGN KEY (slot_id) REFERENCES slots(id)
);
-- +goose Down
DROP TABLE statistics;
DROP TABLE rotation;
DROP TABLE slots;
DROP TABLE banners;
DROP TABLE social_groups;