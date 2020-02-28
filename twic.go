package chess

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func GetTwicArchiveFiles(nr_of_weeks int) (pgnfilename string) {
	// nr_of_weeks gives the latest <nr_of_weeks> Twic archive files concatenated to a large pgn file
	// if nr_of_week is 0: take all the available archive files (now: 0920 - 1319) and concatenated them all

	// calculation last TWIC week nr
	var start int
	t1 := time.Date(2020, time.January, 24, 0, 0, 0, 0, time.UTC)
	end := week_from_03022020(t1, time.Now()) + 1316

	if nr_of_weeks <= 0 {
		start = 920
	} else {
		start = end - nr_of_weeks + 1
	}

	path := ""
	url := "https://www.theweekinchess.com/zips/"
	return do_get_files(start, end, path, url, false)
}

func do_get_files(van int, tot int, path string, url string, _verbose bool) string {
	localpathTWIC := path + "twic/"
	localpathzips := path + "twic/zips/"
	localpathpgns := path + "twic/pgns/"
	if _, err := os.Stat(localpathTWIC); os.IsNotExist(err) {
		os.MkdirAll(localpathTWIC, 0777)
	}
	if _, err := os.Stat(localpathzips); os.IsNotExist(err) {
		os.MkdirAll(localpathzips, 0777)
	}
	if _, err := os.Stat(localpathpgns); os.IsNotExist(err) {
		os.MkdirAll(localpathpgns, 0777)
	}

	if _verbose {
		fmt.Println(localpathTWIC + "-" + localpathzips + "-" + localpathpgns)
	}
	totr := wget_url_range(van, tot, url, localpathzips, _verbose)
	if totr != tot {
		van--
		tot--
		totr = wget_url_range(van, tot, url, localpathzips, _verbose)
	}
	unzip_twic_range(van, totr, url, localpathzips, localpathpgns, _verbose)

	//twicfn := concat_twic_pgn_files(localpathTWIC, localpathpgns, _verbose)
	twicfn := concat_downloaded_twic_pgn_files(van, totr, localpathTWIC, localpathpgns, _verbose)
	if _verbose {
		fmt.Println("File created: " + twicfn)
	}
	return twicfn
}

func wget_url_range(van int, tot int, url string, localpathzips string, verbose bool) int {
	if verbose {
		fmt.Println("Download TWIC files: " + strconv.Itoa(van) + "-" + strconv.Itoa(tot))
	}
	for i := van; i <= tot; i++ {
		name := "twic" + strconv.Itoa(i) + "g.zip"
		urltot := url + name
		localname := localpathzips + name
		PossibleNewFilename := localpathzips + "twic" + padNumberWithZero(uint32(i)) + "g.zip"
		if fileExists(PossibleNewFilename) {
			if verbose {
				fmt.Println("Bestand al gedownload:" + localname)
			}
		} else {
			if verbose {
				fmt.Println("Bestand " + urltot + " wordt gedownload")
			}
			err, http_status := downloadFile(localname, urltot)
			if err != nil || http_status == 404 {
				if verbose {
					fmt.Println("Filenr " + strconv.Itoa(i) + "not available")
				}
				if http_status == 404 {
					os.Remove(localname)
				}
				write_lastfilenr_from_file(i - 1)
				if verbose {
					fmt.Println("Last Filenr:" + strconv.Itoa(i-1))
				}
				return i - 1
				break
			}
			if localname != PossibleNewFilename {
				renameFile(localname, PossibleNewFilename)
			}
		}
	}
	return tot
}

func unzip_twic_range(van int, tot int, url string, localpathzips string, localpathpgns string, verbose bool) {
	if verbose {
		fmt.Println("Unzip TWIC files")
	}
	for i := van; i <= tot; i++ {
		namezip := "twic" + padNumberWithZero(uint32(i)) + "g.zip"
		namepgn := "twic" + padNumberWithZero(uint32(i)) + ".pgn"
		localnamezip := localpathzips + namezip
		localnamepgn := localpathpgns + namepgn
		if fileExists(localnamepgn) {
			if verbose {
				fmt.Println("Bestand al unzipped: " + localnamepgn)
			}
		} else {
			if verbose {
				fmt.Println("Bestand " + localnamezip + " wordt unzipped")
			}
			// Create a ZipFile Object and load sample.zip in it
			if fileExists(localnamezip) {
				unzip(localnamezip, localpathpgns, verbose)
				localnamepgnshort := localpathpgns + "twic" + strconv.Itoa(i) + ".pgn"
				PossibleNewFilename := localpathpgns + "twic" + padNumberWithZero(uint32(i)) + ".pgn"
				if localnamepgnshort != PossibleNewFilename {
					renameFile(localnamepgnshort, PossibleNewFilename)
				}
			}
		}
	}
}

func concat_twic_pgn_files(localpathTWIC string, localpathpgns string, verbose bool) string {
	var pgnfiles []string

	if verbose {
		fmt.Println("Concat Twic Pgn files")
	}
	daystring := time.Now().Format("20060102")
	twicfilename := localpathTWIC + "TWIC_" + daystring + ".pgn"
	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile(twicfilename, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		log.Fatal(err)
	}

	root := localpathpgns
	errw := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if info.IsDir() {
			if verbose {
				fmt.Printf("skipping a dir without errors: %+v \n", info.Name())
			}
		} else {
			pgnfiles = append(pgnfiles, path)
		}
		return nil
	})

	if errw != nil {
		panic(errw)
	}

	if verbose {
		fmt.Println(pgnfiles)
	}

	for _, file := range pgnfiles {
		data, error := ioutil.ReadFile(file)
		if error != nil {
			panic(error)
		}
		if _, err1 := f.Write([]byte(data)); err1 != nil {
			log.Fatal(err1)
		}
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
	return twicfilename
}

func concat_downloaded_twic_pgn_files(van int, tot int, localpathTWIC string, localpathpgns string, verbose bool) string {
	var pgnfiles []string

	if verbose {
		fmt.Println("Concat Twic Pgn files")
	}
	daystring := time.Now().Format("20060102")
	twicfilename := localpathTWIC + "TWIC_" + daystring + ".pgn"
	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile(twicfilename, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		log.Fatal(err)
	}

	for i := van; i <= tot; i++ {
		namepgn := "twic" + padNumberWithZero(uint32(i)) + ".pgn"
		localnamepgn := localpathpgns + namepgn
		pgnfiles = append(pgnfiles, localnamepgn)
	}

	for _, file := range pgnfiles {
		data, error := ioutil.ReadFile(file)
		if error != nil {
			panic(error)
		}
		if _, err1 := f.Write([]byte(data)); err1 != nil {
			log.Fatal(err1)
		}
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
	return twicfilename
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func downloadFile(filepath string, url string) (error, int) {
	// Get the data
	resp, err := http.Get(url)
	status := resp.StatusCode

	if err != nil {
		return err, status
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err, status
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err, status
}

func unzip(_filename string, targetDir string, verbose bool) {
	zipReader, _ := zip.OpenReader(_filename)
	for _, file := range zipReader.Reader.File {

		zippedFile, err := file.Open()
		if err != nil {
			log.Fatal(err)
		}
		defer zippedFile.Close()
		extractedFilePath := filepath.Join(
			targetDir,
			file.Name,
		)

		if file.FileInfo().IsDir() {
			if verbose {
				log.Println("Directory Created:", extractedFilePath)
			}
			os.MkdirAll(extractedFilePath, file.Mode())
		} else {
			if verbose {
				log.Println("File extracted:", file.Name)
			}

			outputFile, err := os.OpenFile(
				extractedFilePath,
				os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
				file.Mode(),
			)
			if err != nil {
				log.Fatal(err)
			}
			defer outputFile.Close()

			_, err = io.Copy(outputFile, zippedFile)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func renameFile(oldFilename string, newFilename string) {
	err := os.Rename(oldFilename, newFilename)
	if err != nil {
		log.Fatal(err)
	}
}

func nowdate_string() string {
	now := time.Now()
	// current date and time
	year := strconv.Itoa(now.Year())
	month := strconv.Itoa(int(now.Month()))
	day := strconv.Itoa(now.Day())
	return year + month + day
}

func week_from_03022020(d1 time.Time, d2 time.Time) int {
	diff := d2.Sub(d1)
	days := int(diff.Hours() / 24)
	weeks := int(days / 7)
	return weeks
}

func padNumberWithZero(value uint32) string {
	return fmt.Sprintf("%04d", value)
}

func read_lastfilenr_from_file(verbose bool) int {
	var content, err = ioutil.ReadFile("Twic_lastfile_nr.txt")
	if err != nil {
		log.Fatal(err)
	}
	text := string(content)
	if verbose {
		fmt.Println(text)
	}
	i, err := strconv.Atoi(text)
	if err == nil {
		return i
	}
	return 1316
}

func write_lastfilenr_from_file(lastfilenr int) {
	//message := []byte()
	err := ioutil.WriteFile("Twic_lastfile_nr.txt", []byte(strconv.Itoa(lastfilenr)), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
