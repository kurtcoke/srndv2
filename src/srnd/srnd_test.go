package srnd

import "testing"
import "bytes"

func TestGenFeedsConfig(t *testing.T) {

	err := GenFeedsConfig()
	// Generate default feeds.ini
	if err != nil {

		t.Error("Cannot generate feeds.ini", err)

	}

}


// func (self lineWriter) Write(data []byte) (n int, err error) {



//func OpenFileWriter(fname string) (io.WriteCloser, error) {


func TestOpenFileWriter(t *testing.T) {

	_ ,err := OpenFileWriter("file.txt")
	// Generate default feeds.ini
	if err != nil {

		t.Error("Cant open file writer.", err)

	}

}



//store := createArticleStore(conf.store, nil)

//files, err := store.GetAllAttachments()


/*

func TestGetThreadsPerPage(t *testing.T) {



nntp := createNNTPConnection()
group := nntp.Newsgroup()

//ppb, err := self.database.GetPagesPerBoard(group)

_, err := self.database.GetThreadsPerPage(group)
	// Generate default feeds.ini
	if err != nil {

		t.Error("Cant get threads per page.", err)

	}


}

*/

func TestRenderTo(t *testing.T) {

	c := new(boardModel)

	var b bytes.Buffer

	err := c.RenderTo(&b)
	// Generate default feeds.ini
	if err != nil {

		t.Error("Error rendering bytes.", err)

	}

}


 //(self boardModel) RenderTo(wr io.Writer) error {















