# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make measure_vector_space.pdf
	make graph.pdf

measure_vector_space.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > measure_vector_space.tex

measure_vector_space.pdf: ../../*.tex ../../trademark.pdf *.tex measure_vector_space.tex
	pdflatex --file-line-error -interaction=nonstopmode measure_vector_space.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: measure_vector_space.pdf
	open measure_vector_space.pdf

o:
	make open

defs: measure_vector_space.tex
	defs measure_vector_space.tex

spell:
	aspell -c measure_vector_space.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f measure_vector_space.*
	make reset
	bbk_sheet -mode mk > Makefile