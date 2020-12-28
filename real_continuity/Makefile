real_function_continuity.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex real_function_continuity.tex
	make defs

make:
	make real_function_continuity.pdf

open: real_function_continuity.pdf
	open real_function_continuity.pdf

o:
	make open

defs: real_function_continuity.tex
	defs real_function_continuity.tex

spell:
	aspell -c real_function_continuity.tex

s:
	make spell
