measurable_sections.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex measurable_sections.tex
	make defs

make:
	make measurable_sections.pdf

open: measurable_sections.pdf
	open measurable_sections.pdf

defs: measurable_sections.tex
	defs measurable_sections.tex
