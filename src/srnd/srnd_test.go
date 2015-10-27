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




// func signArticle(nntp NNTPMessage, seed []byte) (signed nntpArticle, err error) {



/*

func TestsignArticle(t *testing.T) {

	c := new(NNTPMessage)
	//f  byte := "\xbd" 
	f := []byte{1, 2, 3, 0, 0, 0}
	_, err := signArticle(&c,f)
	
	if err != nil {

		t.Error("Error signing article.", err)

	}

}


*/


/*
func (self articleStore) ReadMessage(r io.Reader) (NNTPMessage, error) {
  return read_message(r)
}

  */


/* 
Convert string to io.Reader

b := bytes.NewBufferString("your string")

*/



/*



func TestReadMessage(t *testing.T) {

	store := articleStore{
	    directory: "/some/dir",
	    temp: config["incoming_dir"],
	    attachments: config["attachments_dir"],
	    thumbs: config["thumbs_dir"],
	    convert_path: config["convert_bin"],
	    ffmpeg_path: config["ffmpegthumbnailer_bin"],
	    sox_path: config["sox_bin"],
	    database: database,
	  }
// Get sox path: `which sox`


	b := bytes.NewBufferString("your string")
	 _, err := store.ReadMessage(b)
	
	if err != nil {

		t.Error("Unable to read message.", err)

	}

}


*/






func TestGenSRNdConfig(t *testing.T) {



	err := GenSRNdConfig()
	// Generate default feeds.ini
	if err != nil {

		t.Error("Error generting srnd.ini", err)

	}

}





/*




func TestGenerateThumbnail(t *testing.T) {


	err := GenSRNdConfig() 


	conf := ReadConfig()


	store := createArticleStore(conf.store, nil)




	err := store.GenerateThumbnail("testdata/i2p.jpeg")
	// Generate default feeds.ini
	if err != nil {

		t.Error("Error Generating thumbnail.", err)

	}

}



  */








