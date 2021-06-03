# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make linearly_dependent_vectors.pdf

linearly_dependent_vectors.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > linearly_dependent_vectors.tex

linearly_dependent_vectors.pdf: ../../*.tex ../../trademark.pdf *.tex linearly_dependent_vectors.tex
	# pdflatex --file-line-error -interaction=nonstopmode linearly_dependent_vectors.tex
	../../latexrun linearly_dependent_vectors.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: linearly_dependent_vectors.pdf
	open linearly_dependent_vectors.pdf

o:
	make open

terms: linearly_dependent_vectors.tex
	bbk_sheet -mode ts linearly_dependent_vectors.tex

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
	rm -f linearly_dependent_vectors.*
	make reset
	bbk_sheet -mode mk > Makefile