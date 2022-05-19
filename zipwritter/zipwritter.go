package zipwritter

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type archiveManager struct {
	zip.Compressor
	zip.Decompressor
}

type ArchiveManager interface {
	Zip(targetFolder string, filepaths ...string) error
	Unzip(archivePath string) error
}

func (am *archiveManager) Zip(targetfolder string, filePaths ...string) error {
	var archive *os.File

	if _, err := os.Stat(targetfolder); os.IsNotExist(err) {
		err = os.Mkdir(targetfolder, 0755)
		if err != nil {
			return err
		}
	}

	increment := findNextNumber(targetfolder) 

	log.Println("Creating archive...")
	archive, err := os.Create(targetfolder + "/archive_" + increment + ".zip")
	if err != nil {
		return err
	}

	zr := zip.NewWriter(archive)

	for _, f := range filePaths {
		f1, err := os.Open(f)
		if err != nil {
			os.Remove(targetfolder)
			return err
		}
		defer f1.Close()

		a1, err := zr.Create("archive_"+ increment + "/" + f1.Name())
		if err != nil {
			os.Remove(targetfolder)
			return err
		}

		_, err = io.Copy(a1, f1)
		if err != nil {
			os.Remove(targetfolder)
			return err
		}
	}

	zr.Close()

	return nil
}

func (am *archiveManager) Unzip(archivePath string) error {
	panic("Implement method")
}

func NewArchiveManager(zip.Compressor, zip.Decompressor) ArchiveManager {
	return &archiveManager{}
}

func findNextNumber(path string) string {
	folder, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer folder.Close()

	xd, err := folder.ReadDir(0)
	if err != nil {
		log.Fatal(err)
	}

	lastFileName := strings.TrimSuffix(xd[len(xd)-1].Name(), ".zip")

	num, err := strconv.Atoi(strings.Split(lastFileName,"_")[1])
	if err != nil {
		log.Fatal(err)
	}

	//gerer le changement d'unite 

	num++

	return "00" + strconv.Itoa(num)
}
