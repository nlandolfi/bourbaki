# generated automatically by sheet.go

make:
	make mutual_information_graph.pdf

mutual_information_graph.tex: sheet.tex
	sheet -mode c -in sheet.tex > mutual_information_graph.tex

mutual_information_graph.pdf: ../../*.tex ../../trademark.pdf *.tex mutual_information_graph.tex
	pdflatex mutual_information_graph.tex
	make defs


open: mutual_information_graph.pdf
	open mutual_information_graph.pdf

o:
	make open

defs: mutual_information_graph.tex
	defs mutual_information_graph.tex

spell:
	aspell -c mutual_information_graph.tex

s:
	make spell