# generated automatically by sheet.go

make:
	make differential_mutual_information.pdf

differential_mutual_information.tex: sheet.tex
	sheet -mode c -in sheet.tex > differential_mutual_information.tex

differential_mutual_information.pdf: ../../*.tex ../../trademark.pdf *.tex differential_mutual_information.tex
	pdflatex differential_mutual_information.tex
	make defs


open: differential_mutual_information.pdf
	open differential_mutual_information.pdf

o:
	make open

defs: differential_mutual_information.tex
	defs differential_mutual_information.tex

spell:
	aspell -c differential_mutual_information.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile