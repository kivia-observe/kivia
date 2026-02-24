CREATE OR REPLACE FUNCTION blacklist_tokens()

RETURNS TRIGGER AS $$

BEGIN
    UPDATE refresh_tokens SET blacklisted = true WHERE user_id = NEW.user_id;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER blacklist_trigger

BEFORE INSERT ON refresh_tokens

FOR EACH ROW

EXECUTE FUNCTION blacklist_tokens();