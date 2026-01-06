-- +goose Up
-- +goose StatementBegin
CREATE TABLE plans(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    "name" TEXT NOT NULL,
    min_amount NUMERIC(18,6) NOT NULL,
    max_amount NUMERIC(18,6) NOT NULL,
    duration_days INT NOT NULL CHECK (duration_days > 0),
    roi NUMERIC(18,6) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now()

)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS plans;
-- +goose StatementEnd
