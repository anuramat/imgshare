-- +goose Up
-- +goose StatementBegin
CREATE TABLE images (
    fileid varchar PRIMARY KEY,
    userid bigint not null,
    description varchar
);
CREATE TABLE votes (
    fileid varchar not null,
    userid bigint not null, 
    upvote boolean not null,
    PRIMARY KEY (fileid, userid)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE images;
DROP TABLE votes;
-- +goose StatementEnd
