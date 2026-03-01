CREATE OR REPLACE FUNCTION revoke_api_trigger()
    RETURNS trigger AS $$
    BEGIN
        UPDATE api_keys SET revoked = true WHERE user_id = NEW.user_id AND project_id = NEW.project_id;
        RETURN NEW;
    END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER revoke_api_trigger
BEFORE INSERT ON api_keys
FOR EACH ROW
EXECUTE FUNCTION revoke_api_trigger();