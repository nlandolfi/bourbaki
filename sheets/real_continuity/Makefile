real_continuity.tex: sheet.tex
	sheet -mode c -in sheet.tex > real_continuity.tex

real_continuity.pdf: ../../*.tex ../../trademark.pdf *.tex real_continuity.tex
	pdflatex real_continuity.tex
	make defs

make:
	make real_continuity.pdf

open: real_continuity.pdf
	open real_continuity.pdf

o:
	make open

defs: real_continuity.tex
	defs real_continuity.tex

spell:
	aspell -c real_continuity.tex

s:
	make spell