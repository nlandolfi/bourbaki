groups.tex: sheet.tex
	sheet -mode c -in sheet.tex > groups.tex

groups.pdf: ../../*.tex ../../trademark.pdf *.tex groups.tex
	pdflatex groups.tex
	make defs

make:
	make groups.pdf

open: groups.pdf
	open groups.pdf

o:
	make open

defs: groups.tex
	defs groups.tex

spell:
	aspell -c groups.tex

s:
	make spell