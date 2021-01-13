real_sequences.tex: sheet.tex
	sheet -mode c -in sheet.tex > real_sequences.tex

real_sequences.pdf: ../../*.tex ../../trademark.pdf *.tex real_sequences.tex
	pdflatex real_sequences.tex
	make defs

make:
	make real_sequences.pdf

open: real_sequences.pdf
	open real_sequences.pdf

o:
	make open

defs: real_sequences.tex
	defs real_sequences.tex

spell:
	aspell -c real_sequences.tex

s:
	make spell