bounded_linear_norm.tex: sheet.tex
	sheet -mode c -in sheet.tex > bounded_linear_norm.tex

bounded_linear_norm.pdf: ../../*.tex ../../trademark.pdf *.tex bounded_linear_norm.tex
	pdflatex bounded_linear_norm.tex
	make defs

make:
	make bounded_linear_norm.pdf

open: bounded_linear_norm.pdf
	open bounded_linear_norm.pdf

o:
	make open

defs: bounded_linear_norm.tex
	defs bounded_linear_norm.tex

spell:
	aspell -c bounded_linear_norm.tex

s:
	make spell