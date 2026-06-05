ALTER TABLE public.generations ADD COLUMN error_message TEXT;

CREATE TABLE IF NOT EXISTS public.generated_templates (
    id UUID PRIMARY KEY,
    generation_id UUID NOT NULL
        REFERENCES generations(id)
        ON DELETE CASCADE,

    type VARCHAR(50) NOT NULL,
    content JSONB NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);