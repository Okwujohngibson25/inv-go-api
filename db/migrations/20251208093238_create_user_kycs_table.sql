-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_kycs(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    "status" TEXT NOT NULL DEFAULT 'Pending',
    id_image_front TEXT NOT NULL,
    id_image_back TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now()

)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_kycs;
-- +goose StatementEnd
