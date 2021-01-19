bounded_linear_continuous.tex: sheet.tex
	sheet -mode c -in sheet.tex > bounded_linear_continuous.tex

bounded_linear_continuous.pdf: ../../*.tex ../../trademark.pdf *.tex bounded_linear_continuous.tex
	pdflatex bounded_linear_continuous.tex
	make defs

make:
	make bounded_linear_continuous.pdf

open: bounded_linear_continuous.pdf
	open bounded_linear_continuous.pdf

o:
	make open

defs: bounded_linear_continuous.tex
	defs bounded_linear_continuous.tex

spell:
	aspell -c bounded_linear_continuous.tex

s:
	make spell