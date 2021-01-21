orthogonal_vectors.tex: sheet.tex
	sheet -mode c -in sheet.tex > orthogonal_vectors.tex

orthogonal_vectors.pdf: ../../*.tex ../../trademark.pdf *.tex orthogonal_vectors.tex
	pdflatex orthogonal_vectors.tex
	make defs

make:
	make orthogonal_vectors.pdf

open: orthogonal_vectors.pdf
	open orthogonal_vectors.pdf

o:
	make open

defs: orthogonal_vectors.tex
	defs orthogonal_vectors.tex

spell:
	aspell -c orthogonal_vectors.tex

s:
	make spell