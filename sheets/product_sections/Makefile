product_sections.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex product_sections.tex
	make defs

make:
	make product_sections.pdf

open: product_sections.pdf
	open product_sections.pdf

defs: product_sections.tex
	defs product_sections.tex
