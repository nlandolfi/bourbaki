complete_norm_spaces.tex: sheet.tex
	sheet -mode c -in sheet.tex > complete_norm_spaces.tex

complete_norm_spaces.pdf: ../../*.tex ../../trademark.pdf *.tex complete_norm_spaces.tex
	pdflatex complete_norm_spaces.tex
	make defs

make:
	make complete_norm_spaces.pdf

open: complete_norm_spaces.pdf
	open complete_norm_spaces.pdf

o:
	make open

defs: complete_norm_spaces.tex
	defs complete_norm_spaces.tex

spell:
	aspell -c complete_norm_spaces.tex

s:
	make spell