-- Run after db.ddl has setup the database
-- This file loads sample data into the database for reference purposes

-- Create divisions
INSERT INTO divisions (name) VALUES ('C');
INSERT INTO divisions (name) VALUES ('B');

-- Create teams
INSERT INTO teams (name, division_id, logo_url) VALUES ('Hawks', '1', 'http://placekitten.com/g/64/64');
INSERT INTO teams (name, division_id, logo_url) VALUES ('Bears', '1', 'http://placekitten.com/g/64/64');
INSERT INTO teams (name, division_id, logo_url) VALUES ('Eagles', '2', 'http://placekitten.com/g/64/64');
INSERT INTO teams (name, division_id, logo_url) VALUES ('Grizzlies', '2', 'http://placekitten.com/g/64/64');

-- Create rinks
INSERT INTO rinks (name) VALUES ('Avondale');
INSERT INTO rinks (name) VALUES ('Botany');

-- Create players
INSERT INTO players (jersey_number, player_name, position) VALUES (52, 'Gregg Franco', 'G');
INSERT INTO players (jersey_number, player_name, position) VALUES (73, 'Josh White', 'D');
INSERT INTO players (jersey_number, player_name, position) VALUES (14, 'Joe Shave', 'F');
INSERT INTO players (jersey_number, player_name, position) VALUES (23, 'Andrew Lobb', 'F');

INSERT INTO players (jersey_number, player_name, position) VALUES (21, 'Ben Melville', 'F');
INSERT INTO players (jersey_number, player_name, position) VALUES (69, 'Chen Yao Huang', 'D');

INSERT INTO players (jersey_number, player_name, position) VALUES (86, 'Geoff Combs', 'F');

INSERT INTO players (jersey_number, player_name, position) VALUES (12, 'Chris Noble', 'D');

-- Register players with teams
INSERT INTO team_players (team_id, player_id) VALUES (1, 5);
INSERT INTO team_players (team_id, player_id) VALUES (1, 6);

INSERT INTO team_players (team_id, player_id) VALUES (2, 1);
INSERT INTO team_players (team_id, player_id) VALUES (2, 2);
INSERT INTO team_players (team_id, player_id) VALUES (2, 3);
INSERT INTO team_players (team_id, player_id) VALUES (2, 4);

INSERT INTO team_players (team_id, player_id) VALUES (3, 7);

INSERT INTO team_players (team_id, player_id) VALUES (4, 8);

-- Create matches
INSERT INTO matches (start, season, status, away_id, home_id, away_score, home_score, rink_id) VALUES (NOW(), 2017, 'Upcoming', 1, 2, 0, 0, 1);
INSERT INTO matches (start, season, status, away_id, home_id, away_score, home_score, rink_id) VALUES (NOW() - INTERVAL '2' MONTH, 2017, 'Over', 3, 4, 3, 1, 2);

-- TODO inserts for all the remaining tables

