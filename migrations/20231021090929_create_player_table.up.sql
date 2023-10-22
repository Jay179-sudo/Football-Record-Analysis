CREATE TABLE IF NOT EXISTS Player(
	Player_ID bigserial PRIMARY KEY,
	Player_Name varchar(30),
	DOB DATE,
	Position varchar(30)
);