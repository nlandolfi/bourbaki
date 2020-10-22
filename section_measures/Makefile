section_measures.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex section_measures.tex
	make defs

make:
	make section_measures.pdf

open: section_measures.pdf
	open section_measures.pdf

defs: section_measures.tex
	defs section_measures.tex
