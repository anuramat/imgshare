-- +goose Up
-- +goose StatementBegin
CREATE TABLE images (
    fileid varchar unique not null,
    description varchar
);
CREATE TABLE votes (
    fileid varchar(256) not null,
    userid bigint not null, 
    upvote boolean not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE images;
DROP TABLE votes;
-- +goose StatementEnd
