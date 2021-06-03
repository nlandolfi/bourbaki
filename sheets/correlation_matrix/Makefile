# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make correlation_matrix.pdf

correlation_matrix.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > correlation_matrix.tex

correlation_matrix.pdf: ../../*.tex ../../trademark.pdf *.tex correlation_matrix.tex
	# pdflatex --file-line-error -interaction=nonstopmode correlation_matrix.tex
	../../latexrun correlation_matrix.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: correlation_matrix.pdf
	open correlation_matrix.pdf

o:
	make open

terms: correlation_matrix.tex
	bbk_sheet -mode ts correlation_matrix.tex

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
	rm -f correlation_matrix.*
	make reset
	bbk_sheet -mode mk > Makefile