package storage

import "testing"

func TestNoteStruct(t *testing.T) {
	note := Note{
		Id:       1,
		Title:    "Test",
		FullText: "Content",
	}

	if note.Id < 1 {
		t.Error("Not struct not working (Id cannot be < 1)")
	}

	if note.Id != 1 {
		t.Error("Note struct not working (Id is incorrect)")
	}

	if note.Title != "Test" {
		t.Error("Note struct not working (Title is incorrect)")
	}

	if note.FullText != "Content" {
		t.Error("Note struct not working (FullText is incorrect)")
	}
}
