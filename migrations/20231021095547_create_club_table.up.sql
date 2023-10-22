CREATE TABLE IF NOT EXISTS Club(
	Club_ID bigserial PRIMARY KEY UNIQUE,
	Team_Name varchar(30),
	Stadium_Capacity integer
);
