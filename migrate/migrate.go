package main

import (
	"log"

	"github.com/beinganukul/rest_api_golang/initializers"
	"github.com/beinganukul/rest_api_golang/models"
)

func init() {
	initializers.LoadEnVars()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Book{})
	initializers.DB.AutoMigrate(&models.User{})

	// creating a new column with weighted lexems leveraging postgres full text search feature
	initial_result := initializers.DB.Exec(`
	ALTER TABLE books ADD COLUMN search tsvector;
	UPDATE books SET
		search = 
			setweight(to_tsvector('simple', title), 'A') || ' ' ||
			setweight(to_tsvector('simple', array_to_string((COALESCE(author, ARRAY[]::varchar[])), ' '::text)), 'B') || ' ' ||
			setweight(to_tsvector('simple', array_to_string((COALESCE(genre, ARRAY[]::varchar[])), ' '::text)), 'C') || ' ' ||
			setweight(to_tsvector('simple', isbn), 'D');
	`)

	if initial_result.Error != nil {
		log.Fatal("failed to create new search column!")
	}

	create_trigger_function_result := initializers.DB.Exec(`
	CREATE OR REPLACE FUNCTION books_search_trigger() RETURNS trigger AS $$
	begin
	  new.search :=
		setweight(to_tsvector('pg_catalog.simple', coalesce(new.title,'')), 'A') ||
		setweight(to_tsvector('pg_catalog.simple', coalesce(array_to_string(new.author, ' '),'')), 'B') ||
		setweight(to_tsvector('pg_catalog.simple', coalesce(array_to_string(new.genre, ' '),'')), 'C') ||
		setweight(to_tsvector('pg_catalog.simple', coalesce(new.isbn,'')), 'D');
	  return new;
	end
	$$ LANGUAGE plpgsql;	
	`)
	if create_trigger_function_result.Error != nil {
		log.Fatal("failed to create a function to update search column!")
	}

	trigger_function_result := initializers.DB.Exec(`
	CREATE TRIGGER books_search_update BEFORE INSERT OR UPDATE
	ON books FOR EACH ROW EXECUTE PROCEDURE books_search_trigger();
	`)

	if trigger_function_result.Error != nil {
		log.Fatal("failed to create a trigger for update function!")
	}

	// creating a index of search column
	new_result := initializers.DB.Exec(`
		CREATE INDEX IF NOT EXISTS idx_search 
		ON books USING GIN(search);
	`)

	if new_result.Error != nil {
		log.Fatal("failed to create create index of search column!")
	}

	create_query_function_result := initializers.DB.Exec(`
	CREATE OR REPLACE FUNCTION search_books(search_term text)
	RETURNS TABLE(id int, title text, image text[], user_id int, rank numeric) AS $$
	BEGIN
	  RETURN QUERY
		SELECT 
		  books.id::int, 
		  books.title::text, 
		  books.image::text[], 
		  books.user_id::int, 
		  ts_rank(books.search, websearch_to_tsquery('simple', search_term))::numeric as rank
		FROM books
		WHERE books.search @@ websearch_to_tsquery('simple', search_term)
		ORDER BY rank DESC;
	END; $$ LANGUAGE plpgsql;
	`)

	if create_query_function_result.Error != nil {
		log.Fatal("failed to create search function!")
	}

}
