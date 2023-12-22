-- Create function to trigger update last_update
CREATE OR REPLACE FUNCTION update_doc_and_last_update()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = CURRENT_TIMESTAMP;
  NEW.doc_version = OLD.doc_version + 1;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;