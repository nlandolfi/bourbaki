vector_space_dimension.tex: sheet.tex
	sheet -mode c -in sheet.tex > vector_space_dimension.tex

vector_space_dimension.pdf: ../../*.tex ../../trademark.pdf *.tex vector_space_dimension.tex
	pdflatex vector_space_dimension.tex
	make defs

make:
	make vector_space_dimension.pdf

open: vector_space_dimension.pdf
	open vector_space_dimension.pdf

o:
	make open

defs: vector_space_dimension.tex
	defs vector_space_dimension.tex

spell:
	aspell -c vector_space_dimension.tex

s:
	make spell