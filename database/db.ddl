create table teams
(
  team_id serial not null
    constraint teams_pkey
    primary key,
  name varchar(20) not null
    constraint teams_name_pk
    unique,
  division_id integer not null,
  logo_url varchar(128)
)
;

create unique index teams_team_id_uindex
  on teams (team_id)
;

create unique index teams_name_uindex
  on teams (name)
;

create table divisons
(
  division_id serial not null
    constraint divisons_pkey
    primary key,
  name varchar(20) not null
)
;

create unique index divisons_division_id_uindex
  on divisons (division_id)
;

create unique index divisons_name_uindex
  on divisons (name)
;

alter table teams
  add constraint teams_divisons_division_id_fk
foreign key (division_id) references divisons
;

create table players
(
  player_id serial not null
    constraint players_pkey
    primary key,
  jersey_number integer not null,
  player_name varchar(48) not null,
  position varchar(1) not null
)
;

create unique index players_player_id_uindex
  on players (player_id)
;

create table matches
(
  match_id serial not null
    constraint matches_pkey
    primary key,
  start time not null,
  season integer default 2017,
  away_id integer not null
    constraint away_team_id_fk
    references teams,
  home_id integer not null
    constraint home_team_id_fk
    references teams,
  away_score integer default '-1'::integer,
  home_score integer default '-1'::integer,
  rink_id integer not null,
  constraint matches_start_home_id_away_id_pk
  unique (start, home_id, away_id)
)
;

create unique index matches_match_id_uindex
  on matches (match_id)
;

create table rinks
(
  rink_id serial not null
    constraint rink_pkey
    primary key,
  name varchar(20) not null
)
;

create unique index rink_rink_id_uindex
  on rinks (rink_id)
;

create unique index rink_name_uindex
  on rinks (name)
;

alter table matches
  add constraint rink_id_fk
foreign key (rink_id) references rinks
;

create table goals
(
  goal_id serial not null
    constraint goals_pkey
    primary key,
  goal_type char(2) not null,
  team_id integer not null
    constraint goals_teams_team_id_fk
    references teams,
  period smallint default 1 not null,
  time integer default 0 not null,
  scorer_id integer not null
    constraint goals_players_player_id_fk
    references players,
  match_id integer not null
    constraint goals_matches_match_id_fk
    references matches,
  constraint goals_team_id_period_time_scorer_id_pk
  unique (team_id, period, time, scorer_id)
)
;

create unique index goals_goal_id_uindex
  on goals (goal_id)
;

create table penalties
(
  penalty_id serial not null
    constraint penalties_pkey
    primary key,
  team_id integer not null
    constraint penalties_teams_team_id_fk
    references teams,
  period integer default 1 not null,
  time integer default 0 not null,
  offense varchar(20) not null,
  offender_id integer not null
    constraint penalties_players_player_id_fk
    references players,
  pim smallint default 2,
  match_id integer not null
    constraint penalties_matches_match_id_fk
    references matches,
  constraint penalties_team_id_period_time_offender_id_offense_pk
  unique (team_id, period, time, offender_id, offense)
)
;

create unique index penalties_penalty_id_uindex
  on penalties (penalty_id)
;

create table match_goals
(
  match_id integer not null
    constraint match_goals_matches_match_id_fk
    references matches,
  goal_id integer not null
    constraint match_goals_goals_goal_id_fk
    references goals,
  constraint match_goals_match_id_goal_id_pk
  primary key (match_id, goal_id)
)
;

create table match_penalties
(
  match_id integer not null
    constraint match_penalties_matches_match_id_fk
    references matches,
  penalty_id integer not null
    constraint match_penalties_penalties_penalty_id_fk
    references penalties,
  constraint match_penalties_match_id_penalty_id_pk
  primary key (match_id, penalty_id)
)
;

create table team_players
(
  team_id integer not null
    constraint team_players_teams_team_id_fk
    references teams,
  player_id integer not null
    constraint team_players_players_player_id_fk
    references players,
  constraint team_players_team_id_player_id_pk
  primary key (team_id, player_id)
)
;

create table goal_assists
(
  goal_id integer not null
    constraint goal_assists_goals_goal_id_fk
    references goals,
  player_id integer not null
    constraint goal_assists_players_player_id_fk
    references players,
  constraint goal_assists_goal_id_player_id_pk
  primary key (goal_id, player_id)
)
;

create table match_player_stats
(
  match_id integer not null
    constraint match_player_stats_matches_match_id_fk
    references matches,
  player_id integer not null
    constraint match_player_stats_players_player_id_fk
    references players,
  goals smallint default 0,
  assists smallint default 0,
  pim smallint default 0,
  constraint match_player_stats_match_id_player_id_pk
  primary key (match_id, player_id)
)
;

create table match_goalie_stats
(
  match_id integer not null
    constraint match_goalie_stats_matches_match_id_fk
    references matches,
  goalie_id integer not null
    constraint match_goalie_stats_players_player_id_fk
    references players,
  shots_faced smallint default 0,
  saves smallint default 0,
  mins smallint default 0,
  constraint match_goalie_stats_match_id_goalie_id_pk
  primary key (match_id, goalie_id)
)
;

