# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make monotone_sequences.pdf
	make graph.pdf

monotone_sequences.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > monotone_sequences.tex

monotone_sequences.pdf: ../../*.tex ../../trademark.pdf *.tex monotone_sequences.tex
	pdflatex --file-line-error -interaction=nonstopmode monotone_sequences.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: monotone_sequences.pdf
	open monotone_sequences.pdf

o:
	make open

defs: monotone_sequences.tex
	defs monotone_sequences.tex

spell:
	aspell -c monotone_sequences.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f monotone_sequences.*
	make reset
	bbk_sheet -mode mk > Makefile