equivalence_relations.pdf: *.tex ../*.tex ../trademark.pdf
	pdflatex equivalence_relations.tex
	make defs

make:
	make equivalence_relations.pdf

open: equivalence_relations.pdf
	open equivalence_relations.pdf

defs: equivalence_relations.tex
	defs equivalence_relations.tex
