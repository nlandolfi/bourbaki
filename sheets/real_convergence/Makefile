real_convergence.tex: sheet.tex
	sheet -mode c -in sheet.tex > real_convergence.tex

real_convergence.pdf: ../../*.tex ../../trademark.pdf *.tex real_convergence.tex
	pdflatex real_convergence.tex
	make defs

make:
	make real_convergence.pdf

open: real_convergence.pdf
	open real_convergence.pdf

o:
	make open

defs: real_convergence.tex
	defs real_convergence.tex

spell:
	aspell -c real_convergence.tex

s:
	make spell