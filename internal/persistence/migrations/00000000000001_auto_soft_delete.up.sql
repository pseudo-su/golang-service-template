CREATE OR REPLACE FUNCTION manage_auto_soft_delete(_tbl regclass) RETURNS VOID AS $$
BEGIN
    EXECUTE format('CREATE TABLE _deleted_%s (
        CHECK (deleted_at IS NOT NULL),
        _deleted_id uuid DEFAULT gen_random_uuid() PRIMARY KEY
    ) INHERITS (%s)', _tbl, _tbl);
    EXECUTE format('CREATE TRIGGER auto_soft_delete_%s AFTER UPDATE OF deleted_at OR DELETE ON %s
                    FOR EACH ROW EXECUTE PROCEDURE soft_delete_record()', _tbl, _tbl);
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION soft_delete_record()
RETURNS TRIGGER AS $$
BEGIN
    IF (TG_OP = 'UPDATE' AND NEW.deleted_at IS NOT NULL) THEN
        EXECUTE format('DELETE FROM %I.%I WHERE id = $1', TG_TABLE_SCHEMA, TG_TABLE_NAME) USING OLD.id;
        RETURN OLD;
    END IF;
    IF (TG_OP = 'DELETE') THEN
        IF (OLD.deleted_at IS NULL) THEN
            OLD.deleted_at := now();
        END IF;
        EXECUTE format('INSERT INTO %I.%I SELECT $1.*'
                    , TG_TABLE_SCHEMA, '_deleted_' || TG_TABLE_NAME)
        USING OLD;
    END IF;
    RETURN NULL;
END;
$$ LANGUAGE plpgsql;
