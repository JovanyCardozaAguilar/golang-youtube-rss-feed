CREATE TABLE CATEGORY (
	categoryId VARCHAR(255) PRIMARY KEY,
	catName VARCHAR(255) NOT NULL
);

CREATE TABLE CHANNEL (
	channelId VARCHAR(255) PRIMARY KEY,
	username VARCHAR(255) NOT NULL,
	avatar VARCHAR(255) NOT NULL
);

CREATE TABLE VIDEO (
	videoId VARCHAR(255) PRIMARY KEY,
	vChannelId VARCHAR(255) NOT NULL,
	title VARCHAR(255) NOT NULL,
	thumbnail VARCHAR(255) NOT NULL,
	watched BOOLEAN NOT NULL,
	FOREIGN KEY(vChannelId) REFERENCES CHANNEL(channelId)
);

CREATE TABLE CHANNEL_CATEGORY (
	ccChannelId VARCHAR(255) NOT NULL,
	ccCategoryId VARCHAR(255) NOT NULL,
	PRIMARY KEY (ccChannelId, ccCategoryId),
	FOREIGN KEY(ccChannelId) REFERENCES CHANNEL(channelId),
	FOREIGN KEY(ccCategoryId) REFERENCES CATEGORY(categoryId)
);

CREATE TABLE VIDEO_CATEGORY (
	vcVideoId VARCHAR(255) NOT NULL,
	vcCategoryId VARCHAR(255) NOT NULL,
	PRIMARY KEY (vcVideoId, vcCategoryId),
	FOREIGN KEY(vcVideoId) REFERENCES VIDEO(videoId),
	FOREIGN KEY(vcCategoryId) REFERENCES CATEGORY(categoryId)
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

INSERT INTO VIDEO (videoId, vChannelId, title, thumbnail, watched)
SELECT DISTINCT ON (videoId) videoId, channelId, title, thumbnail, watched
FROM temp
ORDER BY videoId;

INSERT INTO CATEGORY (categoryId, catName)
SELECT DISTINCT ON (categoryId) categoryId, catName
FROM temp
ORDER BY categoryId;

INSERT INTO CHANNEL_CATEGORY (ccChannelId, ccCategoryId)
SELECT DISTINCT channelId, categoryId
FROM temp;

INSERT INTO VIDEO_CATEGORY (vcVideoId, vcCategoryId)
SELECT videoId, categoryId
FROM temp;

DROP TABLE IF EXISTS temp;
