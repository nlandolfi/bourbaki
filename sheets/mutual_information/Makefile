# generated automatically by sheet.go

make:
	make mutual_information.pdf

mutual_information.tex: sheet.tex
	sheet -mode c -in sheet.tex > mutual_information.tex

mutual_information.pdf: ../../*.tex ../../trademark.pdf *.tex mutual_information.tex
	pdflatex mutual_information.tex
	make defs


open: mutual_information.pdf
	open mutual_information.pdf

o:
	make open

defs: mutual_information.tex
	defs mutual_information.tex

spell:
	aspell -c mutual_information.tex

s:
	make spell