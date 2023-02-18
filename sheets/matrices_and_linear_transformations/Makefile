# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make matrices_and_linear_transformations.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

matrices_and_linear_transformations.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > matrices_and_linear_transformations.tex

matrices_and_linear_transformations.pdf: ../../*.tex ../../trademark.pdf *.tex matrices_and_linear_transformations.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode matrices_and_linear_transformations.tex
	../../latexrun matrices_and_linear_transformations.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: matrices_and_linear_transformations.pdf
	open matrices_and_linear_transformations.pdf

o:
	make open

terms: matrices_and_linear_transformations.tex
	bbk_sheet -mode ts matrices_and_linear_transformations.tex

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
	rm -f matrices_and_linear_transformations.*
	make reset
	bbk_sheet -mode mk > Makefile