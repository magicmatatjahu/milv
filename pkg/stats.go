package pkg

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

type FileStats struct {
	SuccessLinks SuccessLinks
	FailedLinks  FailedLinks
}

type SuccessLinks struct {
	Count int
	Links []Link
}

type FailedLinks struct {
	Count int
	Links []Link
}

type FilesStats []*FileStats

func NewFileStats(file *File) *FileStats {
	fileStat := &FileStats{}
	for _, link := range file.Links {
		if link.Result.Status {
			fileStat.SuccessLinks.Count++
			fileStat.SuccessLinks.Links = append(fileStat.SuccessLinks.Links, link)
		} else {
			fileStat.FailedLinks.Count++
			fileStat.FailedLinks.Links = append(fileStat.FailedLinks.Links, link)
		}
	}
	return fileStat
}

func NewFilesStats(files Files) FilesStats {
	var fileStats FilesStats
	for _, file := range files {
		if file.Stats == nil {
			fileStats = append(fileStats, NewFileStats(file))
		} else {
			fileStats = append(fileStats, file.Stats)
		}
	}
	return fileStats
}

func writeStats(file *File) {
	fmt.Printf("----- %s - status: %v\n", file.RelPath, file.Status)
	for _, link := range file.Links {
		if link.TypeOf == ExternalLink {
			fmt.Printf("- %s", link.AbsPath)
		} else {
			fmt.Printf("- %s", link.RelPath)
		}
		fmt.Printf(" - status: %v", link.Result.Status)
		if link.Result.Message != "" {
			fmt.Printf(", message: %s", link.Result.Message)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func summaryOfFile(file *File) {
	fmt.Printf("----- %s -----", file.RelPath)

	data := [][]string{}
	for _, link := range file.Links {
		var path string
		if link.TypeOf == ExternalLink {
			path = link.AbsPath
		} else {
			path = link.RelPath
		}
		data = append(data, []string{
			path,
			link.Result.Message,
			fmt.Sprintf("%v", link.Result.Status),
		})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Link", "Description", "Status"})
	table.SetRowLine(true)
	table.AppendBulk(data)
	table.Render()
}

func summaryOfFiles(files Files) bool {
	failed := false

	data := [][]string{}
	for _, file := range files {
		if len(file.Stats.FailedLinks.Links) > 0 {
			failed = true
			for _, link := range file.Stats.FailedLinks.Links {
				var path string
				if link.TypeOf == ExternalLink {
					path = link.AbsPath
				} else {
					path = link.RelPath
				}
				data = append(data, []string{
					file.RelPath,
					path,
					link.Result.Message,
				})
			}
		}
	}

	if failed {
		fmt.Printf("#################################################\n")
		fmt.Printf("#                     SUMMARY                   #\n")
		fmt.Printf("#################################################\n\n")
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"File", "Link", "Description"})
		table.SetAutoMergeCells(true)
		table.SetRowLine(true)
		table.AppendBulk(data)
		table.Render()
	}

	return failed
}
