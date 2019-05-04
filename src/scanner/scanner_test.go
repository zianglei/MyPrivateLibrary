package scanner

import (
	"os"
	"strings"
	"testing"
	"path/filepath"
)


func TestScanFolder(t *testing.T){

	// Scan folder not existed
	if book, err := ScanFolder("./testlib/"); err == nil || book != nil {
		t.Error("Scan folder that not exist didn't raise error!")
	}

	err := os.MkdirAll("./testlib/test1", os.ModePerm)
	if err != nil { t.Errorf("Failed to create a folder for testing scanner function. Error: %s", err.Error())}

	err = os.MkdirAll("./testlib/test1/test2", os.ModePerm)
	if err != nil { t.Errorf("Failed to create a folder for testing scanner function. Error: %s", err.Error())}

	// Scan empty folder
	books, err := ScanFolder("./testlib/")
	if len(books) != 0 { t.Errorf("The number of books when scanning empry folder is incorrect, got: %d, want: %d",
		len(books), 0)
	}

	f, err := os.Create("./testlib/test1/test1.pdf")
	if f != nil { err = f.Close() } else { t.Errorf("Failed to create test1.pdf")}
	f, err = os.Create("./testlib/test1/test2/test2.pdf")
	if f != nil { err = f.Close() } else { t.Errorf("Failed to create test2.pdf")}
	f, err = os.Create("./testlib/test1/test3.txt")
	if f != nil { err = f.Close() } else { t.Errorf("Failed to create test3.pdf")}

	books, err = ScanFolder("./testlib/")
	if len(books) != 2 {
		t.Errorf("The number of books is incorrect, got: %d, want: %d", len(books), 2)
	} else {
		test1Path := filepath.Join("testlib", "test1", "test1.pdf")
		test2Path := filepath.Join("testlib", "test1", "test2", "test2.pdf")

		if strings.Compare(books[0], test1Path) != 0 {
			t.Errorf("The scanned path of test1.pdf is not incorrect, got: %s, want: %s", books[0], test1Path)
		}
		if strings.Compare(books[1], test2Path) != 0 {
			t.Errorf("The scanned path of test2.pdf is not incorrect, got: %s, want: %s", books[1], test2Path)
		}
	}

	err = os.RemoveAll("./testlib")
	if err != nil { t.Errorf("Failed to remove test folder.")}

}
