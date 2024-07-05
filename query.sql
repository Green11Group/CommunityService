CREATE TYPE event_type AS ENUM ('workshop', 'seed_exchange', 'community_planting', 'farmers_market');

CREATE TABLE communities (
                             id uuid PRIMARY KEY,
                             name VARCHAR(100) NOT NULL,
                             description TEXT,
                             location VARCHAR(100),
                             created_at VARCHAR
);

CREATE TABLE community_members (
                                   community_id uuid primary key REFERENCES communities(id),
                                   user_id uuid,
                                   joined_at VARCHAR
);

CREATE TABLE events (
                        id uuid PRIMARY KEY,
                        community_id uuid REFERENCES communities(id),
                        name VARCHAR(100) NOT NULL,
                        description TEXT,
                        type event_type,
                        start_time VARCHAR,
                        end_time VARCHAR,
                        location VARCHAR(255)
);

CREATE TABLE forum_posts (
                             id uuid PRIMARY KEY,
                             community_id uuid REFERENCES communities(id),
                             user_id uuid,
                             title VARCHAR(255) NOT NULL,
                             content TEXT,
                             created_at VARCHAR
);

CREATE TABLE Comments (
                          forum_id uuid primary key REFERENCES forum_posts(id),
                          user_id VARCHAR,
                          content TEXT
) ;
