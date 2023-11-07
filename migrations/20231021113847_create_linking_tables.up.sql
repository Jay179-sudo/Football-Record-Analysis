CREATE TABLE IF NOT EXISTS Club_Stats(
	Club_ID bigserial,
	Season integer,
	Wins integer,
	Losses integer,
	Draws integer

);

-- type Club_Stats struct {
-- 	Club_ID int64
-- 	Season  int64
-- 	Wins    int32
-- 	Losses  int32
-- 	Draws   int32
-- }

CREATE TABLE IF NOT EXISTS Player_Stats(
	Player_ID bigserial REFERENCES Player(Player_ID),
	Current_Club_ID bigserial REFERENCES Club(Club_ID),
	Season integer,
	Yellow_Cards NUMERIC(8, 3),
	Red_Cards NUMERIC(8, 3),
	Goals NUMERIC(8, 3),
	Assists NUMERIC(8, 3),
	Minutes_Played integer,
	Player_Valuations integer

);