-- TODO refactor with opensource discordd
use opensource_discord;

CREATE TABLE User
(
    id INTEGER AUTO_INCREMENT,
    username TEXT,
    password TEXT,
    PRIMARY KEY (id)
);

CREATE TABLE Friend
(
    id INTEGER AUTO_INCREMENT,
    user1 INTEGER,
    user2 INTEGER,
    FOREIGN KEY (user1) REFERENCES User(id),
    FOREIGN KEY (user2) REFERENCES User(id),
    PRIMARY KEY (id)
);

CREATE TABLE FileUpload
(
    id INTEGER AUTO_INCREMENT,
    filePath TEXT,
    fileSize INTEGER,
    hash INTEGER,
    PRIMARY KEY (id)
);

CREATE TABLE FriendMessage
(
    id INTEGER AUTO_INCREMENT,
    fromUser INTEGER,
    friendRelation INTEGER,
    messageContent TEXT,
    readMessage BOOLEAN,
    sentAt TIMESTAMP,
    file INTEGER,
    FOREIGN KEY (friendRelation) REFERENCES Friend(id),
    FOREIGN KEY (fromUser) REFERENCES User(id),
    FOREIGN KEY (file) REFERENCES FileUpload(id),

    PRIMARY KEY (id)
);

CREATE TABLE Server
(
    id INTEGER,
    name TEXT,
    profilePhoto TEXT,
    PRIMARY KEY (id)
);

CREATE TABLE ServerUser
(
    id INTEGER AUTO_INCREMENT,
    user INTEGER,
    server INTEGER,
    FOREIGN KEY (user) REFERENCES User(id),
    FOREIGN KEY (server) REFERENCES Server(id),
    PRIMARY KEY (id)
);

CREATE TABLE ServerTextChannel
(
    id INTEGER AUTO_INCREMENT,
    server INTEGER,
    FOREIGN KEY (server) REFERENCES Server(id),
    PRIMARY KEY (id)
);

CREATE TABLE ServerMessage
(
    id INTEGER AUTO_INCREMENT,
    serverChannel INTEGER,
    FOREIGN KEY (serverChannel) REFERENCES ServerTextChannel(id),
    PRIMARY KEY (id)
);

CREATE TABLE ServerMessageRead
(
    id INTEGER AUTO_INCREMENT,
    message INTEGER,
    serverChannel INTEGER,
    serverUser INTEGER,
    hasRead BOOLEAN,
    FOREIGN KEY (message) REFERENCES ServerMessage(id),
    FOREIGN KEY (serverChannel) REFERENCES ServerMessageRead(id),
    FOREIGN KEY (serverUser) REFERENCES ServerUser(id),
    PRIMARY KEY (id)
);

CREATE TABLE ServerVoiceChannel
(
    id INTEGER AUTO_INCREMENT,
    server INTEGER,
    channelName TEXT,
    FOREIGN KEY (server) REFERENCES Server(id),
    PRIMARY KEY (id)
);

CREATE TABLE ServerUserInVoiceChannel
(
    id INTEGER AUTO_INCREMENT,
    serverUser INTEGER,
    serverVoiceChannel INTEGER,
    joinedAt TIMESTAMP,
    FOREIGN KEY (serverUser) REFERENCES ServerUser(id),
    FOREIGN KEY (serverVoiceChannel) REFERENCES ServerVoiceChannel(id),
    PRIMARY KEY (id)
);