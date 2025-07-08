CREATE TABLE CHANNEL (
	channelId VARCHAR(255) PRIMARY KEY,
	username VARCHAR(255) NOT NULL,
	avatar VARCHAR(255) NOT NULL
);

CREATE TABLE VIDEO (
	videoId VARCHAR(255) PRIMARY KEY,
	title VARCHAR(255) NOT NULL,
	thumbnail VARCHAR(255) NOT NULL,
	watched BOOLEAN NOT NULL,
	videoChannel VARCHAR(255) NOT NULL,
	FOREIGN KEY(videoChannel) REFERENCES CHANNEL(channelId)
);

CREATE TABLE CATEGORY (
	categoryId VARCHAR(255) PRIMARY KEY,
	catName VARCHAR(255) NOT NULL,
	catChannel VARCHAR(255) NOT NULL,
	FOREIGN KEY(catChannel) REFERENCES CHANNEL(channelId)
);

CREATE TABLE temp (
	channelId VARCHAR(255) NOT NULL,
	username VARCHAR(255) NOT NULL,
	avatar VARCHAR(255) NOT NULL,
	videoId VARCHAR(255) NOT NULL,
	title VARCHAR(255) NOT NULL,
	thumbnail VARCHAR(255) NOT NULL,
	watched BOOLEAN NOT NULL,
	categoryId VARCHAR(255) NOT NULL,
	catName VARCHAR(255) NOT NULL
);

COPY temp
FROM '/docker-entrypoint-initdb.d/test.csv'
DELIMITER ','
CSV HEADER;

INSERT INTO CHANNEL (channelId, username, avatar)
SELECT DISTINCT channelId, username, avatar
FROM temp;

INSERT INTO VIDEO (videoId, title, thumbnail, watched, videoChannel)
SELECT DISTINCT ON (videoId) videoId, title, thumbnail, watched, channelId
FROM temp
ORDER BY videoId;

INSERT INTO CATEGORY (categoryId, catName, catChannel)
SELECT DISTINCT ON (categoryId) categoryId, catName, channelId
FROM temp
ORDER BY categoryId;

DROP TABLE IF EXISTS temp;
