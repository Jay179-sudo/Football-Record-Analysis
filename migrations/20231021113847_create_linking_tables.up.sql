CREATE TABLE IF NOT EXISTS Club_Stats(
	Club_ID bigserial REFERENCES Club(Club_ID),
	Season integer,
	Average_Age NUMERIC(3, 1),
	Squad_Size integer

);

CREATE TABLE IF NOT EXISTS Player_Stats(
	Player_ID bigserial REFERENCES Player(Player_ID),
	Current_Club_ID bigserial REFERENCES Club(Club_ID),
	Season integer,
	Yellow_Cards integer,
	Red_Cards integer,
	Goals integer,
	Assists integer,
	Minutes_Played integer,
	Player_Valuations integer

);