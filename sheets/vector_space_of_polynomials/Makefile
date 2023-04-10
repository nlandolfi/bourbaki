# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make vector_space_of_polynomials.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

vector_space_of_polynomials.tex: sheet.tex macros.tex
	bbk_sheet -mode c -in sheet.tex > vector_space_of_polynomials.tex

vector_space_of_polynomials.pdf: ../../*.tex ../../trademark.pdf *.tex vector_space_of_polynomials.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode vector_space_of_polynomials.tex
	../../latexrun vector_space_of_polynomials.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: vector_space_of_polynomials.pdf
	open vector_space_of_polynomials.pdf

o:
	make open

terms: vector_space_of_polynomials.tex
	bbk_sheet -mode ts vector_space_of_polynomials.tex

spell:
	aspell -c sheet.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f vector_space_of_polynomials.*
	make reset
	bbk_sheet -mode mk > Makefile