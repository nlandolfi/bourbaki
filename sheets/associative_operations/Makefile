# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make associative_operations.pdf
	make graph.pdf

associative_operations.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > associative_operations.tex

associative_operations.pdf: ../../*.tex ../../trademark.pdf *.tex associative_operations.tex
	pdflatex --file-line-error -interaction=nonstopmode associative_operations.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: associative_operations.pdf
	open associative_operations.pdf

o:
	make open

terms: associative_operations.tex
	bbk_sheet -mode ts associative_operations.tex

spell:
	aspell -c associative_operations.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f associative_operations.*
	make reset
	bbk_sheet -mode mk > Makefile