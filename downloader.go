package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

type DownloadFile struct {
	Filename string
	URL      string
	MD5      string
}

//var Verify = flag.Bool("verify", false, "Rechecks files after downloading. This should be left on unless you know what you're doing.")
var Threads = flag.Int("threads", 5, "Amount of files to download at the same time.")

var DownloadCh = make(chan DownloadFile, *Threads*10)
var DownloadWg = sync.WaitGroup{}

func init() {
	for i := 0; i < *Threads; i++ {
		DownloadWg.Add(1)
		go func() {
			downloadthread()
			DownloadWg.Done()
		}()
	}
}

func downloadthread() {
	for download := range DownloadCh {
		ModPrintf("DOWNLOAD", "Downloading %s", download.URL)

		err := actualdownload(download)
		if err == nil {
			continue
		}

		retry := 1
		for err != nil {
			if retry == *Retries {
				ModPrintf("DOWNLOAD", "%d retries exceeded while downloading %s. Failing with error: %v", *Retries, download.URL, err)
				break
			}

			ModPrintf("DOWNLOAD", "Error while downloading %s: %v, retrying (attempt %d/%d)", download.URL, err, retry, *Retries)
			retry++

			err = actualdownload(download)
		}

	}
}

func actualdownload(download DownloadFile) error {
	req, err := http.NewRequest("GET", download.URL, nil)
	if err != nil {
		return err
	}

	out, err := os.Create(filepath.Join(*savePath, download.Filename))
	if err != nil {
		return err
	}

	defer out.Close()

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	n, err := io.Copy(out, res.Body)
	if err != nil {
		return err
	}

	//if !*Verify {
	//	return nil
	//}

	if n != res.ContentLength {
		return fmt.Errorf("content length didn't match with downloaded file size")
	}

	//out.Seek(0, 0)

	//hash := md5.New()
	//_, err = io.Copy(hash, out)
	//if err != nil {
	//	return err
	//}

	//if hex.EncodeToString(hash.Sum(nil)) != download.MD5 {
	//	return fmt.Errorf("hashes didn't match")
	//}

	return nil
}

func Download(download DownloadFile) {
	DownloadCh <- download
}

func WaitDownloadFinish() {
	close(DownloadCh)
	DownloadWg.Wait()
}
