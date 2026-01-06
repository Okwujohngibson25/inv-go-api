-- +goose Up
-- +goose StatementBegin
CREATE TABLE investments(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    plan_id UUID NOT NULL REFERENCES plans(id) ON DELETE CASCADE,
    amount NUMERIC(18,6) NOT NULL,
    "start_date" TIMESTAMP NOT NULL DEFAULT now(),
    end_date TIMESTAMP NOT NULL,
    expected_profit NUMERIC(18,6) NOT NULL,
    "status" TEXT NOT NULL DEFAULT 'Ongoing',
    created_at TIMESTAMP NOT NULL DEFAULT now()
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS investments;
-- +goose StatementEnd
