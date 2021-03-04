# generated automatically by sheet.go

make:
	make differential_mutual_information_graph.pdf

differential_mutual_information_graph.tex: sheet.tex
	sheet -mode c -in sheet.tex > differential_mutual_information_graph.tex

differential_mutual_information_graph.pdf: ../../*.tex ../../trademark.pdf *.tex differential_mutual_information_graph.tex
	pdflatex differential_mutual_information_graph.tex
	make defs


open: differential_mutual_information_graph.pdf
	open differential_mutual_information_graph.pdf

o:
	make open

defs: differential_mutual_information_graph.tex
	defs differential_mutual_information_graph.tex

spell:
	aspell -c differential_mutual_information_graph.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile