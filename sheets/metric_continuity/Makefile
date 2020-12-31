metric_continuity.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex metric_continuity.tex
	make defs

make:
	make metric_continuity.pdf

open: metric_continuity.pdf
	open metric_continuity.pdf

o:
	make open

defs: metric_continuity.tex
	defs metric_continuity.tex

spell:
	aspell -c metric_continuity.tex

s:
	make spell
