package main

import (
	sgg "github.com/henomis/simple-go-gnuplot"
)

func main() {
	gnuplot := sgg.New("/usr/bin/gnuplot")

	gnuplot.AddEnv("filename", "data.csv")

	gnuplot.AddEnv("width", 800)
	gnuplot.AddEnv("height", 600)

	gnuplot.AddEnv("filetitle", "output.svg")
	gnuplot.AddEnv("headtitle", "This is a graph")

	gnuplot.SetPlotFilePath("./line.plot")

	gnuplot.Exec()

}
