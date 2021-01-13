metric_convergence.tex: sheet.tex
	sheet -mode c -in sheet.tex > metric_convergence.tex

metric_convergence.pdf: ../../*.tex ../../trademark.pdf *.tex metric_convergence.tex
	pdflatex metric_convergence.tex
	make defs

make:
	make metric_convergence.pdf

open: metric_convergence.pdf
	open metric_convergence.pdf

o:
	make open

defs: metric_convergence.tex
	defs metric_convergence.tex

spell:
	aspell -c metric_convergence.tex

s:
	make spell