solving_equations.tex: sheet.tex
	sheet -mode c -in sheet.tex > solving_equations.tex

solving_equations.pdf: ../../*.tex ../../trademark.pdf *.tex solving_equations.tex
	pdflatex solving_equations.tex
	make defs

make:
	make solving_equations.pdf

open: solving_equations.pdf
	open solving_equations.pdf

o:
	make open

defs: solving_equations.tex
	defs solving_equations.tex

spell:
	aspell -c solving_equations.tex

s:
	make spell