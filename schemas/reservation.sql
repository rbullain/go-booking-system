DROP TABLE IF EXISTS reservation;

CREATE TABLE reservation
(
    id      BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id BIGINT NOT NULL,
    room_id BIGINT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES user (id),
    FOREIGN KEY (room_id) REFERENCES room (id)
);