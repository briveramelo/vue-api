package data

import "testing"

func Test_Ping(t *testing.T) {
	err := testDB.Ping()
	if err != nil {
		t.Error("Failed to ping database")
	}
}

func TestUser_GetAll(t *testing.T) {
	all, err := models.Book.GetAll()
	if err != nil {
		t.Error("failed to get all books", err)
	}

	if len(all) != 1 {
		t.Error("failed to get the correct number of books")
	}
}

func Test_Book_GetOneByID(t *testing.T) {
	b, err := models.Book.GetOneById(1)
	if err != nil {
		t.Error("failed to get one book by id", err)
	}

	if b.Title != "My Book" {
		t.Errorf("expected title to be 'My Book', but go %s", b.Title)
	}
}

func Test_Book_GetOneBySlug(t *testing.T) {
	b, err := models.Book.GetOneBySlug("my-book")
	if err != nil {
		t.Error("failed to get one book by slug", err)
	}

	if b.Title != "My Book" {
		t.Errorf("expected title to be 'My Book', but go %s", b.Title)
	}

	_, err = models.Book.GetOneBySlug("bad-slug")
	if err == nil {
		t.Error("did not get an error when attempting to fetch non-existing slug")
	}
}